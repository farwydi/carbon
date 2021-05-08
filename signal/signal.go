package signal

import (
	"context"
	"github.com/farwydi/carbon/server"

	"github.com/drone/signal"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

func New(ctx context.Context, logger *zap.Logger) *Signal {
	return &Signal{
		ctx:    signal.WithContext(ctx),
		logger: logger,
	}
}

type Signal struct {
	ctx    context.Context
	eg     errgroup.Group
	logger *zap.Logger
}

func (s *Signal) Add(f func(ctx context.Context) error) {
	s.eg.Go(func() error {
		return f(s.ctx)
	})
}

func (s *Signal) AddIf(condition bool, f func(ctx context.Context) error) {
	if condition {
		s.eg.Go(func() error {
			return f(s.ctx)
		})
	}
}

func (s *Signal) AddServer(srv server.Server) {
	s.eg.Go(func() error {
		return srv.Run(s.ctx)
	})
}

func (s *Signal) AddServerIf(condition bool, srv server.Server) {
	if condition {
		s.eg.Go(func() error {
			return srv.Run(s.ctx)
		})
	}
}

func (s *Signal) Run() {
	if err := s.eg.Wait(); err != nil {
		s.logger.Fatal("program terminated",
			zap.Error(err))
	}
}
