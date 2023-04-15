package dto

type CommentDTO struct {
	ID      uint   `json:"id"`
	UserID  string `json:"user_id"`
	PhotoID string `json:"photo_id"`
	Message string `json:"message"`
}
