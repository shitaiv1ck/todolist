package tasks_transport

import (
	"time"

	"github.com/shitaiv1ck/todolist/internal/core/domains"
)

type CreateTaskRequest struct {
	Title       string  `json:"title" validate:"required"`
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

type GetTaskResponse struct {
	ID          int        `json:"id"`
	UserID      int        `json:"user_id"`
	Title       string     `json:"title"`
	Description *string    `json:"description"`
	Completed   bool       `json:"completed"`
	CreatedAt   time.Time  `json:"created_at"`
	CompletedAt *time.Time `json:"completed_at"`
}

type PatchTaskRequest struct {
	Title       domains.Nullable[string] `json:"title"`
	Description domains.Nullable[string] `json:"description"`
	Completed   domains.Nullable[bool]   `json:"completed"`
}

type PatchTaskResponse struct {
	ID          int        `json:"id"`
	UserID      int        `json:"user_id"`
	Title       string     `json:"title"`
	Description *string    `json:"description"`
	Completed   bool       `json:"completed"`
	CreatedAt   time.Time  `json:"created_at"`
	CompletedAt *time.Time `json:"completed_at"`
}

func parseTasksToResponse(tasks []*domains.Task) []*GetTaskResponse {
	responseTasks := make([]*GetTaskResponse, 0)

	for _, task := range tasks {
		responseTasks = append(responseTasks, (*GetTaskResponse)(task))
	}

	return responseTasks
}
