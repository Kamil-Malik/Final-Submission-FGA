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
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest,
				dto.GenericResponseDTO{
					Code:    http.StatusBadRequest,
					Message: "Please provide a valid data",
				})
			return
		}
	} else {
		if err := ctx.ShouldBind(&socialMediaDTO); err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest,
				dto.GenericResponseDTO{
					Code:    http.StatusBadRequest,
					Message: "Please provide a valid data",
				})
			return
		}
	}

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	socialMediaDTO.UserID = uint(userData["id"].(float64))

	if _, err := valid.ValidateStruct(socialMediaDTO); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dto.GenericResponseDTO{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	if err := service.InsertSocialMedia(mapper.SocialMediaDTOToEntity(socialMediaDTO)); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dto.GenericResponseDTO{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(
		http.StatusCreated,
		dto.GenericResponseDTO{
			Code:    http.StatusCreated,
			Message: "OK",
		})
}

func GetAllSocialMedia(ctx *gin.Context) {
	socialMediaEntites := service.GetAllSocialMedia()
	socialMediasDTO := []dto.SocialMediaDTO{}
	for _, entity := range socialMediaEntites {
		socialMediasDTO = append(socialMediasDTO, mapper.SocialMediaEntityToDto(entity))
	}

	ctx.JSON(
		http.StatusOK,
		gin.H{
			"status": dto.GenericResponseDTO{
				Code:    http.StatusOK,
				Message: "OK",
			},
			"data": socialMediasDTO,
		},
	)
}

func GetSocialMediaByID(ctx *gin.Context) {
	id := ctx.Param("id")
	socialMediaID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			dto.GenericResponseDTO{
				Code:    http.StatusNotFound,
				Message: "Please provide a valid ID",
			})
		return
	}

	socialMediaEntity, err := service.GetSocialMediaByID(uint(socialMediaID))
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			dto.GenericResponseDTO{
				Code:    http.StatusNotFound,
				Message: "Social Media not found",
			})
		return
	}

	ctx.JSON(
		http.StatusOK,
		gin.H{
			"status": dto.GenericResponseDTO{
				Code:    http.StatusOK,
				Message: "OK",
			},
			"data": mapper.SocialMediaEntityToDto(socialMediaEntity),
		},
	)
}

func UpdateSocialMediaByID(ctx *gin.Context) {
	id := ctx.Param("id")
	socialMediaID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			dto.GenericResponseDTO{
				Code:    http.StatusNotFound,
				Message: "Please provide a valid ID",
			})
		return
	}

	socialMediaDTO := dto.SocialMediaDTO{}
	contentType := helper.GetContentType(ctx)
	if contentType == "application/json" {
		if err := ctx.ShouldBindJSON(&socialMediaDTO); err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest,
				dto.GenericResponseDTO{
					Code:    http.StatusBadRequest,
					Message: "Please provide a valid data",
				})
			return
		}
	} else {
		if err := ctx.ShouldBind(&socialMediaDTO); err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest,
				dto.GenericResponseDTO{
					Code:    http.StatusBadRequest,
					Message: "Please provide a valid data",
				})
			return
		}
	}

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	socialMediaDTO.UserID = uint(userData["id"].(float64))
	socialMediaDTO.ID = uint(socialMediaID)

	if _, err := valid.ValidateStruct(socialMediaDTO); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dto.GenericResponseDTO{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	if err := service.UpdateSocialMedia(mapper.SocialMediaDTOToEntity(socialMediaDTO)); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dto.GenericResponseDTO{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(
		http.StatusOK,
		dto.GenericResponseDTO{
			Code:    http.StatusOK,
			Message: "OK",
		})
}

func DeleteSocialMediaByID(ctx *gin.Context) {
	id := ctx.Param("id")
	socialMediaID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			dto.GenericResponseDTO{
				Code:    http.StatusBadRequest,
				Message: "Please provide a valid ID",
			})
		return
	}

	if err := service.DeletePhotoByID(uint(socialMediaID)); err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			dto.GenericResponseDTO{
				Code:    http.StatusNotFound,
				Message: "Photo not found",
			})
		return
	}

	ctx.JSON(
		http.StatusCreated,
		dto.GenericResponseDTO{
			Code:    http.StatusCreated,
			Message: "OK",
		})
}
