package session

import (
	"database/sql"

	"forum/internal/models"
)

type SessionStorage struct {
	db *sql.DB
}

func NewSessionStorage(db *sql.DB) *SessionStorage {
	return &SessionStorage{db: db}
}

func (s *SessionStorage) CreateSession(session *models.Session) error {
	_, err := s.db.Exec("INSERT INTO session (uuid, user_id, expire_at) VALUES ($1, $2, $3)", session.UUID, session.User_id, session.ExpireAt)
	if err != nil {
		return err
	}

	return nil
}

func (s *SessionStorage) GetSessionByUserID(userID int) (*models.Session, error) {
	var session models.Session
	err := s.db.QueryRow("SELECT * FROM session WHERE user_ID = $1", userID).Scan(&session.UUID, &session.User_id, &session.ExpireAt)
	if err != nil {
		return nil, err
	}

	return &session, nil
}

func (s *SessionStorage) GetSessionByUUID(sessionID string) (*models.Session, error) {
	var session models.Session
	err := s.db.QueryRow("SELECT * FROM session WHERE uuid = $1", sessionID).Scan(&session.UUID, &session.User_id, &session.ExpireAt)
	if err != nil {
		return nil, err
	}

	return &session, nil
}

func (s *SessionStorage) DeleteSessionByUUID(sessionID string) error {
	_, err := s.db.Exec("DELETE FROM session WHERE uuid = $1", sessionID)
	if err != nil {
		return err
	}

	return nil
}

func (s *SessionStorage) UpdateByUserID(session *models.Session) error {
	_, err := s.db.Exec("UPDATE session SET uuid = $1 WHERE user_id = $2", session.UUID, session.User_id)
	if err != nil {
		return err
	}

	return nil
}
