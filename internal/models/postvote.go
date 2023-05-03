package models

type PostVote struct {
	AuthorID int   `json:"author_id"`
	PostID   int   `json:"post_id"`
	Status   uint8 `json:"status"`
}

type PostVoteService interface {
	CreateVote(v *PostVote) error
	GetVotesByPostID(postID int) ([]*PostVote, error)
	GetVotesByAuthorID(authorID int) ([]*PostVote, error)
	UpdateVote(v *PostVote) error
	DeleteVote(v *PostVote) error
	GetLikesAndDislikes(postID int) (int, int, error)
}

type PostVoteRepo interface {
	CreateVote(v *PostVote) error
	GetVotesByPostID(postID int) ([]*PostVote, error)
	GetVotesByAuthorID(authorID int) ([]*PostVote, error)
	GetVoteByPostIDAndAuthorID(v *PostVote) (*PostVote, error)
	UpdateVote(v *PostVote) error
	DeleteVote(postID, authorID int) error
}
