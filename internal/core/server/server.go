package core_server

import (
	"context"
	"errors"
	"net/http"

	core_logger "github.com/shitaiv1ck/todolist/internal/core/logger"
	"go.uber.org/zap"
)

type Server struct {
	config Config
	router *http.ServeMux

	logger *core_logger.Logger
}

func NewServer(router *http.ServeMux, logger *core_logger.Logger) *Server {
	return &Server{
		config: NewConfigMust(),
		router: router,
		logger: logger,
	}
}

func (s *Server) Run(ctx context.Context) error {
	httpServer := &http.Server{
		Addr:    s.config.Addres,
		Handler: s.router,
	}

	ch := make(chan error, 1)

	go func() {
		defer close(ch)
		s.logger.Debug("start the server", zap.String("addres", s.config.Addres))

		err := httpServer.ListenAndServe()
		if !errors.Is(err, http.ErrServerClosed) {
			ch <- err
		}
	}()

	select {
	case err := <-ch:
		s.logger.Error("server's work", zap.Error(err))
	case <-ctx.Done():
		httpServer.Close()

		s.logger.Info("the server was closed")
	}

	return nil
}
