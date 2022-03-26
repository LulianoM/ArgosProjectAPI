package main

import (
	"github.com/LulianoM/ArgosProjectAPI/src/database"
	"github.com/LulianoM/ArgosProjectAPI/src/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()
	database.SetupRedis()
	database.SetupCacheChannel()

	app := fiber.New()
	routes.Setup(app)

	app.Listen(":8000")
}
