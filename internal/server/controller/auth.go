package controller

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	firebaseAuthPkg "firebase.google.com/go/v4/auth" // Firebase Auth package
	"github.com/GymLens/Cloud-Computing/config"
	customAuthPkg "github.com/GymLens/Cloud-Computing/pkg/auth" // custom Auth package
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// AuthController for auth-related ops
type AuthController struct {
	AuthClient firebaseAuthPkg.Client
	Config     *config.Config
	Validator  *validator.Validate
}

// NewAuthController init a new AuthController.
func NewAuthController(firebaseAuth *customAuthPkg.FirebaseAuth, cfg *config.Config, v *validator.Validate) *AuthController {
	if firebaseAuth == nil || firebaseAuth.Client == nil {
		return &AuthController{}
	}
	return &AuthController{
		AuthClient: *firebaseAuth.Client,
		Config:     cfg,
		Validator:  v,
	}
}

// SignUpRequest for sign-up
type SignUpRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Name     string `json:"name" validate:"required"`
}

type SignUpResponse struct {
	UID    string `json:"uid"`
	Email  string `json:"email"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

// SignInRequest for sign-in
type SignInRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type SignInResponse struct {
	IDToken      string `json:"idToken"`
	RefreshToken string `json:"refreshToken"`
	ExpiresIn    string `json:"expiresIn"`
	TokenType    string `json:"tokenType"`
}

// SignUp handle via Firebase.
func (ac *AuthController) SignUp(c *fiber.Ctx) error {
	var req SignUpRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	if err := ac.Validator.Struct(req); err != nil {
		var errors []string
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, fmt.Sprintf("Field '%s' failed on the '%s' tag", err.Field(), err.Tag()))
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": errors,
		})
	}

	// Create a new user with Firebase Auth
	params := (&firebaseAuthPkg.UserToCreate{}).
		Email(req.Email).
		EmailVerified(false).
		Password(req.Password).
		DisplayName(req.Name).
		Disabled(false)

	userRecord, err := ac.AuthClient.CreateUser(context.Background(), params)
	if err != nil {
		var firebaseError map[string]interface{}
		if err := json.Unmarshal([]byte(err.Error()), &firebaseError); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to parse Firebase error",
			})
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": firebaseError["error"],
		})
	}

	res := SignUpResponse{
		UID:    userRecord.UID,
		Email:  userRecord.Email,
		Name:   userRecord.DisplayName,
		Status: "User created successfully",
	}

	return c.Status(fiber.StatusCreated).JSON(res)
}

// SignIn for user auth via Firebase Auth REST API.
func (ac *AuthController) SignIn(c *fiber.Ctx) error {
	if ac.Config == nil || ac.Config.FirebaseAPIKey == "" {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Firebase API Key is not configured",
		})
	}

	var req SignInRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	// Validate the request payload
	if err := ac.Validator.Struct(req); err != nil {
		// Collect all validation errors
		var errors []string
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, fmt.Sprintf("Field '%s' failed on the '%s' tag", err.Field(), err.Tag()))
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": errors,
		})
	}

	// Prepare the payload for Firebase Auth REST API
	payload := map[string]string{
		"email":             req.Email,
		"password":          req.Password,
		"returnSecureToken": "true",
	}
	payloadBytes, _ := json.Marshal(payload)

	// Make the POST request to Firebase Auth REST API
	resp, err := http.Post(
		fmt.Sprintf("https://identitytoolkit.googleapis.com/v1/accounts:signInWithPassword?key=%s", ac.Config.FirebaseAPIKey),
		"application/json",
		bytes.NewBuffer(payloadBytes),
	)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to communicate with Firebase Auth",
		})
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var firebaseError map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&firebaseError); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to parse Firebase error",
			})
		}
		return c.Status(resp.StatusCode).JSON(fiber.Map{
			"error": firebaseError["error"],
		})
	}

	var signInRes SignInResponse
	if err := json.NewDecoder(resp.Body).Decode(&signInRes); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to parse Firebase Auth response",
		})
	}

	return c.JSON(signInRes)
}
