package server

import (
	"context"
	"net/http"
	"time"

	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

func NewHTTP(addr string, logger *zap.Logger, router http.Handler) Server {
	srv := &http.Server{
		Addr:    addr,
		Handler: router,
	}
	return &httpServer{srv: srv, logger: logger, addr: addr}
}

type httpServer struct {
	addr   string
	srv    *http.Server
	logger *zap.Logger
}

func (s *httpServer) Run(ctx context.Context) error {
	var g errgroup.Group
	g.Go(func() error {
		<-ctx.Done()

		s.logger.Info("Shutting down http server...")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		return s.srv.Shutdown(ctx)
	})
	g.Go(func() error {
		s.logger.Info("Starting http server",
			zap.String("addr", s.addr))
		if err := s.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			return err
		}
		return nil
	})
	return g.Wait()
}
