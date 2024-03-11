package main

import (
	"GoHtmxGL/models"
	"GoHtmxGL/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"log"
)

func init() {
	// Initialize the database (gorm)
	//TODO replace with sqlx
	models.InitDB()
}
func main() {

	//r := chi.NewRouter()
	//r.Use(middleware.Logger)
	//r.Get("/", func(w http.ResponseWriter, r *http.Request) {
	//    w.Write([]byte("Hello World!"))
	//})
	//http.ListenAndServe(":3000", r)

	// initialize fiber template engine
	engine := html.New("./views", ".html")
	engine.Reload(true) // Disable this in production
	engine.AddFunc("getCssAsset", utils.GetCssAsset)

	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "layouts/main",
	})
	app.Static("/public", "./public")

	MenuRoutes(app)
	ArticleRoutes(app)
	UserRoutes(app)
	Components(app)
	log.Fatal(app.Listen(":8080"))
}
