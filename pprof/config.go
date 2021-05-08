package pprof

// Config defines the config for middleware.
type Config struct {
	DefaultPath string
}

// ConfigDefault is the default config
var ConfigDefault = Config{
	DefaultPath: "/pprof",
}

// Helper function to set default values
func configDefault(config ...Config) Config {
	// Return default config if nothing provided
	if len(config) < 1 {
		return ConfigDefault
	}

	// Override default config
	cfg := config[0]

	if cfg.DefaultPath == "" {
		cfg.DefaultPath = ConfigDefault.DefaultPath
	}

	return cfg
}
