package main

import (
	"GoHtmxGL/app/pages"
	"GoHtmxGL/models"
	"GoHtmxGL/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"log"
)

func init() {
	// Initialize the database (gorm)
	models.InitDB()
}
func main() {

	// initialize fiber template engine
	engine := html.New("./views", ".html")
	engine.Reload(true) // Disable this in production
	engine.AddFunc("getCssAsset", utils.GetCssAsset)

	app := fiber.New(fiber.Config{
		Views: engine,
		// Views Layout is the global layout for all template render until override on Render function.
		ViewsLayout: "layouts/main",
	})
	app.Static("/public", "./public")

	//app.Get("/", func(c *fiber.Ctx) error {
	//	return c.Status(fiber.StatusOK).Render("pages/index", fiber.Map{})
	//	//return c.Render("layouts/main", fiber.Map{})
	//})
	//app.Get("/auth", func(c *fiber.Ctx) error {
	//	return c.Status(fiber.StatusOK).Render("pages/index", fiber.Map{}, "layouts/authed")
	//	//return c.Render("layouts/main", fiber.Map{})
	//})

	pages.NavPageRoutes(app)
	pages.Components(app)
	log.Fatal(app.Listen(":8080"))
}
