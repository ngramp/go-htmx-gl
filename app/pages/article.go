package pages

import (
	"GoHtmxGL/models"

	"github.com/gofiber/fiber/v2"
)

func GetArticles(ctx *fiber.Ctx) error {
	// Logic to fetch and return all users
	var articles []models.Article
	models.DB.Find(&articles)
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"data": articles})
}

func GetArticle(ctx *fiber.Ctx) error {
	// Logic to fetch and return a specific user
	// todo: check id is sound first
	id := ctx.Params("id")
	var article models.Article
	models.DB.First(&article, id)
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"data": article})
}

func CreateArticle(ctx *fiber.Ctx) error {
	// Logic to create a new user
	ctx.Status(fiber.StatusNotImplemented)
	return ctx.SendString("Endpoint not implemented yet")
}

func UpdateArticle(ctx *fiber.Ctx) error {
	// Logic to update a user
	ctx.Status(fiber.StatusNotImplemented)
	return ctx.SendString("Endpoint not implemented yet")
}

func DeleteArticle(ctx *fiber.Ctx) error {
	// Logic to delete a user
	ctx.Status(fiber.StatusNotImplemented)
	return ctx.SendString("Endpoint not implemented yet")
}
