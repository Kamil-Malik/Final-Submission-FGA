package dto

type PhotoDTO struct {
	ID       uint   `json:"id" form:"id" valid:"optional"`
	Title    string `json:"title" form:"title" valid:"required~Image Title should not be empty"`
	Caption  string `json:"caption" form:"caption" valid:"required~Caption should not be empty"`
	PhotoURL string `json:"photo_url" form:"photo_url" valid:"required~Photo URL should not be empty"`
	UserID   uint   `json:"user_id" form:"user_id" valid:"required~User ID should not be empty"`
}
