package gormkit

import (
	"time"

	"github.com/pkg/errors"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewGorm(cfg Config) (*gorm.DB, func(), error) {
	dbConfig := &gorm.Config{
		Logger: logger.New(
			zap.NewStdLog(cfg.Logger),
			logger.Config{
				// Slow SQL threshold
				SlowThreshold: time.Second,
				// Log level
				LogLevel: logger.Info,
				// Disable color
				Colorful: cfg.Colorful,
			},
		),
	}

	if cfg.DisableLogger {
		dbConfig.Logger.LogMode(logger.Silent)
	}

	ds, err := cfg.ToString()
	if err != nil {
		return nil, nil, errors.Wrap(
			errors.WithMessage(err, "config to source"), "new gorm")
	}

	db, err := gorm.Open(postgres.Open(ds), dbConfig)
	if err != nil {
		return nil, nil, errors.Wrap(
			errors.WithMessage(err, "gorm Open (postgres)"), "new gorm")
	}

	dd, err := db.DB()
	if err != nil {
		return nil, nil, errors.Wrap(
			errors.WithMessage(err, "take sql.DB"), "new gorm")
	}

	cleanup := func() {
		if err := dd.Close(); err != nil {
			cfg.Logger.Error("Fail close database content",
				zap.String("link", ds),
				zap.Error(err),
			)
		}
	}

	dd.SetMaxOpenConns(cfg.MaxOpenConns)
	dd.SetMaxIdleConns(cfg.MaxIdleConns)
	dd.SetConnMaxLifetime(cfg.ConnMaxLifetime)

	return db, cleanup, nil
}
