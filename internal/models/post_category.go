package models

// PostCategory represents a post-category relationship, but not using in the project as a struct.
type PostCategory struct {
	PostID     int64 `json:"post_id"`
	CategoryID int64 `json:"category_id"`
}
