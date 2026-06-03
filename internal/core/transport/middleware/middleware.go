package core_middleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/shitaiv1ck/todolist/internal/core/domains"
	core_errors "github.com/shitaiv1ck/todolist/internal/core/errors"
	core_logger "github.com/shitaiv1ck/todolist/internal/core/logger"
	core_response "github.com/shitaiv1ck/todolist/internal/core/transport/response"
	"go.uber.org/zap"
)

type Middleware func(http.Handler) http.Handler

var (
	requestID = "X-Request-ID"
	csrfToken = "X-CSRF-Token"
)

type SessionsService interface {
	FindByToken(token string) (*domains.Session, error)
}

func CORS() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			allowedOrigins := map[string]struct{}{
				"http://localhost:8080": {},
			}

			origin := r.Header.Get("Origin")

			if _, ok := allowedOrigins[origin]; ok {
				w.Header().Set("Access-Control-Allow-Origin", origin)
				w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE, OPTIONS")
				w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			}

			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusOK)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

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
			rw := core_response.NewResponseHandler(w)

			log := logger.With(
				zap.String("request_id", r.Header.Get(requestID)),
				zap.String("URL", r.URL.String()),
				zap.String("method", r.Method),
			)

			ctx := context.WithValue(r.Context(), "logger", log)

			log.Debug(">>> incoming http request")

			next.ServeHTTP(rw, r.WithContext(ctx))

			log.Info("<<< done http request", zap.Int("status code", rw.GetStatusCode()))
		})
	}
}

func Authentication(sessionsService SessionsService) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			responseHandler := core_response.NewResponseHandler(w)

			sessionToken, err := r.Cookie("session_token")
			if err != nil {
				responseHandler.ErrorResponse("failed to authentication", core_errors.ErrCookie)

				return
			}

			session, err := sessionsService.FindByToken(sessionToken.Value)
			if err != nil {
				responseHandler.ErrorResponse("failed to authentication", core_errors.ErrCookie)

				return
			}

			ctx := context.WithValue(r.Context(), "user_id", session.UserID)
			ctx = context.WithValue(ctx, "csrf_token", session.CSRFToken)

			next.ServeHTTP(w, r.WithContext(ctx))

		})
	}
}

func ValidateCSRF() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			responseHandler := core_response.NewResponseHandler(w)

			csrfFromReq := r.Header.Get(csrfToken)
			if csrfFromReq == "" {
				responseHandler.ErrorResponse("failed to authentication", core_errors.ErrCookie)

				return
			}

			csrfFromCoockie := r.Context().Value("csrf_token").(string)
			if csrfFromReq != csrfFromCoockie {
				responseHandler.ErrorResponse("failed to authentication", core_errors.ErrCookie)

				return
			}

			next.ServeHTTP(w, r)
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

func ChainCommon(h http.Handler, logger *core_logger.Logger) http.Handler {
	return ChainMiddleware(
		h,
		CORS(),
		RequestID(),
		Logger(logger),
	)
}

func ChainAuthenticated(h http.Handler, sessionsService SessionsService) http.Handler {
	return ChainMiddleware(
		h,
		Authentication(sessionsService),
	)
}

func ChainProtected(h http.Handler, sessionsService SessionsService) http.Handler {
	return ChainMiddleware(
		h,
		Authentication(sessionsService),
		ValidateCSRF(),
	)
}
