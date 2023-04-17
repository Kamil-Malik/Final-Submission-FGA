package entity

import "time"

type SocialMediaEntity struct {
	ID             uint       `gorm:"primaryKey;autoIncrement"`
	Name           string     `gorm:"not null"`
	SocialMediaURL string     `gorm:"not null"`
	UserID         uint       `gorm:"not null;column:user_id;index;associationForeignKey:ID"`
	User           UserEntity `gorm:"foreignKey:UserID"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
