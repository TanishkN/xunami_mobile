package main

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	"github.com/TanishkN/xunami_mobile/cmd/controller"
	"github.com/TanishkN/xunami_mobile/cmd/router"
	db "github.com/TanishkN/xunami_mobile/database"
	"google.golang.org/api/option"
)

func main() {
	// Set up the database connection
	dbConn, err := db.ConnectToDatabase("host=localhost port=5432 user=postgres dbname=postgres password=member sslmode=disable")
	if err != nil {
		log.Fatal("Failed to connect to database")
	}
	defer dbConn.Close()

	// Create a Firebase app instance
	opt := option.WithCredentialsFile("./service-account-key.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("Failed to create Firebase app: %v", err)
	}

	// Create a Firebase auth client instance
	authClient, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalf("Failed to create Firebase auth client: %v", err)
	}

	// Run database migrations
	dbConn.AutoMigrate(&auth.User{})

	// Set up the authentication service and middleware
	authService := &auth.AuthService{
		DB:       dbConn,
		FireAuth: authClient,
	}
	authController := controller.NewAuthController(authService)

	// Set up the HTTP router
	r := router.NewRouter(authController)

	// Start the server
	r.Run(":8000")
}

/*package webserver

import (
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm" // Ensure to import GORM for database interactions
)

var SecretKey = os.Getenv("SECRET_KEY")
var DB *gorm.DB // This should be initialized in your main application

// User represents the user model in the database
type User struct {
	ID       uint `gorm:"primaryKey"`
	Name     string
	Email    string `gorm:"unique"`
	Password []byte
}

// Utility function to hash passwords
func hashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// Utility function to check passwords
func checkPassword(hashedPassword []byte, plainPassword string) error {
	return bcrypt.CompareHashAndPassword(hashedPassword, []byte(plainPassword))
}

// Utility function to generate JWT tokens
func generateJWT(userID uint) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(userID)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	return claims.SignedString([]byte(SecretKey))
}

// Register handles user registration
func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid input")
	}

	password, err := hashPassword(data["password"])
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to hash password")
	}

	user := User{
		Name:     data["name"],
		Email:    data["email"],
		Password: password,
	}

	if result := DB.Create(&user); result.Error != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to create user")
	}

	return c.JSON(user)
}

// Login handles user login
func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid input")
	}

	var user User

	if result := DB.Where("email = ?", data["email"]).First(&user); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, "User not found")
	}

	if err := checkPassword(user.Password, data["password"]); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Incorrect password")
	}

	token, err := generateJWT(user.ID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Could not login")
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

// User retrieves user information
func User(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil || !token.Valid {
		return fiber.NewError(fiber.StatusUnauthorized, "Unauthenticated")
	}

	claims := token.Claims.(*jwt.StandardClaims)

	var user User

	if result := DB.Where("id = ?", claims.Issuer).First(&user); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, "User not found")
	}

	return c.JSON(user)
}

// Logout handles user logout
func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}*/
