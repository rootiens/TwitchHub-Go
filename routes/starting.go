package routes

import (
	"TwitchHub/controllers"

	"github.com/gofiber/fiber/v2"
)

func StartingRoutes(app fiber.Router) {
	route := app.Group("/startings")
	startingController := &controllers.StartingController{}

	route.Get("/", startingController.Index)
	route.Post("/new", startingController.NewStarting)
}
