package dto

type PhotoDTO struct {
	ID       uint   `json:"id" form:"id" valid:"optional"`
	Title    string `json:"title" form:"title" valid:"required"`
	Caption  string `json:"caption" form:"caption" valid:"required"`
	PhotoURL string `json:"photo_url" form:"photo_url" valid:"required"`
	UserID   uint   `json:"user_id" form:"user_id" valid:"required"`
}
