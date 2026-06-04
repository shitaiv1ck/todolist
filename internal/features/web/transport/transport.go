package web_transport

import (
	"net/http"

	core_logger "github.com/shitaiv1ck/todolist/internal/core/logger"
	core_response "github.com/shitaiv1ck/todolist/internal/core/transport/response"
)

type WebTransport struct {
	service WebService
}

type WebService interface {
	GetMainPage() ([]byte, error)
}

func NewTransport(service WebService) *WebTransport {
	return &WebTransport{
		service: service,
	}
}

func (wt *WebTransport) GetMainPageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := core_logger.FromContext(r.Context())
		responseHandler := core_response.NewResponseHandler(w)

		logger.Debug("invoke GetMainPage handler")

		html, err := wt.service.GetMainPage()
		if err != nil {
			responseHandler.ErrorResponse("failed to get index.html", err)

			return
		}

		responseHandler.HtmlResponse(html)
	}
}
