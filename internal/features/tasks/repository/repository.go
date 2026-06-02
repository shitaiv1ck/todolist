package tasks_repository

import (
	"database/sql"
	"errors"

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

func (r *TasksRepository) FindByUserID(id int, userID int) (*domains.Task, error) {
	db := r.store.GetDB()

	query := `
		SELECT * FROM todolist.tasks
		WHERE id = $1 AND user_id = $2;
	`

	var foundTask domains.Task
	if err := db.QueryRow(
		query,
		id,
		userID,
	).Scan(
		&foundTask.ID,
		&foundTask.UserID,
		&foundTask.Title,
		&foundTask.Description,
		&foundTask.Completed,
		&foundTask.CreatedAt,
		&foundTask.CompletedAt,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, core_errors.ErrNotFound
		}

		return nil, err
	}

	return &foundTask, nil
}

func (r *TasksRepository) UpdateTask(task *domains.Task) (*domains.Task, error) {
	db := r.store.GetDB()

	query := `
		UPDATE todolist.tasks
		SET title = $1, description = $2, completed = $3, completed_at = $4
		WHERE id = $5 AND user_id = $6
		RETURNING *;
	`

	var patchedTask domains.Task
	if err := db.QueryRow(
		query,
		task.Title,
		task.Description,
		task.Completed,
		task.CompletedAt,
		task.ID,
		task.UserID,
	).Scan(
		&patchedTask.ID,
		&patchedTask.UserID,
		&patchedTask.Title,
		&patchedTask.Description,
		&patchedTask.Completed,
		&patchedTask.CreatedAt,
		&patchedTask.CompletedAt,
	); err != nil {
		return nil, err
	}

	return &patchedTask, nil
}
