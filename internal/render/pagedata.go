package render

import (
	"forum/internal/models"
	"forum/pkg/forms"
)

type PageData struct {
	AuthenticatedUser *models.User
	FlashMessage      string
	Form              *forms.Form
	Posts             []*models.Post
	Post              *models.Post
	Categories        []*models.Category
	Comments          []*models.Comment
}
