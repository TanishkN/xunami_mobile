package main

import (
	"github.com/gofiber/fiber/v2"
)

func statusHandler(c *fiber.Ctx, err error, code int) error {
	c.Status(code)
	return c.JSON(fiber.Map{
		"error":   true,
		"message": err.Error(),
	})
}
