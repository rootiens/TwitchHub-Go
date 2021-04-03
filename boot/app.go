package boot

import (
	"github.com/rootiens/TwitchHub-Go/routes"
	"github.com/rootiens/TwitchHub-Go/utils"

	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func BootApp() *fiber.App {

	err := godotenv.Load()
	utils.ErrorHandler(err, "env not imported")

	app := fiber.New()

	api := app.Group("/api", cors.New())
	// startings.StartingRoutes(api)
	routes.LiteratureRoutes(api)

	return app

}
