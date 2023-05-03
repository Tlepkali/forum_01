package commentvote

import "forum/internal/models"

type CommentVoteService struct {
	repo models.CommentVoteRepo
}

func NewCommentVoteService(repo models.CommentVoteRepo) *CommentVoteService {
	return &CommentVoteService{repo}
}

func (s *CommentVoteService) CreateVote(vote *models.CommentVote) error {
	currentVote, err := s.repo.GetVoteByCommentIDAndAuthorID(vote)
	if err != nil && err != models.ErrSqlNoRows {
		return err
	}

	if currentVote != nil {
		if currentVote.Status == vote.Status {
			return s.repo.DeleteVote(vote)
		}

		return s.repo.UpdateVote(vote)
	}

	return s.repo.CreateVote(vote)
}

func (s *CommentVoteService) GetLikesAndDislikes(commentID int) (int, int, error) {
	votes, err := s.repo.GetVotesByCommentID(commentID)
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

func (s *CommentVoteService) GetVotesByCommentID(id int) ([]*models.CommentVote, error) {
	return s.repo.GetVotesByCommentID(id)
}

func (s *CommentVoteService) UpdateVote(vote *models.CommentVote) error {
	return s.repo.UpdateVote(vote)
}

func (s *CommentVoteService) DeleteVote(vote *models.CommentVote) error {
	return s.repo.DeleteVote(vote)
}
