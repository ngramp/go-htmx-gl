package main

import (
	"GoHtmxGL/app/pages"
	"github.com/gofiber/fiber/v2"
)

func MenuRoutes(router fiber.Router) {
	nav := router.Group("/")
	nav.Get("/", pages.Home)
	nav.Get("/about", pages.About)
	nav.Get("/contact", pages.Contact)
	nav.Get("/donate", pages.Donate)
}

func ArticleRoutes(router fiber.Router) {
	articles := router.Group("/articles")
	articles.Get("/", pages.GetArticles)
	articles.Get("/:id", pages.GetArticle)
	articles.Post("/", pages.CreateArticle)
	articles.Put("/:id", pages.UpdateArticle)
	articles.Delete("/:id", pages.DeleteArticle)
}
func UserRoutes(router fiber.Router) {
	users := router.Group("/users")
	users.Get("/", pages.GetUsers)
	users.Get("/:id", pages.GetUser)
	users.Post("/", pages.CreateUser)
	users.Put("/:id", pages.UpdateUser)
	users.Delete("/:id", pages.DeleteUser)
}
func Components(t fiber.Router) {
}
