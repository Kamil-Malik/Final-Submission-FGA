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

// Create Social Media
// @Summary Create Social Media
// @Schemes
// @Description Create Social Media and link it with User Account
// @Tags social_media
// @Produce json
// @Success 200 {object} dto.GenericResponseDTO
// @Failure 400 {object} dto.GenericResponseDTO
// @Failure 401 {object} dto.GenericResponseDTO
// @Router /social/ [post]
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

// Get All Social Media
// @Summary Get All Social Media
// @Schemes
// @Description Get All Social Media
// @Tags social_media
// @Produce json
// @Success 200 {object} dto.GenericResponseWithDataDTO
// @Failure 401 {object} dto.GenericResponseDTO
// @Router /social/ [get]
func GetAllSocialMedia(ctx *gin.Context) {
	socialMediaEntites := service.GetAllSocialMedia()
	socialMediasDTO := []dto.SocialMediaDTO{}
	for _, entity := range socialMediaEntites {
		socialMediasDTO = append(socialMediasDTO, mapper.SocialMediaEntityToDto(entity))
	}

	helper.SendGenericResponseWithData(ctx, socialMediasDTO)
}

// Get Social Media By ID
// @Summary Create Social Media
// @Schemes
// @Description Get Social Media with specific ID
// @Tags social_media
// @Produce json
// @Success 200 {object} dto.GenericResponseWithDataDTO
// @Failure 400 {object} dto.GenericResponseDTO
// @Failure 401 {object} dto.GenericResponseDTO
// @Router /social/:id [get]
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

// Update Social Media By ID
// @Summary Update Social Media
// @Schemes
// @Description Update Social Media with specific ID
// @Tags social_media
// @Accept json
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Success 200 {object} dto.GenericResponseDTO
// @Failure 400 {object} dto.GenericResponseDTO
// @Failure 401 {object} dto.GenericResponseDTO
// @Failure 404 {object} dto.GenericResponseDTO
// @Router /social/:id [put]
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
			http.StatusNotFound,
			err.Error(),
		)
		return
	}

	helper.SendGenericResponse(ctx)
}

// Delete Social Media By ID
// @Summary Delete Social Media
// @Schemes
// @Description Delete Social Media with specific ID
// @Tags social_media
// @Produce json
// @Success 200 {object} dto.GenericResponseDTO
// @Failure 400 {object} dto.GenericResponseDTO
// @Failure 401 {object} dto.GenericResponseDTO
// @Failure 404 {object} dto.GenericResponseDTO
// @Router /social/:id [delete]
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
			http.StatusNotFound,
			"Social media not found",
		)
		return
	}

	helper.SendGenericResponse(ctx)
}
