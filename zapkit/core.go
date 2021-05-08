package zapkit

import (
	"go.uber.org/zap/zapcore"
)

type CoreMakeFunc func(cfg Config) (zapcore.Core, error)
