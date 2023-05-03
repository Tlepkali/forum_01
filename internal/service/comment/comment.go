package comment

import (
	"time"

	"forum/internal/models"
)

type CommentService struct {
	repo models.CommentRepo
}

func NewCommentService(repo models.CommentRepo) *CommentService {
	return &CommentService{repo}
}

func (c *CommentService) CreateComment(commentDTO *models.CreateCommentDTO) error {
	comment := &models.Comment{
		Content:    commentDTO.Content,
		AuthorID:   commentDTO.AuthorID,
		AuthorName: commentDTO.AuthorName,
		PostID:     commentDTO.PostID,
		CreatedAt:  time.Now(),
	}

	err := c.repo.CreateComment(comment)
	if err != nil {
		return err
	}
	return nil
}

func (c *CommentService) GetAllByPostID(postID int) ([]*models.Comment, error) {
	comments, err := c.repo.GetAllByPostID(postID)
	if err != nil {
		return nil, err
	}

	return comments, nil
}

func (c *CommentService) GetCommentByID(id int) (*models.Comment, error) {
	return nil, nil
}

func (c *CommentService) DeleteComment(id int) error {
	return nil
}

func (c *CommentService) UpdateComment(comment *models.Comment) error {
	return nil
}
