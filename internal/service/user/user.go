package user

import (
	"time"

	"forum/internal/models"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo models.UserRepo
}

func NewUserService(repo models.UserRepo) *UserService {
	return &UserService{repo}
}

func (u *UserService) CreateUser(userDTO *models.CreateUserDTO) error {
	hashedPW, err := bcrypt.GenerateFromPassword([]byte(userDTO.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &models.User{
		Username:  userDTO.Username,
		Email:     userDTO.Email,
		HashedPW:  string(hashedPW),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = u.repo.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserService) GetUserByID(id int) (*models.User, error) {
	user, err := u.repo.GetUserByID(id)
	if err != nil {
		switch err {
		case models.ErrSqlNoRows:
			return nil, models.ErrSqlNoRows
		default:
			return nil, err
		}
	}

	return user, nil
}

func (u *UserService) GetUserByUsername(username string) (*models.User, error) {
	return nil, nil
}

func (u *UserService) LoginUser(userDTO *models.LoginUserDTO) (int, error) {
	user, err := u.repo.GetUserByEmail(userDTO.Email)
	if err != nil {
		switch err {
		case models.ErrSqlNoRows:
			return 0, models.ErrInvalidCredentials
		default:
			return 0, err
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.HashedPW), []byte(userDTO.Password))
	if err != nil {
		return 0, models.ErrInvalidCredentials
	}

	return user.ID, nil
}

func (u *UserService) GetUserByEmail(email string) (*models.User, error) {
	return nil, nil
}

func (u *UserService) UpdateUser(user *models.UpdateUserDTO) error {
	return nil
}
