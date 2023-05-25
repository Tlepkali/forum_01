package post

import (
	"io/ioutil"
	"strings"
	"time"

	"forum/internal/models"

	"github.com/gofrs/uuid"
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
	return p.repo.CreatePost(post)
}

func (p *PostService) CreatePostWithImage(postDTO *models.CreatePostDTO) (int, error) {
	if postDTO.ImageFile == nil {
		return p.CreatePost(postDTO)
	}

	data, err := ioutil.ReadAll(postDTO.ImageFile)
	if err != nil {
		return 0, err
	}

	fileName, err := uuid.NewV4()
	if err != nil {
		return 0, err
	}
	filePath := "ui/static/img/" + fileName.String()

	post := &models.Post{
		Title:      postDTO.Title,
		Content:    postDTO.Content,
		AuthorID:   postDTO.Author,
		AuthorName: postDTO.AuthorName,
		Categories: postDTO.Categories,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		Version:    1,
		ImagePath:  filePath,
	}
	id, err := p.repo.CreatePostWithImage(post)
	if err != nil {
		return 0, err
	}

	err = ioutil.WriteFile(filePath, data, 0o666)
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

	if post.ImagePath == "" {
		return post, nil
	}

	post.ImagePath = ".." + strings.TrimPrefix(post.ImagePath, "ui")

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
