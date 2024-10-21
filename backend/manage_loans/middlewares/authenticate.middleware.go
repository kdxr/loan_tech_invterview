package middlewares

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

func Authenticate() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		authorization := c.Get("Authorization")

		splitToken := strings.Split(authorization, " ")
		token := splitToken[len(splitToken)-1]

		// if token != services.Global.TokenC {
		// 	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		// 		"message": fiber.ErrUnauthorized.Message,
		// 	})
		// }

		c.Set("token", token)

		c.Next()

		return nil
	}
}
