package pages

import "github.com/gofiber/fiber/v2"

func Donate(ctx *fiber.Ctx) error {
	return ctx.Render("pages/donate", fiber.Map{})
}
func Home(ctx *fiber.Ctx) error {
	return ctx.Render("pages/home", fiber.Map{})
}
func Contact(ctx *fiber.Ctx) error {
	return ctx.Render("pages/contact", fiber.Map{})
}
func About(ctx *fiber.Ctx) error {
	return ctx.Render("pages/about", fiber.Map{})
}
