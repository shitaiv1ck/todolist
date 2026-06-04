package statistics_transport

import (
	"net/http"

	"github.com/shitaiv1ck/todolist/internal/core/domains"
	core_errors "github.com/shitaiv1ck/todolist/internal/core/errors"
	core_logger "github.com/shitaiv1ck/todolist/internal/core/logger"
	core_response "github.com/shitaiv1ck/todolist/internal/core/transport/response"
)

type StatisticsTransport struct {
	service StatisticsService
}

type StatisticsService interface {
	GetStatistics(userID int) (*domains.Statistics, error)
}

func NewTransport(service StatisticsService) *StatisticsTransport {
	return &StatisticsTransport{
		service: service,
	}
}

func (st *StatisticsTransport) GetStatisticsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := core_logger.FromContext(r.Context())
		responseHandler := core_response.NewResponseHandler(w)

		logger.Debug("invoke GetStatistics handler")

		userID := r.Context().Value("user_id")
		if userID == nil {
			responseHandler.ErrorResponse("failed to authentication", core_errors.ErrCookie)

			return
		}

		statistics, err := st.service.GetStatistics(userID.(int))
		if err != nil {
			responseHandler.ErrorResponse("failed to get statistics", err)

			return
		}

		response := stasticsToResponse(statistics)

		responseHandler.JsonResponse(response, http.StatusOK)
	}
}
