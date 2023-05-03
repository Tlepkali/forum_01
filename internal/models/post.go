package models

import "time"

type Post struct {
	ID         int         `json:"id"`
	Title      string      `json:"title"`
	Content    string      `json:"content"`
	AuthorID   int         `json:"author_id"`
	AuthorName string      `json:"authorname"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
	Version    int         `json:"version"`
	Categories []*Category `json:"category_id"`
	Comments   []*Comment  `json:"comments"`
	Likes      int         `json:"likes"`
	Dislikes   int         `json:"dislikes"`
}

type CreatePostDTO struct {
	Title      string      `json:"title"`
	Content    string      `json:"content"`
	Author     int         `json:"author"`
	AuthorName string      `json:"authorname"`
	Categories []*Category `json:"category"`
}

type UpdatePostDTO struct {
	Title       string `json:"title"`
	Content     string `json:"content"`
	Category_id int    `json:"category_id"`
}

type DeletePostDTO struct {
	ID int `json:"id"`
}

type PostService interface {
	CreatePost(post *CreatePostDTO) (int, error)
	GetAllPosts(offset, limit int) ([]*Post, error)
	GetLikedPosts(userID int) ([]*Post, error)
	GetPostByID(id int) (*Post, error)
	GetPostsByTitle(title string) ([]*Post, error)
	GetPostsByAuthorID(author int) ([]*Post, error)
	GetPostsByCategory(category string) ([]*Post, error)
	UpdatePost(post *Post) error
	DeletePost(id int) error
}

type PostRepo interface {
	CreatePost(post *Post) (int, error)
	GetAllPosts(offset, limit int) ([]*Post, error)
	GetLikedPosts(userID int) ([]*Post, error)
	GetPostByID(id int) (*Post, error)
	GetPostByTitle(title string) ([]*Post, error)
	GetPostsByAuthor(author int) ([]*Post, error)
	GetPostsByCategory(category string) ([]*Post, error)
	UpdatePost(post *Post) error
	DeletePost(id int) error
}
