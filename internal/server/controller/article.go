package controller

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"github.com/GymLens/Cloud-Computing/db"
	"github.com/gofiber/fiber/v2"
)

type ArticleController struct {
	app *firebase.App
}

func NewArticleController(app *firebase.App) *ArticleController {
	return &ArticleController{
		app: app,
	}
}

func (ac *ArticleController) GetArticles(c *fiber.Ctx) error {
	ctx := context.Background()

	articles, err := db.GetArticles(ctx, ac.app)
	if err != nil {
		log.Printf("Error fetching articles: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  500,
			"message": "Internal Server Error",
		})
	}

	response := fiber.Map{
		"status":  200,
		"message": "Success",
		"data":    articles,
	}

	return c.JSON(response)
}
