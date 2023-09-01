package pages

import "github.com/gofiber/fiber/v2"

func getHome(ctx *fiber.Ctx) error {
	return ctx.Render("pages/home", fiber.Map{})
}
