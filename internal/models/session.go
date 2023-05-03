package models

import "time"

type Session struct {
	User_id  int       `json:"user_id"`
	UUID     string    `json:"uuid"`
	ExpireAt time.Time `json:"expire_at"`
}

type SessionService interface {
	CreateSession(userID int) (*Session, error)
	GetSessionByUUID(uuid string) (*Session, error)
	UpdateByUserID(s *Session) error
	DeleteSessionByUUID(uuid string) error
}

type SessionRepo interface {
	CreateSession(s *Session) error
	GetSessionByUUID(uuid string) (*Session, error)
	GetSessionByUserID(userID int) (*Session, error)
	UpdateByUserID(s *Session) error
	DeleteSessionByUUID(uuid string) error
}
