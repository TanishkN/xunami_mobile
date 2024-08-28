package main

import (
	"./database"
	"./routes"
	"github.com/gofiber/fiber/v2"
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

/*package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	"xunami_mobile/webserver"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

type config struct {
	port int
	env  string
}

type application struct {
	config config
	logger *log.Logger
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "Server port to listen on")
	flag.StringVar(&cfg.env, "env", "development", "Application environment (development|production)")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	engine := html.New("./src", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Define routes
	app.Post("/register", webserver.Register)
	app.Post("/login", webserver.Login)
	app.Get("/user", webserver.User)
	app.Post("/logout", webserver.Logout)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("Starting server on port %d in %s mode", cfg.port, cfg.env)
	err := srv.ListenAndServe()
	if err != nil {
		logger.Fatalf("Could not start server: %v", err)
	}
}
*/
