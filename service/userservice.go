package service

import (
	"Proyek-Akhir-Golang/db"
	"Proyek-Akhir-Golang/entity"
)

func InsertUser(user entity.UserEntity) (err error) {
	db := db.GetDB()
	err = db.Create(&user).Error
	return err
}

func GetUserByEmail(email string) (user entity.UserEntity, err error) {
	db := db.GetDB()
	err = db.Where("email = ?", email).First(&user).Error
	return user, err
}
