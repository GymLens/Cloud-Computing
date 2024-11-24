package controller

import (
	"log"
	"time"

	firebase "firebase.google.com/go"
	"github.com/GymLens/Cloud-Computing/api"
	"github.com/GymLens/Cloud-Computing/db"
	"github.com/GymLens/Cloud-Computing/models"
	"github.com/GymLens/Cloud-Computing/services"
	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	authService *services.AuthService
	firebaseApp *firebase.App
}

func NewAuthController(app *firebase.App) *AuthController {
	authService, err := services.NewAuthService(app)
	if err != nil {
		panic(err)
	}
	return &AuthController{
		authService: authService,
		firebaseApp: app,
	}
}

func (ctr *AuthController) SignUp(c *fiber.Ctx) error {
	var req api.SignUpRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	userRecord, err := ctr.authService.CreateUserWithEmailAndPassword(req.Email, req.Password, req.Name)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	user := models.User{
		UID:       userRecord.UID,
		Name:      req.Name,
		Email:     userRecord.Email,
		Created:   time.Now(),
		SignedIn:  time.Now(),
		AvatarURL: "",
	}

	if err := db.CreateUser(c.Context(), ctr.firebaseApp, &user); err != nil {
		log.Printf("Error saving user data to Firestore: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to save user data", "details": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"uid": user.UID, "email": user.Email, "name": user.Name})
}

func (ctr *AuthController) SignIn(c *fiber.Ctx) error {
	var req api.SignInRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	idToken, err := ctr.authService.SignIn(req.Email, req.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}

	// Get user by email to retrieve UID
	userRecord, err := ctr.authService.GetUserByEmail(req.Email)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve user"})
	}

	// Update SignedIn timestamp
	if err := db.UpdateUser(c.Context(), ctr.firebaseApp, userRecord.UID, map[string]interface{}{
		"signedIn": time.Now(),
	}); err != nil {
		log.Printf("Error updating user sign-in time: %v", err)
	}

	return c.JSON(fiber.Map{"token": idToken})
}
