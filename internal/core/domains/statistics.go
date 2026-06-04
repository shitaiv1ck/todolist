package domains

import "time"

type Statistics struct {
	TotalTasks         int
	TotalCompleted     int
	TotalCompletedRate *float64
	AvgCompletedTime   *time.Duration
}

func NewStatistics(
	totalTasks int,
	totalCompleted int,
	TotalCompletedRate *float64,
	avgCompletedTime *time.Duration,
) *Statistics {
	return &Statistics{
		TotalTasks:         totalTasks,
		TotalCompleted:     totalCompleted,
		TotalCompletedRate: TotalCompletedRate,
		AvgCompletedTime:   avgCompletedTime,
	}
}
