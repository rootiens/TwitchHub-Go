package routes

import (
	"github.com/rootiens/TwitchHub-Go/controllers"

	"github.com/gofiber/fiber/v2"
)

func LiteratureRoutes(app fiber.Router) {
	route := app.Group("/literatures")
	literaturesController := &controllers.LiteratureController{}

	route.Get("/", literaturesController.Index)
	route.Get("/id/:ID", literaturesController.GetLiterature)
	route.Post("/create", literaturesController.Create)
}
