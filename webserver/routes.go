package main

import (
	"xunami_mobile/webserver"

	"github.com/gofiber/fiber/v2"
)

func initializeRoutes(app *fiber.App) {
	// Define your routes
	app.Post("/register", webserver.Register)
	app.Post("/login", webserver.Login)
	app.Get("/user", webserver.User)
	app.Post("/logout", webserver.Logout)
}
