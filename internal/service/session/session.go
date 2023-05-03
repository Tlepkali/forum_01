package session

import (
	"time"

	"forum/internal/models"

	"github.com/gofrs/uuid"
)

type SessionService struct {
	repo models.SessionRepo
}

func NewSessionService(repo models.SessionRepo) *SessionService {
	return &SessionService{repo}
}

func (s *SessionService) CreateSession(userID int) (*models.Session, error) {
	oldSession, _ := s.repo.GetSessionByUserID(userID)
	if oldSession != nil {
		return nil, models.ErrSessionAlreadyExists
	}

	uuid, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	session := &models.Session{
		User_id:  userID,
		UUID:     uuid.String(),
		ExpireAt: time.Now().Add(time.Hour),
	}

	err = s.repo.CreateSession(session)
	if err != nil {
		return nil, err
	}

	return session, nil
}

func (s *SessionService) GetSessionByUUID(uuid string) (*models.Session, error) {
	session, err := s.repo.GetSessionByUUID(uuid)

	switch err {
	case nil:
		if session.ExpireAt.Before(time.Now()) {
			return nil, models.ErrSessionExpired
		}
		return session, nil
	case models.ErrSqlNoRows:
		return nil, models.ErrSqlNoRows
	default:
		return nil, err
	}
}

func (s *SessionService) UpdateByUserID(session *models.Session) error {
	return s.repo.UpdateByUserID(session)
}

func (s *SessionService) DeleteSessionByUUID(uuid string) error {
	return s.repo.DeleteSessionByUUID(uuid)
}
