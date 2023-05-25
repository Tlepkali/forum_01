package models

type Image struct {
	ID        int    `json:"id"`
	PostID    int    `json:"post_id"`
	UserID    int    `json:"user_id"`
	ImagePath string `json:"image_path"`
}

type ImageService interface {
	UploadImage(file []byte, filename string) error
	GetImage(id int, originID int, origin string) (*Image, error)
	DeleteImage(id int) error
}
