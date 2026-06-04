package sessions_transport

import (
	"net/http"

	"github.com/shitaiv1ck/todolist/internal/core/domains"
	core_errors "github.com/shitaiv1ck/todolist/internal/core/errors"
	core_logger "github.com/shitaiv1ck/todolist/internal/core/logger"
	core_request "github.com/shitaiv1ck/todolist/internal/core/transport/request"
	core_response "github.com/shitaiv1ck/todolist/internal/core/transport/response"
)

type SessionsTransport struct {
	service SessionsService
}

type SessionsService interface {
	Authenticate(username string, password string) (int, error)
	CreateSession(userID int) (*domains.Session, error)
	DeleteByToken(sessionToken string) error
}

func NewTransport(service SessionsService) *SessionsTransport {
	return &SessionsTransport{
		service: service,
	}
}

func (st *SessionsTransport) CreateSessionHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := core_logger.FromContext(r.Context())
		responseHandler := core_response.NewResponseHandler(w)

		logger.Debug("invoke CreateSession handler")

		var request CreateSessionRequest
		if err := core_request.DecodeAndValidate(r, &request); err != nil {
			responseHandler.ErrorResponse("failed to decode and validate", err)

			return
		}

		userID, err := st.service.Authenticate(request.Username, request.Password)
		if err != nil {
			responseHandler.ErrorResponse("failed to authentication", err)

			return
		}

		session, err := st.service.CreateSession(userID)
		if err != nil {
			responseHandler.ErrorResponse("failed to create session", err)

			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:     "session_token",
			Value:    session.SessionToken,
			Expires:  session.ExpiresAt,
			Path:     "/",
			HttpOnly: true,
			SameSite: http.SameSiteLaxMode,
		})

		http.SetCookie(w, &http.Cookie{
			Name:     "csrf_token",
			Value:    session.CSRFToken,
			Expires:  session.ExpiresAt,
			Path:     "/",
			HttpOnly: false,
			SameSite: http.SameSiteLaxMode,
		})

		responseHandler.WriteHeader(http.StatusCreated)
	}
}

func (st *SessionsTransport) DeleteSessionHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := core_logger.FromContext(r.Context())
		responseHandler := core_response.NewResponseHandler(w)

		logger.Debug("invoke DeleteSession handler")

		sessionToken, err := r.Cookie("session_token")
		if err != nil {
			responseHandler.ErrorResponse("failed to authentication", core_errors.ErrCookie)

			return
		}

		if err := st.service.DeleteByToken(sessionToken.Value); err != nil {
			responseHandler.ErrorResponse("failed to delete session", err)
		}

		responseHandler.WriteHeader(http.StatusNoContent)
	}
}
