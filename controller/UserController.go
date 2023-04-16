package controller

import (
	"Proyek-Akhir-Golang/dto"
	"Proyek-Akhir-Golang/helper"
	"Proyek-Akhir-Golang/mapper"
	"Proyek-Akhir-Golang/service"
	"net/http"

	valid "github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {

	var user dto.UserDTO
	var contentType string

	if err := ctx.BindJSON(&user); err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			dto.GenericResponseDTO{
				Code:    http.StatusBadRequest,
				Message: "Please provide a valid data",
			})
		return
	}

	if _, err := valid.ValidateStruct(user); err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			dto.GenericResponseDTO{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			})
		return
	}

	if user.Age <= 8 {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			dto.GenericResponseDTO{
				Code:    http.StatusBadRequest,
				Message: "You should be at least 9 years old",
			})
		return
	}

	user.Password = helper.HashPass(user.Password)
	if err := service.InsertUser(mapper.DtoToEntity(user)); err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			dto.GenericResponseDTO{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
ct.JSN(
	http.StusCreated,
		to.GeneicResponseTO{
		Code:    http.StatsCrated,
			Message "OK",
		})
}

uc Login(ctx *ginContext) {


