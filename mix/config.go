package mix

import (
	"github.com/farwydi/carbon/pkg/health"
	"github.com/farwydi/carbon/pkg/logger"
	"github.com/farwydi/carbon/pkg/pprof"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type Config struct {
	FiberConfig fiber.Config

	Logger           *zap.Logger
	LoggerConfig     logger.Config
	LoggerMiddalware fiber.Handler

	DisablePProf    bool
	PProfConfig     pprof.Config
	PProfMiddalware fiber.Handler

	DisableHealth        bool
	HeathCheckConfig     health.Config
	HeathCheckMiddalware fiber.Handler

	RequestIDCtx string
}

var ConfigDefault = Config{
	Logger: zap.L(),
	LoggerConfig: logger.Config{
		Logger: zap.L(),
		SkipPaths: []string{
			health.ConfigDefault.DefaultPath,
		},
		RequestIDCtx: "req-id-ctx",
	},
	HeathCheckConfig: health.Config{
		DefaultPath: health.ConfigDefault.DefaultPath,
	},
	PProfConfig: pprof.Config{
		DefaultPath: pprof.ConfigDefault.DefaultPath,
	},
	RequestIDCtx: "req-id-ctx",
}

func configDefault(config ...Config) Config {
	// Return default config if nothing provided
	if len(config) < 1 {
		return ConfigDefault
	}

	cfg := config[0]

	if cfg.RequestIDCtx == "" {
		cfg.RequestIDCtx = ConfigDefault.RequestIDCtx
	}

	if cfg.LoggerMiddalware == nil {
		if cfg.LoggerConfig.Logger == nil {
			cfg.LoggerConfig.Logger = cfg.Logger
		}

		if cfg.LoggerConfig.RequestIDCtx == "" {
			cfg.LoggerConfig.RequestIDCtx = cfg.RequestIDCtx
		}

		cfg.LoggerMiddalware = logger.New(cfg.LoggerConfig)
	}

	if !cfg.DisableHealth && cfg.HeathCheckMiddalware == nil {
		if cfg.HeathCheckConfig.DefaultPath == "" {
			cfg.HeathCheckConfig.DefaultPath = health.ConfigDefault.DefaultPath
		}

		cfg.HeathCheckMiddalware = health.New(cfg.HeathCheckConfig)
	}

	if !cfg.DisablePProf && cfg.PProfMiddalware == nil {
		if cfg.PProfConfig.DefaultPath == "" {
			cfg.PProfConfig.DefaultPath = pprof.ConfigDefault.DefaultPath
		}

		cfg.PProfMiddalware = pprof.New(cfg.PProfConfig)
	}

	return cfg
}
