package users_repository

import (
	"database/sql"
	"errors"

	"github.com/lib/pq"
	"github.com/shitaiv1ck/todolist/internal/core/domains"
	core_errors "github.com/shitaiv1ck/todolist/internal/core/errors"
	core_postgres "github.com/shitaiv1ck/todolist/internal/core/store/postgres"
)

type UsersRepository struct {
	store *core_postgres.Store
}

func NewRepository(store *core_postgres.Store) *UsersRepository {
	return &UsersRepository{
		store: store,
	}
}

func (r *UsersRepository) CreateUser(user *domains.User) (*domains.User, error) {
	db := r.store.GetDB()

	query := `
		INSERT INTO todolist.users(username, encrypted_password)
		VALUES ($1, $2)
		RETURNING id, username;
	`

	var savedUser domains.User
	if err := db.QueryRow(
		query,
		user.Username,
		user.EncryptedPassword,
	).Scan(
		&savedUser.ID,
		&savedUser.Username,
	); err != nil {
		if errPQ, ok := err.(*pq.Error); ok {
			if errPQ.Code == "23505" {
				return nil, core_errors.ErrConflict
			}
		}

		return nil, err
	}

	return &savedUser, nil
}

func (r *UsersRepository) FindByUsername(username string) (*domains.User, error) {
	db := r.store.GetDB()

	query := `
		SELECT id, username, encrypted_password
		FROM todolist.users
		WHERE username = $1;
	`

	var foundUser domains.User
	if err := db.QueryRow(
		query,
		username,
	).Scan(
		&foundUser.ID,
		&foundUser.Username,
		&foundUser.EncryptedPassword,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, core_errors.ErrNotFound
		}

		return nil, err
	}

	return &foundUser, nil
}
