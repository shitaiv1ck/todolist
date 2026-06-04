package web_service

import (
	"fmt"
	"os"
	"path"
)

type WebService struct {
	repository WebRepository
}

type WebRepository interface {
	GetFile(path string) ([]byte, error)
}

func NewService(rep WebRepository) *WebService {
	return &WebService{
		repository: rep,
	}
}

func (ws *WebService) GetMainPage() ([]byte, error) {
	htmlFilePath := path.Join(
		os.Getenv("PROJECT_ROOT"),
		"/public/index.html",
	)

	html, err := ws.repository.GetFile(htmlFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to get file: %w", err)
	}

	return html, err
}
