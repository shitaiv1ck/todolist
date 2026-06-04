package statistics_service

import (
	"fmt"
	"time"

	"github.com/shitaiv1ck/todolist/internal/core/domains"
)

type StatisticsService struct {
	repository TasksRepository
}

type TasksRepository interface {
	GetTasks(userID int) ([]*domains.Task, error)
}

func NewService(rep TasksRepository) *StatisticsService {
	return &StatisticsService{
		repository: rep,
	}
}

func (s *StatisticsService) GetStatistics(userID int) (*domains.Statistics, error) {
	tasks, err := s.repository.GetTasks(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get tasks")
	}

	statistics := calculateStatistics(tasks)

	return statistics, nil
}

func calculateStatistics(tasks []*domains.Task) *domains.Statistics {
	if len(tasks) == 0 {
		return domains.NewStatistics(
			0,
			0,
			nil,
			nil,
		)
	}

	totalTasks := len(tasks)
	totalCompleted := 0
	var totalCompletedTime time.Duration

	for _, task := range tasks {
		if task.Completed {
			totalCompleted++

			complitionTime := task.CompletionTime()
			if complitionTime != nil {
				totalCompletedTime += *complitionTime
			}
		}
	}

	totalCompletedRate := float64(totalCompleted) / float64(totalTasks) * 100

	var avgCompletedTime *time.Duration
	if totalCompletedTime > 0 && totalCompleted != 0 {
		avgTime := totalCompletedTime / time.Duration(totalCompleted)
		avgCompletedTime = &avgTime
	}

	return domains.NewStatistics(
		totalTasks,
		totalCompleted,
		&totalCompletedRate,
		avgCompletedTime,
	)
}
