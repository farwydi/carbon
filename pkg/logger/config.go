package logger

import "go.uber.org/zap"

// Config defines the config for middleware.
type Config struct {
	Logger       *zap.Logger
	SkipPaths    []string
	RequestIDCtx string

	cacheSkipPath map[string]struct{}
}

func (c Config) IsSkipPath(path string) bool {
	_, found := c.cacheSkipPath[path]
	return found
}

// ConfigDefault is the default config
var ConfigDefault = Config{}

// Helper function to set default values
func configDefault(config ...Config) Config {
	// Return default config if nothing provided
	if len(config) < 1 {
		return ConfigDefault
	}

	// Override default config
	cfg := config[0]

	if len(cfg.SkipPaths) > 0 {
		cfg.cacheSkipPath = map[string]struct{}{}
		for _, skipPath := range cfg.SkipPaths {
			cfg.cacheSkipPath[skipPath] = struct{}{}
		}
	}

	return cfg
}
