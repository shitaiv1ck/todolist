package users_transport

import (
	"net/http"

	"github.com/shitaiv1ck/todolist/internal/core/domains"
	core_logger "github.com/shitaiv1ck/todolist/internal/core/logger"
	core_request "github.com/shitaiv1ck/todolist/internal/core/transport/request"
	core_response "github.com/shitaiv1ck/todolist/internal/core/transport/response"
)

type UsersTransport struct {
	service UsersService
}

type UsersService interface {
	CreateUser(user *domains.User) (*domains.User, error)
}

func NewTransport(service UsersService) *UsersTransport {
	return &UsersTransport{
		service: service,
	}
}

func (ut *UsersTransport) CreateUserHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := core_logger.FromContext(r.Context())
		responseHandler := core_response.NewResponseHandler(w)

		logger.Debug("invoke CreateUser handler")

		var request CreateUserRequest
		if err := core_request.DecodeAndValidate(r, &request); err != nil {
			responseHandler.ErrorResponse("decode and validate", err)

			return
		}

		user := domains.NewUninitializedUser(
			request.Username,
			request.Password,
		)

		createdUser, err := ut.service.CreateUser(user)
		if err != nil {
			responseHandler.ErrorResponse("create user", err)

			return
		}

		response := CreateUserResponse{
			ID:       createdUser.ID,
			Username: createdUser.Username,
		}

		responseHandler.JsonResponse(response, http.StatusCreated)
	}
}
