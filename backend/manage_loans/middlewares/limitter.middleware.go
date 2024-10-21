package middlewares

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func LimitterMiddleWare() func(*fiber.Ctx) error {
	return limiter.New(limiter.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.IP() == "127.0.0.1" || c.IP() == "localhost"
		},
		Max:        30,
		Expiration: 30 * time.Second,
	})
}
