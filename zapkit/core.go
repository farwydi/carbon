package zapkit

import (
	"go.uber.org/zap/zapcore"
)

type CoreMakeFunc func(cfg Config) (zapcore.Core, error)

type CoreBuildFunc func() CoreMakeFunc

func CoreBuilder(initFuncs ...CoreBuildFunc) []CoreMakeFunc {
	var cores []CoreMakeFunc
	for _, initFunc := range initFuncs {
		if f := initFunc(); f != nil {
			cores = append(cores, f)
		}
	}
	return cores
}
