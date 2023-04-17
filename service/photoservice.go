package service

import (
	"Proyek-Akhir-Golang/db"
	"Proyek-Akhir-Golang/entity"
	"errors"
)

func GetAllPhoto() (photo []entity.PhotoEntity) {
	db := db.GetDB()
	db.Find(&photo)
	return photo
}

func GetPhotoByID(id uint) (photo entity.PhotoEntity, err error) {
	db := db.GetDB()
	err = db.Where("id = ?", id).Take(&photo).Error
	return photo, err
}

func InsertPhoto(photo entity.PhotoEntity) (err error) {
	db := db.GetDB()
	err = db.Create(&photo).Error
	return err
}

func UpdatePhoto(photo entity.PhotoEntity) (err error) {
	db := db.GetDB()
	model := entity.PhotoEntity{}
	if result := db.Model(model).Where("id = ?", photo.ID).Updates(&photo); result.RowsAffected == 0 {
		err = errors.New("photo is not exist")
		return err
	}

	return nil
}

func DeletePhotoByID(id uint) (err error) {
	db := db.GetDB()
	model := entity.PhotoEntity{}
	result := db.Delete(&model, id)
	if result.Error != nil {
		err = result.Error
		return err
	}
	if result.RowsAffected == 0 {
		err = errors.New("photo is not exist")
		return err
	}
	return nil
}
