package post

import (
	"context"
	"database/sql"
	"time"

	"forum/internal/models"
)

type PostStorage struct {
	db *sql.DB
}

func NewPostStorage(db *sql.DB) *PostStorage {
	return &PostStorage{db: db}
}

func (s *PostStorage) GetAllPosts(offset, limit int) ([]*models.Post, error) {
	query := `SELECT * FROM post ORDER BY id DESC LIMIT $1 OFFSET $2`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := s.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []*models.Post

	for rows.Next() {
		post := models.Post{}

		err := rows.Scan(&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.AuthorName,
			&post.CreatedAt,
			&post.UpdatedAt,
			&post.Version)
		if err != nil {
			return nil, err
		}

		posts = append(posts, &post)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

func (s *PostStorage) GetLikedPosts(userID int) ([]*models.Post, error) {
	query := `SELECT p.id, p.title, p.content, p.author_id, p.authorname, p.created_at, p.updated_at, p.version FROM post p
	JOIN post_vote a ON p.id = a.post_id
	WHERE a.author_id = $1 AND a.status = 1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := s.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []*models.Post

	for rows.Next() {
		post := models.Post{}

		err := rows.Scan(&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.AuthorName,
			&post.CreatedAt,
			&post.UpdatedAt,
			&post.Version)
		if err != nil {
			return nil, err
		}

		posts = append(posts, &post)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

func (s *PostStorage) CreatePost(p *models.Post) (int, error) {
	query := `INSERT INTO post (title, content, author_id, authorname, created_at, updated_at, version) 
	VALUES ($1, $2, $3, $4, $5, $6, $7) 
	RETURNING id, created_at, updated_at, version`

	args := []interface{}{p.Title, p.Content, p.AuthorID, p.AuthorName, p.CreatedAt, p.UpdatedAt, p.Version}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := s.db.QueryRowContext(ctx, query, args...).Scan(&p.ID, &p.CreatedAt, &p.UpdatedAt, &p.Version)
	if err != nil {
		return 0, err
	}

	// Create post_categories entries
	for _, category := range p.Categories {
		query = `INSERT INTO post_categories (post_id, category_name) VALUES ($1, $2)`
		_, err = s.db.ExecContext(ctx, query, p.ID, category.Name)
		if err != nil {
			return 0, err
		}
	}

	return p.ID, nil
}

func (s *PostStorage) GetPostByID(id int) (*models.Post, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}

	query := `SELECT * FROM post WHERE id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	post := &models.Post{}

	err := s.db.QueryRowContext(ctx, query, id).Scan(&post.ID,
		&post.Title,
		&post.Content,
		&post.AuthorID,
		&post.AuthorName,
		&post.CreatedAt,
		&post.UpdatedAt,
		&post.Version)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	// Get post categories
	query = `SELECT c.name FROM category c
	JOIN post_categories pc ON c.name = pc.category_name
	WHERE pc.post_id = $1`

	rows, err := s.db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		category := models.Category{}

		err := rows.Scan(&category.Name)
		if err != nil {
			return nil, err
		}

		post.Categories = append(post.Categories, &category)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return post, nil
}

func (s *PostStorage) GetPostByTitle(title string) ([]*models.Post, error) {
	query := `SELECT * FROM post WHERE title = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := s.db.QueryContext(ctx, query, title)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []*models.Post

	for rows.Next() {
		post := models.Post{}

		err := rows.Scan(&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.AuthorName,
			&post.CreatedAt,
			&post.UpdatedAt,
			&post.Version)
		if err != nil {
			return nil, err
		}

		posts = append(posts, &post)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

func (s *PostStorage) GetPostsByAuthor(author int) ([]*models.Post, error) {
	query := `SELECT * FROM post WHERE author_id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := s.db.QueryContext(ctx, query, author)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []*models.Post

	for rows.Next() {
		post := models.Post{}

		err := rows.Scan(&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.AuthorName,
			&post.CreatedAt,
			&post.UpdatedAt,
			&post.Version)
		if err != nil {
			return nil, err
		}

		posts = append(posts, &post)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

func (s *PostStorage) GetPostsByCategory(category string) ([]*models.Post, error) {
	query := `SELECT p.* FROM post p
	JOIN post_categories pc ON p.id = pc.post_id
	WHERE pc.category_name = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := s.db.QueryContext(ctx, query, category)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []*models.Post

	for rows.Next() {
		post := models.Post{}

		err := rows.Scan(&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.AuthorName,
			&post.CreatedAt,
			&post.UpdatedAt,
			&post.Version)
		if err != nil {
			return nil, err
		}

		posts = append(posts, &post)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

func (s *PostStorage) UpdatePost(p *models.Post) error {
	query := `UPDATE post SET title = $1, content = $2, updated_at = $3, version = $4 WHERE id = $5`

	args := []interface{}{p.Title, p.Content, p.UpdatedAt, p.Version, p.ID}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := s.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (s *PostStorage) DeletePost(id int) error {
	query := `DELETE FROM post WHERE id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := s.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
