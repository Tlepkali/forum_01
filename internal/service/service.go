package service

import (
	"forum/internal/models"
	"forum/internal/repository"
	"forum/internal/service/category"
	"forum/internal/service/comment"
	"forum/internal/service/commentvote"
	"forum/internal/service/post"
	"forum/internal/service/postvote"
	"forum/internal/service/session"
	"forum/internal/service/user"
)

type Service struct {
	UserService        models.UserService
	PostService        models.PostService
	CommentService     models.CommentService
	PostVoteService    models.PostVoteService
	SessionService     models.SessionService
	CategoryService    models.CategoryService
	CommentVoteService models.CommentVoteService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		UserService:        user.NewUserService(repo.UserRepo),
		PostService:        post.NewPostService(repo.PostRepo),
		CategoryService:    category.NewCategoryService(repo.CategoryRepo),
		CommentService:     comment.NewCommentService(repo.CommentRepo),
		PostVoteService:    postvote.NewPostVoteService(repo.PostVoteRepo),
		CommentVoteService: commentvote.NewCommentVoteService(repo.CommentVoteRepo),
		SessionService:     session.NewSessionService(repo.SessionRepo),
	}
}
