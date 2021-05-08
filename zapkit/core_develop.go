package zapkit

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var DefaultDevelopmentCore CoreMakeFunc = func(cfg Config) (zapcore.Core, error) {
	encoder := zap.NewDevelopmentEncoderConfig()
	encoder.EncodeTime = zapcore.ISO8601TimeEncoder
	encoder.EncodeLevel = zapcore.CapitalColorLevelEncoder

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoder),
		zapcore.AddSync(os.Stdout),
		zap.NewAtomicLevelAt(zap.DebugLevel),
	)

	return core, nil
}
