package db

import (
	"Proyek-Akhir-Golang/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func StartDB() {
	connStr := "postgres://xpohkkbv:oVGg1igO9NrC9OFAxXO_RNA2QT57GeA8@john.db.elephantsql.com/xpohkkbv"
	db, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.Debug().AutoMigrate(
		entity.UserEntity{},
		entity.SocialMediaEntity{},
		entity.PhotoEntity{},
		entity.CommentEntity{})
}

func GetDB() *gorm.DB {
	return db
}
