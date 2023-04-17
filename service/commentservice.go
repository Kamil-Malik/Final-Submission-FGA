package service

import (
	"Proyek-Akhir-Golang/db"
	"Proyek-Akhir-Golang/entity"
	"errors"
)

func GetAllComments() (comments []entity.CommentEntity) {
	db := db.GetDB()
	db.Find(&comments).Order("id asc")
	return comments
}

func GetCommentByID(id uint) (comment entity.CommentEntity, err error) {
	db := db.GetDB()
	err = db.Where("id = ?", id).Take(&comment).Error
	return comment, err
}

func InsertComment(comment entity.CommentEntity) (err error) {
	db := db.GetDB()
	err = db.Create(&comment).Error
	return err
}

func UpdateComment(comment entity.CommentEntity) (err error) {
	db := db.GetDB()
	model := entity.CommentEntity{}
	if result := db.Model(model).Where("id = ?", comment.ID).Updates(&comment); result.RowsAffected == 0 {
		err = errors.New("photo is not exist")
		return err
	}

	return nil
}

func DeleteCommentByID(id uint) (err error) {
	db := db.GetDB()
	model := entity.CommentEntity{}
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
