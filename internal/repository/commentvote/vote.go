package commentvote

import (
	"database/sql"

	"forum/internal/models"
)

type CommentVoteStorage struct {
	db *sql.DB
}

func NewCommentVoteStorage(db *sql.DB) *CommentVoteStorage {
	return &CommentVoteStorage{db: db}
}

func (s *CommentVoteStorage) CreateVote(v *models.CommentVote) error {
	_, err := s.db.Exec("INSERT INTO comment_vote (author_id, comment_id, status) VALUES ($1, $2, $3)", v.AuthorID, v.CommentID, v.Status)
	if err != nil {
		return err
	}

	return nil
}

func (s *CommentVoteStorage) GetVotesByCommentID(commentID int) ([]*models.CommentVote, error) {
	var votes []*models.CommentVote
	rows, err := s.db.Query("SELECT * FROM comment_vote WHERE comment_id = $1", commentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var v models.CommentVote
		err := rows.Scan(&v.AuthorID, &v.CommentID, &v.Status)
		if err != nil {
			return nil, err
		}
		votes = append(votes, &v)
	}

	return votes, nil
}

func (s *CommentVoteStorage) GetVotesByAuthorID(authorID int) ([]*models.CommentVote, error) {
	var votes []*models.CommentVote
	rows, err := s.db.Query("SELECT * FROM comment_vote WHERE author_id = $1", authorID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var v models.CommentVote
		err := rows.Scan(&v.AuthorID, &v.CommentID, &v.Status)
		if err != nil {
			return nil, err
		}
		votes = append(votes, &v)
	}

	return votes, nil
}

func (s *CommentVoteStorage) GetVoteByCommentIDAndAuthorID(cV *models.CommentVote) (*models.CommentVote, error) {
	var v models.CommentVote
	err := s.db.QueryRow("SELECT * FROM comment_vote WHERE comment_id = $1 AND author_id = $2", cV.CommentID, cV.AuthorID).Scan(&v.AuthorID,
		&v.CommentID,
		&v.Status)
	if err != nil {
		return nil, err
	}

	return &v, nil
}

func (s *CommentVoteStorage) UpdateVote(v *models.CommentVote) error {
	_, err := s.db.Exec("UPDATE comment_vote SET status = $1 WHERE comment_id = $2 AND author_id = $3", v.Status, v.CommentID, v.AuthorID)
	if err != nil {
		return err
	}

	return nil
}

func (s *CommentVoteStorage) DeleteVote(v *models.CommentVote) error {
	_, err := s.db.Exec("DELETE FROM comment_vote WHERE comment_id = $1 AND author_id = $2", v.CommentID, v.AuthorID)
	if err != nil {
		return err
	}

	return nil
}
