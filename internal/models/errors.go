package models

import (
	"database/sql"
	"errors"
)

var (
	ErrSqlNoRows            = sql.ErrNoRows
	ErrInvalidCredentials   = errors.New("invalid credentials")
	ErrSessionExpired       = errors.New("session expired")
	ErrDuplicateEmail       = errors.New("duplicate email")
	ErrDuplicateUsername    = errors.New("duplicate username")
	ErrSessionAlreadyExists = errors.New("session already exists")
)
