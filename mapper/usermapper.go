package mapper

import (
	"Proyek-Akhir-Golang/dto"
	"Proyek-Akhir-Golang/entity"
)

func DtoToEntity(dto dto.UserDTO) (entity entity.UserEntity) {

	entity.Username = dto.Username
	entity.Email = dto.Email
	entity.Password = dto.Password
	entity.Age = dto.Age

	return entity
}
