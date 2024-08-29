// aut controller for service to client
package controller

import (
	"github.com/gin-gonic/gin"
)

// AuthController is the controller for handling authentication requests
type AuthController struct {
	authService *auth.AuthService
}

// NewAuthController creates a new instance of the AuthController struct
func NewAuthController(authService *auth.AuthService) *AuthController {
	return &AuthController{authService}
}

// Login handles the POST /login route and login a new user with the provided credentials
func (c *AuthController) Login(ctx *gin.Context) {
	// Get the email and password from the request body
	var registrationData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := ctx.ShouldBindJSON(&registrationData); err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	if registrationData.Email == "" || registrationData.Password == "" {
		ctx.AbortWithStatusJSON(400, gin.H{"error": "Email and password are required"})
		return
	}

	// Register the new user and get a custom token for the user
	customToken, err := c.authService.Login(registrationData.Email, registrationData.Password)
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	// Return the custom token to the client
	ctx.JSON(200, gin.H{"token": customToken})
}

// Register handles the POST /register route and creates a new user with the provided credentials
func (c *AuthController) Register(ctx *gin.Context) {
	// Get the email and password from the request body
	var registrationData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := ctx.ShouldBindJSON(&registrationData); err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	if registrationData.Email == "" || registrationData.Password == "" {
		ctx.AbortWithStatusJSON(400, gin.H{"error": "Email and password are required"})
		return
	}

	// Register the new user and get a custom token for the user
	customToken, err := c.authService.Register(registrationData.Email, registrationData.Password)
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	// Return the custom token to the client
	ctx.JSON(200, gin.H{"token": customToken})
}

// package controllers

// import (
// 	"strconv"
// 	"time"

// 	"github.com/TanishkN/xunami_mobile/database"
// 	"github.com/TanishkN/xunami_mobile/models"
// 	"github.com/dgrijalva/jwt-go"
// 	"github.com/gofiber/fiber/v3"
// 	"golang.org/x/crypto/bcrypt"
// )

// const SecretKey = "secret"

// // Register handles user registration
// func Register(c *fiber.Ctx) error {
// 	var data map[string]string

// 	password, err := bcrypt.GenerateFromPassword([]byte(data["password"]), bcrypt.DefaultCost)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusInternalServerError, "Failed to hash password")
// 	}

// 	user := models.User{
// 		Name:     data["name"],
// 		Email:    data["email"],
// 		Password: string(password),
// 	}

// 	if err := database.DB.Create(&user).Error; err != nil {
// 		return fiber.NewError(fiber.StatusInternalServerError, "Failed to create user")
// 	}

// 	return c.JSON(fiber.Map{
// 		"message": "User registered successfully",
// 		"user":    user,
// 	})
// }

// // Login handles user login
// func Login(c *fiber.Ctx) error {
// 	var data map[string]string

// 	if err := c.BodyParser(&data); err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, "Invalid request data")
// 	}

// 	var user models.User

// 	if err := database.DB.Where("email = ?", data["email"]).First(&user).Error; err != nil {
// 		return fiber.NewError(fiber.StatusNotFound, "User not found")
// 	}

// 	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data["password"])); err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, "Incorrect password")
// 	}

// 	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
// 		Issuer:    strconv.Itoa(int(user.ID)),
// 		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // 1 day
// 	})

// 	token, err := claims.SignedString([]byte(SecretKey))
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusInternalServerError, "Could not login")
// 	}

// 	cookie := fiber.Cookie{
// 		Name:     "jwt",
// 		Value:    token,
// 		Expires:  time.Now().Add(time.Hour * 24),
// 		HTTPOnly: true,
// 	}

// 	c.Cookie(cookie)

// 	return c.JSON(fiber.Map{
// 		"message": "Login successful",
// 	})
// }

// // User returns the current logged-in user's information
// func User(c *fiber.Ctx) error {
// 	cookie := c.Cookies("jwt")

// 	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
// 		return []byte(SecretKey), nil
// 	})

// 	if err != nil {
// 		return fiber.NewError(fiber.StatusUnauthorized, "Unauthenticated")
// 	}

// 	claims := token.Claims.(*jwt.StandardClaims)

// 	var user models.User

// 	if err := database.DB.Where("id = ?", claims.Issuer).First(&user).Error; err != nil {
// 		return fiber.NewError(fiber.StatusNotFound, "User not found")
// 	}

// 	return c.JSON(user)
// }

// // Logout handles user logout by clearing the JWT cookie
// func Logout(c *fiber.Ctx) error {
// 	cookie := fiber.Cookie{
// 		Name:     "jwt",
// 		Value:    "",
// 		Expires:  time.Now().Add(-time.Hour),
// 		HTTPOnly: true,
// 	}

// 	c.Cookie(cookie)

// 	return c.JSON(fiber.Map{
// 		"message": "Logout successful",
// 	})
// }
