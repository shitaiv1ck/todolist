package sessions_service

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/shitaiv1ck/todolist/internal/core/domains"
	core_errors "github.com/shitaiv1ck/todolist/internal/core/errors"
)

type SessionsService struct {
	usersRep    UsersRepository
	sessionsRep SessionsRepository
}

type UsersRepository interface {
	GetUserByUsername(username string) (*domains.User, error)
}

type SessionsRepository interface {
	CreateSession(session *domains.Session) (*domains.Session, error)
	GetSessionByToken(token string) (*domains.Session, error)
	DeleteSessionByToken(sessionToken string) error
}

func NewService(usersRep UsersRepository, sessionsRep SessionsRepository) *SessionsService {
	return &SessionsService{
		usersRep:    usersRep,
		sessionsRep: sessionsRep,
	}
}

func (s *SessionsService) Authenticate(username string, password string) (int, error) {
	user, err := s.usersRep.GetUserByUsername(username)
	if err != nil {
		return -1, core_errors.ErrUnautorize
	}

	if err := user.ComparePassword(password); err != nil {
		return -1, core_errors.ErrUnautorize
	}

	return user.ID, nil
}

func (s *SessionsService) CreateSession(userID int) (*domains.Session, error) {
	sessionToken, err := generateToken(32)
	if err != nil {
		return nil, fmt.Errorf("failed to generate session token: %w", err)
	}

	csrfToken, err := generateToken(32)
	if err != nil {
		return nil, fmt.Errorf("failed to generate csrf token: %w", err)
	}

	session := domains.NewSession(
		sessionToken,
		csrfToken,
		userID,
		time.Now(),
		time.Now().Add(24*time.Hour),
	)

	createdSession, err := s.sessionsRep.CreateSession(session)
	if err != nil {
		return nil, fmt.Errorf("failed to create session: %w", err)
	}

	return createdSession, nil

}

func (s *SessionsService) FindByToken(token string) (*domains.Session, error) {
	foundSession, err := s.sessionsRep.GetSessionByToken(token)
	if err != nil {
		return nil, err
	}

	if !foundSession.ExpiresAt.After(time.Now()) {
		return nil, core_errors.ErrUnautorize
	}

	return foundSession, nil
}

func generateToken(len int) (string, error) {
	bytes := make([]byte, len)

	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	token := base64.URLEncoding.EncodeToString(bytes)

	return token, nil
}

func (s *SessionsService) DeleteByToken(sessionToken string) error {
	if err := s.sessionsRep.DeleteSessionByToken(sessionToken); err != nil {
		return err
	}

	return nil
}
