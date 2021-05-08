package mix

import (
	pprof2 "github.com/farwydi/carbon/pkg/pprof"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/google/uuid"
)

// NewStdMix just fiber and carbon.
// Creates a customized fiber.App.
// metricsLabelsModif maybe nil
func NewStdMix(config ...Config) *fiber.App {
	// Set default config
	cfg := configDefault(config...)

	r := fiber.New(cfg.FiberConfig)

	r.Use(requestid.New(requestid.Config{
		Generator: func() string {
			return uuid.New().String()
		},
		ContextKey: cfg.RequestIDCtx,
	}))

	r.Use(cfg.LoggerMiddalware)

	if !cfg.DisableHealth {
		r.Use(cfg.HeathCheckMiddalware)
	}

	if !cfg.DisablePProf {
		r.Use(cfg.PProfMiddalware)
	}

	return r
}

func NewTechMix() *fiber.App {
	r := fiber.New()

	r.Use(pprof2.New())

	return r
}
