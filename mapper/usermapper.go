package mapper

import (
	"Proyek-Akhir-Golang/dto"
	"Proyek-Akhir-Golang/entity"
)

func DtoToEntity(dto dto.UserDTO) entity.UserEntity {
	var user entity.UserEntity

	user.Username = dto.Username
	user.Email = dto.Email
	user.Password = dto.Password
	user.Age = dto.Age

	return user
}
