package dto

type SocialMediaDTO struct {
	ID             uint   `json:"id"`
	Name           string `json:"name"`
	SocialMediaURL string `json:"social_media_url"`
	UserID         string `json:"user_id"`
}
