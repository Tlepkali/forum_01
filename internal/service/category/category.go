package category

import (
	"forum/internal/models"
)

type CategoryService struct {
	repo models.CategoryRepo
}

func NewCategoryService(repo models.CategoryRepo) *CategoryService {
	return &CategoryService{repo}
}

func (c *CategoryService) CreateCategory(category *models.Category) (string, error) {
	name, err := c.repo.CreateCategory(category)
	if err != nil {
		return "", err
	}

	return name, nil
}

func (c *CategoryService) GetAllCategories() ([]*models.Category, error) {
	categories, err := c.repo.GetAllCategories()
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (c *CategoryService) GetCategoryByName(name string) (*models.Category, error) {
	category, err := c.repo.GetCategoryByName(name)
	if err != nil {
		return nil, err
	}

	return category, nil
}
