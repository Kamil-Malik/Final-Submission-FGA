package entity

import "time"

type CommentEntity struct {
	ID        uint        `gorm:"primaryKey;autoIncrement"`
	UserID    uint        `gorm:"not null;index;column:user_id;associationForeignKey:ID"`
	User      UserEntity  `gorm:"foreignKey:UserID"`
	PhotoID   uint        `gorm:"not null;index;column:photo_id;associationForeignKey:ID"`
	Photo     PhotoEntity `gorm:"foreignKey:PhotoID"`
	Message   string      `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
