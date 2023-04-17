package mapper

import (
	"Proyek-Akhir-Golang/dto"
	"Proyek-Akhir-Golang/entity"
)

func SocialMediaDTOToEntity(dto dto.SocialMediaDTO) (entity entity.SocialMediaEntity) {
	entity.ID = dto.ID
	entity.Name = dto.Name
	entity.SocialMediaURL = dto.SocialMediaURL
	entity.UserID = dto.UserID
	return entity
}

func SocialMediaEntityToDto(entity entity.SocialMediaEntity) (dto dto.SocialMediaDTO) {
	dto.ID = entity.ID
	dto.Name = entity.Name
	dto.SocialMediaURL = entity.SocialMediaURL
	dto.UserID = entity.UserID
	return dto
}
