package tasks_repository

import (
	"github.com/lib/pq"
	"github.com/shitaiv1ck/todolist/internal/core/domains"
	core_errors "github.com/shitaiv1ck/todolist/internal/core/errors"
	core_postgres "github.com/shitaiv1ck/todolist/internal/core/store/postgres"
)

type TasksRepository struct {
	store *core_postgres.Store
}

func NewRepository(store *core_postgres.Store) *TasksRepository {
	return &TasksRepository{
		store: store,
	}
}

func (r *TasksRepository) CreateTask(task *domains.Task) (*domains.Task, error) {
	db := r.store.GetDB()

	query := `
		INSERT INTO todolist.tasks(user_id, title, description)
		VALUES ($1, $2, $3)
		RETURNING *;
	`

	var createdTask domains.Task
	if err := db.QueryRow(
		query,
		task.UserID,
		task.Title,
		task.Description,
	).Scan(
		&createdTask.ID,
		&createdTask.UserID,
		&createdTask.Title,
		&createdTask.Description,
		&createdTask.Completed,
		&createdTask.CreatedAt,
		&createdTask.CompletedAt,
	); err != nil {
		if errPQ, ok := err.(*pq.Error); ok {
			if errPQ.Code == "23505" {
				return nil, core_errors.ErrConflict
			}
		}

		return nil, err
	}

	return &createdTask, nil
}
