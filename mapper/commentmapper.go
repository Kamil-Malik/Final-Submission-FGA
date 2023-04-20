package mapper

import (
	"Proyek-Akhir-Golang/dto"
	"Proyek-Akhir-Golang/entity"
)

func CommentDTOToEntity(dto dto.CommentDTO) (entity entity.CommentEntity) {
	entity.ID = dto.ID
	entity.UserID = dto.UserID
	entity.PhotoID = dto.PhotoID
	entity.Message = dto.Message
	return entity
}

func CommentEntityToDTO(entity entity.CommentEntity) (dto dto.CommentDTO) {
	dto.ID = entity.ID
	dto.UserID = entity.UserID
	dto.PhotoID = entity.PhotoID
	dto.Message = entity.Message
	return dto
}
