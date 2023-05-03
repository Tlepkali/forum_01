package models

import "time"

type Comment struct {
	ID         int       `json:"id"`
	Content    string    `json:"content"`
	AuthorID   int       `json:"author_id"`
	AuthorName string    `json:"authorname"`
	PostID     int       `json:"post_id"`
	CreatedAt  time.Time `json:"created_at"`
	Likes      int       `json:"likes"`
	Dislikes   int       `json:"dislikes"`
}

type CreateCommentDTO struct {
	Content    string `json:"content"`
	AuthorID   int    `json:"author_id"`
	AuthorName string `json:"authorname"`
	PostID     int    `json:"post_id"`
}

type CommentService interface {
	CreateComment(comment *CreateCommentDTO) error
	GetAllByPostID(postID int) ([]*Comment, error)
	GetCommentByID(id int) (*Comment, error)
}

type CommentRepo interface {
	CreateComment(comment *Comment) error
	GetAllByPostID(postID int) ([]*Comment, error)
	GetCommentByID(id int) (*Comment, error)
}
