package dto

type CommentDTO struct {
	ID      uint   `json:"id" form:"id" valid:"optional"`
	UserID  uint   `json:"user_id" form:"user_id" valid:"required"`
	PhotoID uint   `json:"photo_id" form:"photo_id" valid:"required; type(int)"`
	Message string `json:"message" form:"message" valid:"required; type(string)"`
}
