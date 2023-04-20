package entity

import (
	"time"
)

type UserEntity struct {
	ID           uint                `gorm:"primaryKey;autoIncrement"`
	Username     string              `gorm:"not null; uniqueIndex"`
	Email        string              `gorm:"not null; uniqueIndex"`
	Password     string              `gorm:"not null"`
	Age          int                 `gorm:"not null"`
	Photos       []PhotoEntity       `gorm:"foreignKey:UserID"`
	SocialMedias []SocialMediaEntity `gorm:"foreignKey:UserID"`
	Comments     []CommentEntity     `gorm:"foreignKey:UserID"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
