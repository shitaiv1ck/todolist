package domains

import "time"

type Task struct {
	ID          int
	UserID      int
	Title       string
	Description *string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

func NewUninitializedTask(
	usedID int,
	title string,
	description *string,
) *Task {
	return &Task{
		ID:          -1,
		UserID:      usedID,
		Title:       title,
		Description: description,
		Completed:   false,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}
}
