package mapper

import (
	"Proyek-Akhir-Golang/dto"
	"Proyek-Akhir-Golang/entity"
)

func PhotoEntityToDTO(entity entity.PhotoEntity) (dto dto.PhotoDTO) {
	dto.ID = entity.ID
	dto.Title = entity.Title
	dto.Caption = entity.Caption
	dto.PhotoURL = entity.PhotoURL
	dto.UserID = entity.UserID
	return dto
}

func PhotoDTOToEntity(dto dto.PhotoDTO) (entity entity.PhotoEntity) {
	entity.ID = dto.ID
	entity.Title = dto.Title
	entity.Caption = dto.Caption
	entity.PhotoURL = dto.PhotoURL
	entity.UserID = dto.UserID
	return entity
}
