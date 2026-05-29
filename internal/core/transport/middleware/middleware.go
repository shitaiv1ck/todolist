package core_middleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	core_logger "github.com/shitaiv1ck/todolist/internal/core/logger"
	"go.uber.org/zap"
)

type Middleware func(http.Handler) http.Handler

var (
	requestID = "X-Request-ID"
)

func RequestID() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			id := r.Header.Get(requestID)

			if id == "" {
				id = uuid.NewString()
			}

			r.Header.Set(requestID, id)
			w.Header().Set(requestID, id)

			next.ServeHTTP(w, r)
		})
	}
}

func Logger(logger *core_logger.Logger) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log := logger.With(
				zap.String("request_id", r.Header.Get(requestID)),
				zap.String("URL", r.URL.String()),
				zap.String("method", r.Method),
			)

			ctx := context.WithValue(r.Context(), "logger", log)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func ChainMiddleware(h http.Handler, m ...Middleware) http.Handler {
	if len(m) == 0 {
		return h
	}

	for i := len(m) - 1; i >= 0; i-- {
		h = m[i](h)
	}

	return h
}

func PublicMiddleware(h http.Handler, logger *core_logger.Logger) http.Handler {
	return ChainMiddleware(
		h,
		RequestID(),
		Logger(logger),
	)
}
