package zapkit

// Config defines the config for middleware.
type Config struct {
	ProjectName   string
	ProectVersion string
	ProjectScope  string

	Cores       []CoreMakeFunc
	SpecConfigs []interface{}
}

// ConfigDefault is the default config
var ConfigDefault = Config{
	Cores: []CoreMakeFunc{
		DefaultDevelopmentCore,
	},
}

// Helper function to set default values
func configDefault(config ...Config) Config {
	// Return default config if nothing provided
	if len(config) < 1 {
		return ConfigDefault
	}

	// Override default config
	cfg := config[0]

	return cfg
}
