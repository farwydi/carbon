package logger

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

// go test -run Test_Logger
func Test_LoggerError(t *testing.T) {
	app := fiber.New()

	loggerCore, loggerObserver := observer.New(zap.InfoLevel)
	testLogger := zap.New(loggerCore)

	app.Use(New(Config{
		Logger: testLogger,
	}))

	returnErr := errors.New("some random error")
	app.Get("/", func(c *fiber.Ctx) error {
		return returnErr
	})

	resp, err := app.Test(httptest.NewRequest("GET", "/", nil))
	assert.NoError(t, err)

	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	b, err := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err)
	assert.EqualValues(t, "some random error", b)

	require.Equal(t, 1, loggerObserver.Len())
	all := loggerObserver.TakeAll()

	logLine := all[0]
	assert.Equal(t, zap.ErrorLevel, logLine.Level)

	ctxMap := logLine.ContextMap()
	assert.NotZero(t, ctxMap["error"])
	assert.NotZero(t, ctxMap["client.ip"])
	assert.NotZero(t, ctxMap["http.request.method"])
	assert.NotZero(t, ctxMap["http.response.status_code"])
	assert.NotZero(t, ctxMap["http.response.bytes"])
	assert.NotZero(t, ctxMap["url.original"])

	for _, field := range logLine.Context {
		switch field.Key {
		case "error":
			assert.True(t, field.Equals(zap.Error(returnErr)))
		case "client.ip":
			assert.Equal(t, "0.0.0.0", field.String)
		case "http.request.method":
			assert.Equal(t, http.MethodGet, field.String)
		case "http.response.status_code":
			assert.EqualValues(t, fiber.StatusInternalServerError, field.Integer)
		case "http.response.bytes":
			assert.EqualValues(t, 17, field.Integer)
		case "url.original":
			assert.Equal(t, "/", field.String)
		}
	}
}
