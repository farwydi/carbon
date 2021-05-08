package zapkit

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(config ...Config) (*zap.Logger, error) {
	// Set default config
	cfg := configDefault(config...)

	var zapCores []zapcore.Core

	for _, coreFunc := range cfg.Cores {
		core, err := coreFunc(cfg)
		if err != nil {
			return nil, err
		}

		zapCores = append(zapCores, core)
	}

	logger := zap.New(zapcore.NewTee(zapCores...))

	var fildes []zap.Field

	if cfg.ProjectName != "" {
		fildes = append(fildes, zap.String("service.name", cfg.ProjectName))
	}
	if cfg.ProectVersion != "" {
		fildes = append(fildes, zap.String("package.install_scope", cfg.ProjectScope))
	}
	if cfg.ProjectScope != "" {
		fildes = append(fildes, zap.String("package.version", cfg.ProectVersion))
	}

	return logger.With(fildes...), nil
}

func RegisterLogger(logger *zap.Logger) {
	// Регистрация логера для доступа по zap.L()
	_ = zap.ReplaceGlobals(logger)
	// Перенаправляет stdlog в zap
	_ = zap.RedirectStdLog(logger)
}
