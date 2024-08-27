package main

import (
	"your_project/controllers"

	"github.com/gofiber/fiber/v2"
)

func initializeRoutes(app *fiber.App) {
	// Define your routes
	app.Post("/register", controllers.Register)
	app.Post("/login", controllers.Login)
	app.Get("/user", controllers.User)
	app.Post("/logout", controllers.Logout)
}
