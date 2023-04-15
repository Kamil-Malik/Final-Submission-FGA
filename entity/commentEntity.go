package entity

import "time"

type CommentEntity struct {
	ID        uint        `gorm:"primaryKey;autoIncrement"`
	UserID    string      `gorm:"not null;index;associationForeignKey:ID"`
	PhotoID   string      `gorm:"not null;index;associationForeignKey:ID"`
	Photo     PhotoEntity `gorm:"foreignKey:PhotoID"`
	Message   string      `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
