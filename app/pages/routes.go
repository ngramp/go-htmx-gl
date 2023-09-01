package pages

import (
	"GoHtmxGL/app/pages/article"
	"GoHtmxGL/app/pages/user"
	"github.com/gofiber/fiber/v2"
)

func NavPageRoutes(router fiber.Router) {
	nav := router.Group("/")
	nav.Get("/", getHome)
	nav.Get("/about", getAbout)
	nav.Get("/contact", getContact)
	nav.Get("/donate", getDonate)
	article.Routes(router)
	user.Routes(router)
}

func Components(t fiber.Router) {
}
