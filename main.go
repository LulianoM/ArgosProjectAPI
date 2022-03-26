package main

import (
	"github.com/LulianoM/ArgosProjectAPI/src/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	routes.Setup(app)

	app.Listen(":3000")
}
