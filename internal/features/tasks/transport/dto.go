package tasks_transport

import "time"

type CreateTaskRequest struct {
	Title       string  `json:"title"`
	Description *string `json:"description"`
}

type CreateTaskResponse struct {
	ID          int        `json:"id"`
	UserID      int        `json:"user_id"`
	Title       string     `json:"title"`
	Description *string    `json:"description"`
	Completed   bool       `json:"completed"`
	CreatedAt   time.Time  `json:"created_at"`
	CompletedAt *time.Time `json:"completed_at"`
}
