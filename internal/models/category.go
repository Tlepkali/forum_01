package models

type Category struct {
	Name string `json:"name"`
}

type CategoryService interface {
	CreateCategory(c *Category) (string, error)
	GetAllCategories() ([]*Category, error)
	GetCategoryByName(name string) (*Category, error)
}

type CategoryRepo interface {
	CreateCategory(c *Category) (string, error)
	GetAllCategories() ([]*Category, error)
	GetCategoryByName(name string) (*Category, error)
}
