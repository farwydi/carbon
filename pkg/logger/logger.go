package logger

import (
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

// New creates a new middleware handler
func New(config ...Config) fiber.Handler {
	// Set default config
	cfg := configDefault(config...)

	// Set variables
	var (
		once       sync.Once
		errHandler fiber.ErrorHandler
	)

	return func(c *fiber.Ctx) error {
		path := c.Path()

		// Log only when path is not being skipped
		if cfg.IsSkipPath(path) {
			return c.Next()
		}

		// Set error handler once
		once.Do(func() {
			// override error handler
			errHandler = c.App().Config().ErrorHandler
		})

		var start, stop time.Time

		// Set latency start time
		start = time.Now()

		// Handle request, store err for logging
		chainErr := c.Next()

		// Manually call error handler
		if chainErr != nil {
			if err := errHandler(c, chainErr); err != nil {
				_ = c.SendStatus(fiber.StatusInternalServerError)
			}
		}

		// Set latency stop time
		stop = time.Now()

		requestId, _ := c.Locals(cfg.RequestIDCtx).(string)
		resp := c.Response()

		logger := cfg.Logger.With(
			zap.Duration("event.duration", stop.Sub(start).Round(time.Millisecond)),
			zap.String("client.ip", c.IP()),
			zap.String("http.request.method", c.Method()),
			zap.Int("http.response.status_code", resp.StatusCode()),
			zap.Int("http.response.bytes", len(resp.Body())),
			zap.String("url.original", c.OriginalURL()),
			zap.String("trace.id", requestId),
		)

		if chainErr != nil {
			logger.Error("Error in handler",
				zap.Error(chainErr))
			return nil
		}

		logger.Info("Access")
		return nil
	}
}
