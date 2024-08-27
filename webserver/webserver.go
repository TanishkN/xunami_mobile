package webserver

import (
	"os"
	"strconv"
	"time"

	//"your_project/database"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

var SecretKey = os.Getenv("SECRET_KEY")

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

func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid input")
	}

	password, err := hashPassword(data["password"])
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to hash password")
	}

	user := database.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: password,
	}

	if result := database.DB.Create(&user); result.Error != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to create user")
	}

	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid input")
	}

	var user database.User

	if result := database.DB.Where("email = ?", data["email"]).First(&user); result.Error != nil {
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

func User(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil || !token.Valid {
		return fiber.NewError(fiber.StatusUnauthorized, "Unauthenticated")
	}

	claims := token.Claims.(*jwt.StandardClaims)

	var user database.User

	if result := database.DB.Where("id = ?", claims.Issuer).First(&user); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, "User not found")
	}

	return c.JSON(user)
}

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
}
