package health

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

// New creates a new middleware handler
func New(config ...Config) fiber.Handler {
	// Set default config
	cfg := configDefault(config...)

	return func(c *fiber.Ctx) error {
		path := c.Path()

		if !strings.HasPrefix(path, cfg.DefaultPath) {
			return c.Next()
		}

		return c.SendStatus(fiber.StatusOK)
	}
}
