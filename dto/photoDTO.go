package dto

type PhotoDTO struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url"`
	UserID   string `json:"user_id"`
}
