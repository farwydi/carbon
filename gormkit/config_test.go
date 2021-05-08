package gormkit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig_ToString_Errs(t *testing.T) {
	var (
		ds     Config
		err    error
		source string
	)

	ds = Config{}
	_, err = ds.ToString()
	assert.Equal(t, ErrSetupHost, err)

	ds = Config{Host: "dbhost"}
	_, err = ds.ToString()
	assert.Equal(t, ErrSetupName, err)

	ds = Config{Host: "dbhost", Name: "dbname"}
	source, err = ds.ToString()
	assert.NoError(t, err)
	assert.Equal(t, "host=dbhost user= password= dbname=dbname sslmode=disable", source)
}
