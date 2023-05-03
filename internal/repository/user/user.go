package user

import (
	"database/sql"

	"forum/internal/models"
)

type UserStorage struct {
	db *sql.DB
}

func NewUserStorage(db *sql.DB) *UserStorage {
	return &UserStorage{db: db}
}

func (s *UserStorage) CreateUser(user *models.User) error {
	_, err := s.db.Exec("INSERT INTO user (username, hashed_pw, email, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)",
		user.Username,
		user.HashedPW,
		user.Email,
		user.CreatedAt,
		user.UpdatedAt)
	if err != nil {
		switch err.Error() {
		case "UNIQUE constraint failed: user.email":
			return models.ErrDuplicateEmail
		case "UNIQUE constraint failed: user.username":
			return models.ErrDuplicateUsername
		default:
			return err
		}
	}

	return nil
}

func (s *UserStorage) GetUserByID(id int) (*models.User, error) {
	var user models.User
	err := s.db.QueryRow("SELECT * FROM user WHERE id = $1", id).Scan(&user.ID,
		&user.Username,
		&user.HashedPW,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *UserStorage) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := s.db.QueryRow("SELECT * FROM user WHERE username = $1", username).Scan(&user.ID, &user.Username, &user.HashedPW, &user.Email)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *UserStorage) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := s.db.QueryRow("SELECT * FROM user WHERE email = $1", email).Scan(&user.ID, &user.Username, &user.HashedPW, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *UserStorage) UpdateUser(user *models.User) error {
	_, err := s.db.Exec("UPDATE user SET username = $1, hashed_pw = $2, email = $3 WHERE id = $4", user.Username, user.HashedPW, user.Email, user.ID)
	if err != nil {
		return err
	}

	return nil
}
