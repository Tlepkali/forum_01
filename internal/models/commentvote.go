package models

type CommentVote struct {
	AuthorID  int   `json:"author_id"`
	CommentID int   `json:"post_id"`
	Status    uint8 `json:"status"`
}

type CommentVoteService interface {
	CreateVote(v *CommentVote) error
	GetVotesByCommentID(commentID int) ([]*CommentVote, error)
	UpdateVote(v *CommentVote) error
	DeleteVote(v *CommentVote) error
	GetLikesAndDislikes(commentID int) (int, int, error)
}

type CommentVoteRepo interface {
	CreateVote(v *CommentVote) error
	GetVotesByCommentID(commentID int) ([]*CommentVote, error)
	GetVoteByCommentIDAndAuthorID(v *CommentVote) (*CommentVote, error)
	UpdateVote(v *CommentVote) error
	DeleteVote(v *CommentVote) error
}
