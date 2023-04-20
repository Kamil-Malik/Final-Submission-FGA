package dto

type SocialMediaDTO struct {
	ID             uint   `json:"id"`
	Name           string `json:"name" form:"name" valid:"required~Social Media Name should not be empty"`
	SocialMediaURL string `json:"social_media_url" form:"social_media_url" valid:"required~Social Media URL should not be empty"`
	UserID         uint   `json:"user_id" form:"user_id" valid:"required~User ID should not be empty"`
}
