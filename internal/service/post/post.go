package post

import (
	"time"

	"forum/internal/models"
)

type PostService struct {
	repo models.PostRepo
}

func NewPostService(repo models.PostRepo) *PostService {
	return &PostService{repo}
}

func (p *PostService) CreatePost(postDTO *models.CreatePostDTO) (int, error) {
	post := &models.Post{
		Title:      postDTO.Title,
		Content:    postDTO.Content,
		AuthorID:   postDTO.Author,
		AuthorName: postDTO.AuthorName,
		Categories: postDTO.Categories,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		Version:    1,
	}
	id, err := p.repo.CreatePost(post)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (p *PostService) GetAllPosts(offset, limit int) ([]*models.Post, error) {
	return p.repo.GetAllPosts(offset, limit)
}

func (p *PostService) GetLikedPosts(userID int) ([]*models.Post, error) {
	return p.repo.GetLikedPosts(userID)
}

func (p *PostService) GetPostByID(id int) (*models.Post, error) {
	post, err := p.repo.GetPostByID(id)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (p *PostService) GetPostsByCategory(category string) ([]*models.Post, error) {
	return p.repo.GetPostsByCategory(category)
}

func (p *PostService) GetPostsByTitle(title string) ([]*models.Post, error) {
	return nil, nil
}

func (p *PostService) GetPostsByAuthorID(authorID int) ([]*models.Post, error) {
	return p.repo.GetPostsByAuthor(authorID)
}

func (p *PostService) GetPostsByCategoryID(categoryID int) ([]*models.Post, error) {
	return nil, nil
}

func (p *PostService) UpdatePost(post *models.Post) error {
	return nil
}

func (p *PostService) DeletePost(id int) error {
	return nil
}
