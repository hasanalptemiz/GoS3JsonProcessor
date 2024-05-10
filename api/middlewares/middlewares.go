package middlewares

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/golangcimri/api/globals"
)

func Fiber(m *fiber.App) {
	m.Use(
		cors.New(),
		helmet.New(),
	)
}

func ApiTokenProtected(c *fiber.Ctx) error {
	authHeader := c.Get("Api-Token")
	if authHeader == "" {
		fmt.Println("Api-Token header not provided")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	if authHeader != globals.Variables.ApiToken {
		fmt.Println("Api-Token header is invalid")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	return c.Next()
}
