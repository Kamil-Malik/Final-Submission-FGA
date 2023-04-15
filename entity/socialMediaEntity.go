package entity

import "time"

type SocialMediaEntity struct {
	ID             uint   `gorm:"primaryKey;autoIncrement"`
	Name           string `gorm:"not null"`
	SocialMediaURL string `gorm:"not null"`
	UserID         string `gorm:"not null"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
