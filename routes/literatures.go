package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rootiens/TwitchHub-Go/controllers"
)

func LiteratureRoutes(app fiber.Router) {
	r := app.Group("/literatures")

	r.Post("/create", controllers.CreateLiterature)
}
