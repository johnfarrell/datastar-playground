package router

import (
	"context"
	"go.uber.org/zap"
	"net/http"
)

type Router struct {
	logger *zap.Logger
}

func New(logger *zap.Logger) *Router {
	return &Router{
		logger: logger.Named("router"),
	}
}

func (s *Router) Routes(ctx context.Context) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/{name}", handleIndex)

	return mux
}
