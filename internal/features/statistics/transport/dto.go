package statistics_transport

import "github.com/shitaiv1ck/todolist/internal/core/domains"

type GetStatisticsResponse struct {
	TotalTasks         int      `json:"total_tasks"`
	TotalCompleted     int      `json:"total_completed"`
	TotalCompletedRate *float64 `json:"total_completed_rate"`
	AvgCompletedTime   *string  `json:"avg_completed_time"`
}

func stasticsToResponse(statistics *domains.Statistics) GetStatisticsResponse {
	var avgCompletedTime *string
	if statistics.AvgCompletedTime != nil {
		avgTime := statistics.AvgCompletedTime.String()
		avgCompletedTime = &avgTime
	}

	return GetStatisticsResponse{
		TotalTasks:         statistics.TotalTasks,
		TotalCompleted:     statistics.TotalCompleted,
		TotalCompletedRate: statistics.TotalCompletedRate,
		AvgCompletedTime:   avgCompletedTime,
	}
}
