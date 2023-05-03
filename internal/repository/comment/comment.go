package comment

import (
	"database/sql"

	"forum/internal/models"
)

type CommentStorage struct {
	db *sql.DB
}

func NewCommentStorage(db *sql.DB) *CommentStorage {
	return &CommentStorage{db: db}
}

func (s CommentStorage) CreateComment(comment *models.Comment) error {
	query := `INSERT INTO comment (content, post_id, author_id, authorname, created_at) VALUES (?, ?, ?, ?, ?)`
	result, err := s.db.Exec(query, comment.Content, comment.PostID, comment.AuthorID, comment.AuthorName, comment.CreatedAt)
	if err != nil {
		return err
	}

	if _, err := result.LastInsertId(); err != nil {
		return err
	}

	if err != nil {
		return err
	}

	return nil
}

func (s CommentStorage) GetAllByPostID(postID int) ([]*models.Comment, error) {
	query := `SELECT id, content, post_id, author_id, authorname, created_at FROM comment WHERE post_id = ?`
	rows, err := s.db.Query(query, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	comments := make([]*models.Comment, 0)
	for rows.Next() {
		comment := new(models.Comment)
		err := rows.Scan(&comment.ID,
			&comment.Content,
			&comment.PostID,
			&comment.AuthorID,
			&comment.AuthorName,
			&comment.CreatedAt)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return comments, nil
}

func (s *CommentStorage) GetCommentByID(id int) (*models.Comment, error) {
	query := `SELECT id, content, post_id, author_id, authorname, created_at FROM comment WHERE id = ?`
	row := s.db.QueryRow(query, id)

	comment := &models.Comment{}
	err := row.Scan(&comment.ID,
		&comment.Content,
		&comment.PostID,
		&comment.AuthorID,
		&comment.AuthorName,
		&comment.CreatedAt)
	if err != nil {
		return nil, err
	}

	return comment, nil
}
