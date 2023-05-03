package repository

import (
	"database/sql"

	"forum/internal/models"
	"forum/internal/repository/category"
	"forum/internal/repository/comment"
	"forum/internal/repository/commentvote"
	"forum/internal/repository/post"
	"forum/internal/repository/postvote"
	"forum/internal/repository/session"
	"forum/internal/repository/user"
)

type Repository struct {
	PostRepo        models.PostRepo
	PostVoteRepo    models.PostVoteRepo
	CommentRepo     models.CommentRepo
	CategoryRepo    models.CategoryRepo
	SessionRepo     models.SessionRepo
	UserRepo        models.UserRepo
	CommentVoteRepo models.CommentVoteRepo
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		PostRepo:        post.NewPostStorage(db),
		PostVoteRepo:    postvote.NewPostVoteStorage(db),
		CommentRepo:     comment.NewCommentStorage(db),
		CategoryRepo:    category.NewCategoryStorage(db),
		SessionRepo:     session.NewSessionStorage(db),
		UserRepo:        user.NewUserStorage(db),
		CommentVoteRepo: commentvote.NewCommentVoteStorage(db),
	}
}
