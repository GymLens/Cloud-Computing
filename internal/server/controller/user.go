package controller

import (
	"github.com/GymLens/Cloud-Computing/db"
	"github.com/GymLens/Cloud-Computing/services"
	"github.com/gofiber/fiber/v2"

	firebase "firebase.google.com/go"
)

type UserController struct {
	firebaseApp *firebase.App
}

func NewUserController(app *firebase.App) *UserController {
	return &UserController{
		firebaseApp: app,
	}
}

func (ctr *UserController) GetProfile(c *fiber.Ctx) error {
	uid := c.Locals("uid").(string)

	user, err := db.GetUser(c.Context(), ctr.firebaseApp, uid)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	return c.JSON(user)
}

func (ctr *UserController) UploadAvatar(c *fiber.Ctx) error {
	uid := c.Locals("uid").(string)

	file, err := c.FormFile("avatar")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Missing file"})
	}

	src, err := file.Open()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Unable to read file"})
	}
	defer src.Close()

	storageService, err := services.NewStorageService(ctr.firebaseApp, "gym-lens-bucket")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to initialize storage service"})
	}

	objectName := "avatars/" + uid + "/" + file.Filename
	if err := storageService.UploadFile(c.Context(), objectName, src); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to upload file"})
	}

	avatarURL := "https://storage.googleapis.com/" + storageService.BucketName + "/" + objectName
	updateData := map[string]interface{}{
		"avatarUrl": avatarURL,
	}
	if err := db.UpdateUser(c.Context(), ctr.firebaseApp, uid, updateData); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update user data"})
	}

	return c.JSON(fiber.Map{"message": "Avatar uploaded successfully", "avatarUrl": avatarURL})
}
