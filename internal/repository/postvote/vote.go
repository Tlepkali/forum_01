package postvote

import (
	"database/sql"

	"forum/internal/models"
)

type PostVoteStorage struct {
	db *sql.DB
}

func NewPostVoteStorage(db *sql.DB) *PostVoteStorage {
	return &PostVoteStorage{db: db}
}

func (s *PostVoteStorage) CreateVote(v *models.PostVote) error {
	_, err := s.db.Exec("INSERT INTO post_vote (author_id, post_id, status) VALUES ($1, $2, $3)", v.AuthorID, v.PostID, v.Status)
	if err != nil {
		return err
	}

	return nil
}

func (s *PostVoteStorage) GetVotesByPostID(postID int) ([]*models.PostVote, error) {
	var votes []*models.PostVote
	rows, err := s.db.Query("SELECT * FROM post_vote WHERE post_id = $1", postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var v models.PostVote
		err := rows.Scan(&v.AuthorID, &v.PostID, &v.Status)
		if err != nil {
			return nil, err
		}
		votes = append(votes, &v)
	}

	return votes, nil
}

func (s *PostVoteStorage) GetVotesByAuthorID(authorID int) ([]*models.PostVote, error) {
	var votes []*models.PostVote
	rows, err := s.db.Query("SELECT * FROM post_vote WHERE author_id = $1", authorID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var v models.PostVote
		err := rows.Scan(&v.AuthorID, &v.PostID, &v.Status)
		if err != nil {
			return nil, err
		}
		votes = append(votes, &v)
	}

	return votes, nil
}

func (s *PostVoteStorage) GetVoteByPostIDAndAuthorID(pV *models.PostVote) (*models.PostVote, error) {
	var v models.PostVote
	err := s.db.QueryRow("SELECT * FROM post_vote WHERE post_id = $1 AND author_id = $2", pV.PostID, pV.AuthorID).Scan(&v.AuthorID,
		&v.PostID,
		&v.Status)
	if err != nil {
		return nil, err
	}

	return &v, nil
}

func (s *PostVoteStorage) UpdateVote(v *models.PostVote) error {
	_, err := s.db.Exec("UPDATE post_vote SET status = $1 WHERE author_id = $2 AND post_id = $3", v.Status, v.AuthorID, v.PostID)
	if err != nil {
		return err
	}

	return nil
}

func (s *PostVoteStorage) DeleteVote(postID, authorID int) error {
	_, err := s.db.Exec("DELETE FROM post_vote WHERE post_id = $1 AND author_id = $2", postID, authorID)
	if err != nil {
		return err
	}

	return nil
}
