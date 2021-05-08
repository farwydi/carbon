package gormkit

import (
	"fmt"
	"go.uber.org/zap"
	"time"

	"github.com/pkg/errors"
)

var (
	ErrSetupHost = errors.New("setup host in database config")
	ErrSetupName = errors.New("setup name in database config")
)

type Config struct {
	Name     string
	User     string `default:"postgres"`
	Password string `default:""`
	Host     string

	MaxOpenConns    int           `yaml:"max_open_conns" default:"0"`
	MaxIdleConns    int           `yaml:"max_idle_conns" default:"2"`
	ConnMaxLifetime time.Duration `yaml:"conn_max_lifetime" default:"1m"`

	Colorful      bool
	DisableLogger bool
	Logger        *zap.Logger
}

func (ds Config) ToString() (string, error) {
	if ds.Host == "" {
		return "", ErrSetupHost
	}

	if ds.Name == "" {
		return "", ErrSetupName
	}

	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		ds.Host,
		ds.User,
		ds.Password,
		ds.Name,
	), nil
}
