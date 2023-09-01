package pages

import "github.com/gofiber/fiber/v2"

func getDonate(ctx *fiber.Ctx) error {
	return ctx.Render("pages/donate", fiber.Map{})
}
