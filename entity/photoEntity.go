package entity

import "time"

type PhotoEntity struct {
	ID        uint            `gorm:"primaryKey;autoIncrement"`
	Title     string          `gorm:"not null"`
	Caption   string          `gorm:"not null"`
	PhotoURL  string          `gorm:"not null"`
	UserID    string          `gorm:"not null;index;associationForeignKey:ID"`
	Comments  []CommentEntity `gorm:"foreignKey:PhotoID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
