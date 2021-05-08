package carbon

import (
	"context"
	"net/http"
	"time"

	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

func NewHTTPServer(addr string, logger *zap.Logger, router http.Handler) Server {
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

func (t *httpServer) Run(ctx context.Context) error {
	var g errgroup.Group
	g.Go(func() error {
		<-ctx.Done()

		t.logger.Info("Shutting down http server...")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		return t.srv.Shutdown(ctx)
	})
	g.Go(func() error {
		t.logger.Info("Starting http server",
			zap.String("addr", t.addr))
		if err := t.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			return err
		}
		return nil
	})
	return g.Wait()
}
