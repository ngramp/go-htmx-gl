package pages

import "github.com/gofiber/fiber/v2"

func getContact(ctx *fiber.Ctx) error {
	return ctx.Render("pages/contact", fiber.Map{})
}
