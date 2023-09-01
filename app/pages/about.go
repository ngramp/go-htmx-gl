package pages

import "github.com/gofiber/fiber/v2"

func getAbout(ctx *fiber.Ctx) error {
	return ctx.Render("pages/about", fiber.Map{})
}
