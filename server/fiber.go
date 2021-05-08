package server

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

func NewFiber(addr string, logger *zap.Logger, app *fiber.App) Server {
	return &fiberServer{
		addr:   addr,
		app:    app,
		logger: logger,
	}
}

type fiberServer struct {
	addr   string
	app    *fiber.App
	logger *zap.Logger
}

func (s *fiberServer) Run(ctx context.Context) error {
	var g errgroup.Group
	g.Go(func() error {
		<-ctx.Done()

		s.logger.Info("Shutting down fiber server...")

		return s.app.Shutdown()
	})
	g.Go(func() error {
		s.logger.Info("Starting fiber server",
			zap.String("addr", s.addr))
		return s.app.Listen(s.addr)
	})
	return g.Wait()
}
