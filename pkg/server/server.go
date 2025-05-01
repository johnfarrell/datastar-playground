package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type Server struct {
	logger *zap.Logger
	server *http.Server
}

func New(l *zap.Logger) *Server {
	server := &http.Server{
		Addr: fmt.Sprintf(":%d", viper.GetInt("port")),
	}
	return &Server{
		logger: l.Named("Server"),
		server: server,
	}
}

// Run initializes and runs the server.
// This blocks
func (s *Server) Run(ctx context.Context, handler http.Handler) error {

	s.server.Handler = handler

	errCh := make(chan error, 1)
	go func() {
		<-ctx.Done()

		s.logger.Info("Server context cancelled")

		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		s.logger.Info("Server shutting down")
		errCh <- s.server.Shutdown(shutdownCtx)
	}()

	s.logger.Info("Starting server", zap.String("address", s.server.Addr))
	if err := s.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		s.logger.Error("Server failed to start", zap.Error(err))
	}

	s.logger.Info("Server shutdown complete")
	if err := <-errCh; err != nil {
		return err
	}
	return nil
}
