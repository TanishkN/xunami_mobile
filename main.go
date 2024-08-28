package main

import (
	"github.com/TanishkN/xunami_mobile/database"
	"github.com/TanishkN/xunami_mobile/routes"
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.Connect()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.Setup(app)

	app.Listen(":8000")
}
