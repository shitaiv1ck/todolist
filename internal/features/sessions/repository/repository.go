package sessions_repository

import (
	"database/sql"
	"errors"

	"github.com/shitaiv1ck/todolist/internal/core/domains"
	core_errors "github.com/shitaiv1ck/todolist/internal/core/errors"
	core_postgres "github.com/shitaiv1ck/todolist/internal/core/store/postgres"
)

type SessionsRepository struct {
	store *core_postgres.Store
}

func NewRepository(store *core_postgres.Store) *SessionsRepository {
	return &SessionsRepository{
		store: store,
	}
}

func (r *SessionsRepository) CreateSession(session *domains.Session) (*domains.Session, error) {
	db := r.store.GetDB()

	query := `
		INSERT INTO todolist.sessions(session_token, csrf_token, user_id, created_at, expires_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING *;
	`

	var createdSession domains.Session
	if err := db.QueryRow(
		query,
		session.SessionToken,
		session.CSRFToken,
		session.UserID,
		session.CreatedAt,
		session.ExpiresAt,
	).Scan(
		&createdSession.SessionToken,
		&createdSession.CSRFToken,
		&createdSession.UserID,
		&createdSession.CreatedAt,
		&createdSession.ExpiresAt,
	); err != nil {
		return nil, err
	}

	return &createdSession, nil
}

func (r *SessionsRepository) FindByToken(token string) (*domains.Session, error) {
	db := r.store.GetDB()

	query := `
		SELECT * FROM todolist.sessions
		WHERE session_token = $1;
	`

	var foundSession domains.Session
	if err := db.QueryRow(
		query,
		token,
	).Scan(
		&foundSession.SessionToken,
		&foundSession.CSRFToken,
		&foundSession.UserID,
		&foundSession.CreatedAt,
		&foundSession.ExpiresAt,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, core_errors.ErrNotFound
		}

		return nil, err
	}

	return &foundSession, nil
}

func (r *SessionsRepository) DeleteByToken(sessionToken string) error {
	db := r.store.GetDB()

	query := `
		DELETE FROM todolist.sessions
		WHERE session_token = $1;
	`

	if _, err := db.Exec(
		query,
		sessionToken,
	); err != nil {
		if errors.Is(err, core_errors.ErrNotFound) {
			return core_errors.ErrNotFound
		}

		return err
	}

	return nil
}
