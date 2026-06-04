package domains

import (
	"fmt"
	"time"

	core_errors "github.com/shitaiv1ck/todolist/internal/core/errors"
)

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

func (t *Task) CompletionTime() *time.Duration {
	if !t.Completed {
		return nil
	}

	if t.CompletedAt == nil {
		return nil
	}

	completionTime := t.CompletedAt.Sub(t.CreatedAt)

	return &completionTime
}

type TaskPatch struct {
	ID          int
	UserID      int
	Title       Nullable[string]
	Description Nullable[string]
	Completed   Nullable[bool]
}

func NewTaskPatch(
	id int,
	userID int,
	title Nullable[string],
	description Nullable[string],
	completed Nullable[bool],
) *TaskPatch {
	return &TaskPatch{
		ID:          id,
		UserID:      userID,
		Title:       title,
		Description: description,
		Completed:   completed,
	}
}

func (t *TaskPatch) Validate() error {
	if t.Title.Set && t.Title.Value == nil {
		return fmt.Errorf("title can't be null: %w", core_errors.ErrInvalidArgument)
	}

	if t.Completed.Set && t.Completed.Value == nil {
		return fmt.Errorf("completed can't be null: %w", core_errors.ErrInvalidArgument)
	}

	return nil
}

func (t *Task) ApplyPatch(patch *TaskPatch) error {
	if err := patch.Validate(); err != nil {
		return err
	}

	temp := *t

	if patch.Title.Set {
		temp.Title = *patch.Title.Value
	}

	if patch.Description.Set {
		temp.Description = patch.Description.Value
	}

	if patch.Completed.Set {
		completed := *patch.Completed.Value

		temp.Completed = completed
		if completed {
			now := time.Now()
			temp.CompletedAt = &now
		} else {
			temp.CompletedAt = nil
		}
	}

	*t = temp

	return nil
}
