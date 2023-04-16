package service

import (
	"Proyek-Akhir-Golang/db"
	"Proyek-Akhir-Golang/entity"
)

func InsertUser(user entity.UserEntity) error {

	db := db.GetDB()
	err := db.Create(&user).Error

	return err
}
