package control

import (
	"Proyek-Akhir-Golang/dto"
	"Proyek-Akhir-Golang/helper"
	"Proyek-Akhir-Golang/mapper"
	"Proyek-Akhir-Golang/service"
	"net/http"
	"strconv"

	valid "github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func CreateSocialMedia(ctx *gin.Context) {
	socialMediaDTO := dto.SocialMediaDTO{}
	contentType := helper.GetContentType(ctx)
	if contentType == "application/json" {
		if err := ctx.ShouldBindJSON(&socialMediaDTO); err != nil {
			helper.AbortGenericResponse(
				ctx,
				http.StatusBadRequest,
				"Please provide a valid data",
			)
			return
		}
	} else {
		if err := ctx.ShouldBind(&socialMediaDTO); err != nil {
			helper.AbortGenericResponse(
				ctx,
				http.StatusBadRequest,
				"Please provide a valid data",
			)
			return
		}
	}

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	socialMediaDTO.UserID = uint(userData["id"].(float64))

	if _, err := valid.ValidateStruct(socialMediaDTO); err != nil {
		helper.AbortGenericResponse(
			ctx,
			http.StatusBadRequest,
			err.Error(),
		)
		return
	}

	if err := service.InsertSocialMedia(mapper.SocialMediaDTOToEntity(socialMediaDTO)); err != nil {
		helper.AbortGenericResponse(
			ctx,
			http.StatusBadRequest,
			err.Error(),
		)
		return
	}

	helper.SendGenericResponse(ctx)
}

func GetAllSocialMedia(ctx *gin.Context) {
	socialMediaEntites := service.GetAllSocialMedia()
	socialMediasDTO := []dto.SocialMediaDTO{}
	for _, entity := range socialMediaEntites {
		socialMediasDTO = append(socialMediasDTO, mapper.SocialMediaEntityToDto(entity))
	}

	helper.SendGenericResponseWithData(ctx, socialMediasDTO)
}

func GetSocialMediaByID(ctx *gin.Context) {
	id := ctx.Param("id")
	socialMediaID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		helper.AbortGenericResponse(
			ctx,
			http.StatusBadRequest,
			"Please provide a valid ID",
		)
		return
	}

	socialMediaEntity, err := service.GetSocialMediaByID(uint(socialMediaID))
	if err != nil {
		helper.AbortGenericResponse(
			ctx,
			http.StatusBadRequest,
			"Social Media not found",
		)
		return
	}

	helper.SendGenericResponseWithData(
		ctx,
		mapper.SocialMediaEntityToDto(socialMediaEntity),
	)
}

func UpdateSocialMediaByID(ctx *gin.Context) {
	id := ctx.Param("id")
	socialMediaID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		helper.AbortGenericResponse(
			ctx,
			http.StatusBadRequest,
			"Please provide a valid ID",
		)
		return
	}

	socialMediaDTO := dto.SocialMediaDTO{}
	contentType := helper.GetContentType(ctx)
	if contentType == "application/json" {
		if err := ctx.ShouldBindJSON(&socialMediaDTO); err != nil {
			helper.AbortGenericResponse(
				ctx,
				http.StatusBadRequest,
				"Please provide a valid data",
			)
			return
		}
	} else {
		if err := ctx.ShouldBind(&socialMediaDTO); err != nil {
			helper.AbortGenericResponse(
				ctx,
				http.StatusBadRequest,
				"Please provide a valid data",
			)
			return
		}
	}

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	socialMediaDTO.UserID = uint(userData["id"].(float64))
	socialMediaDTO.ID = uint(socialMediaID)

	if _, err := valid.ValidateStruct(socialMediaDTO); err != nil {
		helper.AbortGenericResponse(
			ctx,
			http.StatusBadRequest,
			err.Error(),
		)
		return
	}

	if err := service.UpdateSocialMedia(mapper.SocialMediaDTOToEntity(socialMediaDTO)); err != nil {
		helper.AbortGenericResponse(
			ctx,
			http.StatusBadRequest,
			err.Error(),
		)
		return
	}

	helper.SendGenericResponse(ctx)
}

func DeleteSocialMediaByID(ctx *gin.Context) {
	id := ctx.Param("id")
	socialMediaID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		helper.AbortGenericResponse(
			ctx,
			http.StatusBadRequest,
			"Please provide a valid ID",
		)
		return
	}

	if err := service.DeletePhotoByID(uint(socialMediaID)); err != nil {
		helper.AbortGenericResponse(
			ctx,
			http.StatusBadRequest,
			"Social media not found",
		)
		return
	}

	helper.SendGenericResponse(ctx)
}
