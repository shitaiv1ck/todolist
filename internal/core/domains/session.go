package domains

import "time"

type Session struct {
	SessionToken string
	CSRFToken    string
	UserID       int
	CreatedAt    time.Time
	ExpiresAt    time.Time
}

func NewSession(
	sessionToken string,
	csrfToken string,
	userID int,
	createdAt time.Time,
	expiresAt time.Time,
) *Session {
	return &Session{
		SessionToken: sessionToken,
		CSRFToken:    csrfToken,
		UserID:       userID,
		CreatedAt:    createdAt,
		ExpiresAt:    expiresAt,
	}
}
