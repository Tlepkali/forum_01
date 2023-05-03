package category

import (
	"database/sql"
	"errors"

	"forum/internal/models"
)

type CategoryStorage struct {
	db *sql.DB
}

func NewCategoryStorage(db *sql.DB) *CategoryStorage {
	return &CategoryStorage{db: db}
}

func (s *CategoryStorage) GetAllCategories() ([]*models.Category, error) {
	rows, err := s.db.Query("SELECT name FROM category")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	categories := make([]*models.Category, 0)
	for rows.Next() {
		category := new(models.Category)
		err := rows.Scan(&category.Name)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return categories, nil
}

func (s *CategoryStorage) CreateCategory(category *models.Category) (string, error) {
	if category.Name == "" {
		return "", errors.New("category name is empty")
	}

	_, err := s.db.Exec("INSERT INTO category (name) VALUES (?)", category.Name)
	if err != nil {
		return "", err
	}

	return category.Name, nil
}

// TODO: implement
func (s *CategoryStorage) GetCategoryByID(id int) (*models.Category, error) {
	return nil, nil
}

func (s *CategoryStorage) GetCategoryByName(name string) (*models.Category, error) {
	row := s.db.QueryRow("SELECT name FROM category WHERE name = ?", name)

	category := new(models.Category)
	err := row.Scan(&category.Name)
	if err != nil {
		return nil, err
	}

	return category, nil
}
