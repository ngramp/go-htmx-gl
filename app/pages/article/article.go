package article

import (
	"GoHtmxGL/models"

	"github.com/gofiber/fiber/v2"
)

func Routes(api fiber.Router) {
	articlesGroup := api.Group("/articles")
	{
		articlesGroup.Get("/", getArticles)
		articlesGroup.Get("/:id", getArticle)
		articlesGroup.Post("/", createArticle)
		articlesGroup.Put("/:id", updateArticle)
		articlesGroup.Delete("/:id", deleteArticle)
	}
}

// getArticles handles the GET request to fetch all article
func getArticles(ctx *fiber.Ctx) error {
	// Logic to fetch and return all users
	var articles []models.Article
	models.DB.Find(&articles)
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"data": articles})
}

// getArticle handles the GET request to fetch a specific article
func getArticle(ctx *fiber.Ctx) error {
	// Logic to fetch and return a specific user
	// todo: check id is sound first
	id := ctx.Params("id")
	var article models.Article
	models.DB.First(&article, id)
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"data": article})
}

// createArticle handles the POST request to create a new article
func createArticle(ctx *fiber.Ctx) error {
	// Logic to create a new user
	ctx.Status(fiber.StatusNotImplemented)
	return ctx.SendString("Endpoint not implemented yet")
}

// updateArticle handles the PUT request to update an article
func updateArticle(ctx *fiber.Ctx) error {
	// Logic to update a user
	ctx.Status(fiber.StatusNotImplemented)
	return ctx.SendString("Endpoint not implemented yet")
}

// deleteArticle handles the DELETE request to delete an article
func deleteArticle(ctx *fiber.Ctx) error {
	// Logic to delete a user
	ctx.Status(fiber.StatusNotImplemented)
	return ctx.SendString("Endpoint not implemented yet")
}
