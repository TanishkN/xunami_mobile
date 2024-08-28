package routes

import (
	"../controllers"
	"github.com/gofiber/fiber"
)

func Setup(app *fiber.App) {
	//being created from the controllers class
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Get("/api/user", controllers.User)
	app.Post("/api/logout", controllers.Logout)

}

// export GOOGLE_APPLICATION_CREDENTIALS="/desktop/CREATE/xunami_mobile/public/xunami-userbase-firebase-adminsdk-tlmx5-15eeb96930"
