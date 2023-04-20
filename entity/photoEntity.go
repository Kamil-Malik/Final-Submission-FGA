package entity

import "time"

type PhotoEntity struct {
	ID        uint            `gorm:"primaryKey;autoIncrement"`
	Title     string          `gorm:"not null"`
	Caption   string          `gorm:"not null"`
	PhotoURL  string          `gorm:"not null"`
	UserID    uint            `gorm:"not null;index;column:user_id;associationForeignKey:ID"`
	Comments  []CommentEntity `gorm:"foreignKey:PhotoID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
