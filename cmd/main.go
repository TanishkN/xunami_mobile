package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/TanishkN/XUNAMI.IO/webserver"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html" // Import the correct package for HTML templates
)

// Config struct to hold application configuration
type config struct {
	port int
	env  string
}

// Application struct to hold application-wide dependencies
type application struct {
	config config
	logger *log.Logger
}

func main() {
	var cfg config

	// Parse command-line flags for configuration
	flag.IntVar(&cfg.port, "port", 4000, "Server port to listen on")
	flag.StringVar(&cfg.env, "env", "development", "Application environment (development|production)")
	flag.Parse()

	// Set up logger to log to stdout with date and time
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// Set up Fiber with an HTML template engine
	engine := html.New("./src", ".js") // Adjust the file extension if needed
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Define the routing with correct HTTP methods and endpoints
	app.Post("/register", webserver.Register)
	app.Post("/login", webserver.Login)
	app.Get("/user", webserver.User)
	app.Post("/logout", webserver.Logout)

	// Set up the HTTP server with proper timeouts
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	// Start the server
	logger.Printf("Starting server on port %d in %s mode", cfg.port, cfg.env)
	err := srv.ListenAndServe()
	if err != nil {
		logger.Fatalf("Could not start server: %v", err)
	}
}
