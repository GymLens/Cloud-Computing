package controller

import (
	"strconv"
	"sync"

	"github.com/GymLens/Cloud-Computing/models"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	users  map[int64]*models.User
	mu     sync.Mutex
	nextID int64
}

func NewUserController() *UserController {
	return &UserController{
		users:  make(map[int64]*models.User),
		nextID: 1,
	}
}

// CreateUser
func (uc *UserController) CreateUser(c *fiber.Ctx) error {
	var user models.User

	// Parse the request body into the User struct
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	if user.Email == "" || user.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Email and Name are required",
		})
	}

	// Assign a unique ID and store the user
	uc.mu.Lock()
	user.ID = uc.nextID
	uc.nextID++
	uc.users[user.ID] = &user
	uc.mu.Unlock()

	return c.Status(fiber.StatusCreated).JSON(user)
}

// GetUser by ID
func (uc *UserController) GetUser(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}

	uc.mu.Lock()
	user, exists := uc.users[id]
	uc.mu.Unlock()

	if !exists {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	return c.JSON(user)
}

// GetAllUsers
func (uc *UserController) ListUsers(c *fiber.Ctx) error {
	uc.mu.Lock()
	users := make([]*models.User, 0, len(uc.users))
	for _, user := range uc.users {
		users = append(users, user)
	}
	uc.mu.Unlock()

	return c.JSON(users)
}

// UpdateUser by ID
func (uc *UserController) UpdateUser(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}

	var updatedData models.User
	if err := c.BodyParser(&updatedData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	uc.mu.Lock()
	user, exists := uc.users[id]
	if !exists {
		uc.mu.Unlock()
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	if updatedData.Email != "" {
		user.Email = updatedData.Email
	}
	if updatedData.Name != "" {
		user.Name = updatedData.Name
	}

	uc.mu.Unlock()

	return c.JSON(user)
}

// DeleteUser by ID
func (uc *UserController) DeleteUser(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}

	uc.mu.Lock()
	_, exists := uc.users[id]
	if !exists {
		uc.mu.Unlock()
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	delete(uc.users, id)
	uc.mu.Unlock()

	return c.SendStatus(fiber.StatusNoContent) // 204 No Content
}
