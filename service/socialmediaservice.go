package service

import (
	"Proyek-Akhir-Golang/db"
	"Proyek-Akhir-Golang/entity"
	"errors"
)

func GetAllSocialMedia() (socialMedias []entity.SocialMediaEntity) {
	db := db.GetDB()
	db.Find(&socialMedias).Order("id asc")
	return socialMedias
}

func GetSocialMediaByID(id uint) (socialMedia entity.SocialMediaEntity, err error) {
	db := db.GetDB()
	err = db.Where("id = ?", id).Take(&socialMedia).Error
	return socialMedia, err
}

func InsertSocialMedia(socialMedia entity.SocialMediaEntity) (err error) {
	db := db.GetDB()
	err = db.Create(&socialMedia).Error
	return err
}

func UpdateSocialMedia(socialMedia entity.SocialMediaEntity) (err error) {
	db := db.GetDB()
	model := entity.SocialMediaEntity{}
	if result := db.Model(model).Where("id = ?", socialMedia.ID).Updates(&socialMedia); result.RowsAffected == 0 {
		err = errors.New("social media is not exist")
		return err
	}

	return nil
}

func DeleteSocialMediaByID(id uint) (err error) {
	db := db.GetDB()
	model := entity.SocialMediaEntity{}
	result := db.Delete(&model, id)
	if result.Error != nil {
		err = result.Error
		return err
	}
	if result.RowsAffected == 0 {
		err = errors.New("comment is not exist")
		return err
	}
	return nil
}
