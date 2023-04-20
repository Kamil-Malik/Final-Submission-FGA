package helper

import (
	"Proyek-Akhir-Golang/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AbortGenericResponse(ctx *gin.Context, code int, message string) {
	ctx.AbortWithStatusJSON(
		code,
		dto.GenericResponseDTO{
			Code:    code,
			Message: message,
		},
	)
}

func SendGenericResponse(ctx *gin.Context) {
	ctx.JSON(
		http.StatusOK,
		dto.GenericResponseDTO{
			Code:    http.StatusOK,
			Message: "OK",
		},
	)
}

func SendGenericResponseWithData(ctx *gin.Context, data interface{}) {
	ctx.JSON(
		http.StatusOK,
		dto.GenericResponseWithDataDTO{
			Status: dto.GenericResponseDTO{
				Code:    http.StatusOK,
				Message: "OK",
			},
			Data: data,
		},
	)
}
