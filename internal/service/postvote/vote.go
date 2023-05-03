package postvote

import "forum/internal/models"

type PostVoteService struct {
	repo models.PostVoteRepo
}

func NewPostVoteService(repo models.PostVoteRepo) *PostVoteService {
	return &PostVoteService{repo}
}

func (s *PostVoteService) CreateVote(vote *models.PostVote) error {
	currentVote, err := s.repo.GetVoteByPostIDAndAuthorID(vote)
	if err != nil && err != models.ErrSqlNoRows {
		return err
	}

	if currentVote != nil {
		if currentVote.Status == vote.Status {
			return s.repo.DeleteVote(vote.PostID, vote.AuthorID)
		}

		return s.repo.UpdateVote(vote)
	}

	return s.repo.CreateVote(vote)
}

func (s *PostVoteService) GetLikesAndDislikes(postID int) (int, int, error) {
	votes, err := s.repo.GetVotesByPostID(postID)
	if err != nil {
		return 0, 0, err
	}

	var likes, dislikes int
	for _, v := range votes {
		if v.Status == 1 {
			likes++
		} else {
			dislikes++
		}
	}

	return likes, dislikes, nil
}

func (s *PostVoteService) GetVotesByPostID(id int) ([]*models.PostVote, error) {
	return s.repo.GetVotesByPostID(id)
}

func (s *PostVoteService) GetVotesByAuthorID(id int) ([]*models.PostVote, error) {
	return s.repo.GetVotesByAuthorID(id)
}

func (s *PostVoteService) UpdateVote(vote *models.PostVote) error {
	return s.repo.UpdateVote(vote)
}

func (s *PostVoteService) DeleteVote(vote *models.PostVote) error {
	return s.repo.DeleteVote(vote.PostID, vote.AuthorID)
}
