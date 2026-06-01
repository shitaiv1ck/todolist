package tasks_transport

import (
	"net/http"

	"github.com/shitaiv1ck/todolist/internal/core/domains"
	core_errors "github.com/shitaiv1ck/todolist/internal/core/errors"
	core_logger "github.com/shitaiv1ck/todolist/internal/core/logger"
	core_request "github.com/shitaiv1ck/todolist/internal/core/transport/request"
	core_response "github.com/shitaiv1ck/todolist/internal/core/transport/response"
)

type TasksTransport struct {
	service TasksService
}

type TasksService interface {
	CreateTask(task *domains.Task) (*domains.Task, error)
}

func NewTransport(service TasksService) *TasksTransport {
	return &TasksTransport{
		service: service,
	}
}

func (t *TasksTransport) CreateTaskHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := core_logger.FromContext(r.Context())
		responseHandler := core_response.NewResponseHandler(w)

		logger.Debug("invoke CreateTask handler")

		userID := r.Context().Value("user_id")
		if userID == nil {
			responseHandler.ErrorResponse("authenticate user", core_errors.ErrUnautorize)

			return
		}

		var request CreateTaskRequest
		if err := core_request.DecodeAndValidate(r, &request); err != nil {
			responseHandler.ErrorResponse("decode and validate", err)

			return
		}

		task := domains.NewUninitializedTask(
			userID.(int),
			request.Title,
			request.Description,
		)

		createdTask, err := t.service.CreateTask(task)
		if err != nil {
			responseHandler.ErrorResponse("create task", err)

			return
		}

		response := CreateTaskResponse{
			ID:          createdTask.ID,
			UserID:      createdTask.UserID,
			Title:       createdTask.Title,
			Description: createdTask.Description,
			Completed:   createdTask.Completed,
			CreatedAt:   createdTask.CreatedAt,
			CompletedAt: createdTask.CompletedAt,
		}

		responseHandler.JsonResponse(response, http.StatusCreated)
	}
}
