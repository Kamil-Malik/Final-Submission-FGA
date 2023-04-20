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

// Create Photo
// @Summary Create Photo
// @Schemes
// @Description Create Photo
// @Tags photo
// @Accept json
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Success 200 {object} dto.GenericResponseDTO
// @Failure 400 {object} dto.GenericResponseDTO
// @Failure 401 {object} dto.GenericResponseDTO
// @Router /photo/ [post]
func CreatePhoto(ctx *gin.Context) {
	photoDTO := dto.PhotoDTO{}
	contentType := helper.GetContentType(ctx)
	if contentType == "application/json" {
		if err := ctx.ShouldBindJSON(&photoDTO); err != nil {
			helper.AbortGenericResponse(
				ctx,
				http.StatusBadRequest,
				"Please provide a valid data",
			)
			return
		}
	} else {
		if err := ctx.ShouldBind(&photoDTO); err != nil {
			helper.AbortGenericResponse(
				ctx,
				http.StatusBadRequest,
				"Please provide a valid data",
			)
			return
		}
	}

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	photoDTO.UserID = uint(userData["id"].(float64))

	if _, err := valid.ValidateStruct(photoDTO); err != nil {
		helper.AbortGenericResponse(
			ctx,
			http.StatusBadRequest,
			err.Error(),
		)
		return
	}

	if err := service.InsertPhoto(mapper.PhotoDTOToEntity(photoDTO)); err != nil {
		helper.AbortGenericResponse(
			ctx,
			http.StatusBadRequest,
			err.Error(),
		)
		return
	}

	helper.SendGenericResponse(ctx)
}

// Get All Photo
// @Summary Get All Photo
// @Schemes
// @Description Get ALl Photo
// @Tags photo
// @Produce json
// @Success 200 {object} dto.GenericResponseWithDataDTO
// @Failure 401 {object} dto.GenericResponseDTO
// @Router /photo/ [get]
func GetAllPhoto(ctx *gin.Context) {
	photoEntities := service.GetAllPhoto()
	photosDto := []dto.PhotoDTO{}
	for _, entity := range photoEntities {
		photosDto = append(photosDto, mapper.PhotoEntityToDTO(entity))
	}

	helper.SendGenericResponseWithData(ctx, photosDto)
}

// Get Photo By ID
// @Summary Get Photo By ID
// @Schemes
// @Description Get Photo By ID
// @Tags comment
// @Produce json
// @Success 200 {object} dto.GenericResponseWithDataDTO
// @Failure 400 {object} dto.GenericResponseDTO
// @Failure 401 {object} dto.GenericResponseDTO
// @Failure 404 {object} dto.GenericResponseDTO
// @Router /photo/:id [get]
func GetPhotoByID(ctx *gin.Context) {
	id := ctx.Param("id")
	photoID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		helper.AbortGenericResponse(
			ctx,
			http.StatusBadRequest,
			"Please provide a valid ID",
		)
		return
	}

	photoEntity, err := service.GetPhotoByID(uint(photoID))
	if err != nil {
		helper.AbortGenericResponse(
			ctx,
			http.StatusNotFound,
			"Photo not found",
		)
		return
	}

	helper.SendGenericResponseWithData(
		ctx,
		mapper.PhotoEntityToDTO(photoEntity),
	)
}

// Update Photo By ID
// @Summary Update Photo By Specific ID
// @Schemes
// @Description Update Photo By Specific ID
// @Tags comment
// @Accept json
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Success 200 {object} dto.GenericResponseDTO
// @Failure 400 {object} dto.GenericResponseDTO
// @Failure 401 {object} dto.GenericResponseDTO
// @Failure 404 {object} dto.GenericResponseDTO
// @Router /photo/:id [put]
func UpdatePhotoByID(ctx *gin.Context) {
	id := ctx.Param("id")
	photoID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		helper.AbortGenericResponse(
			ctx,
			http.StatusBadRequest,
			"Please provide a valid ID",
		)
		return
	}

	photoDTO := dto.PhotoDTO{}
	contentType := helper.GetContentType(ctx)
	if contentType == "application/json" {
		if err := ctx.ShouldBindJSON(&photoDTO); err != nil {
			helper.AbortGenericResponse(
				ctx,
				http.StatusBadRequest,
				"Please provide a valid data",
			)
			return
		}
	} else {
		if err := ctx.ShouldBind(&photoDTO); err != nil {
			helper.AbortGenericResponse(
				ctx,
				http.StatusBadRequest,
				"Please provide a valid data",
			)
			return
		}
	}

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	photoDTO.UserID = uint(userData["id"].(float64))
	photoDTO.ID = uint(photoID)

	if _, err := valid.ValidateStruct(photoDTO); err != nil {
		helper.AbortGenericResponse(
			ctx,
			http.StatusBadRequest,
			err.Error(),
		)
		return
	}

	if err := service.UpdatePhoto(mapper.PhotoDTOToEntity(photoDTO)); err != nil {
		helper.AbortGenericResponse(
			ctx,
			http.StatusBadRequest,
			err.Error(),
		)
		return
	}

	helper.SendGenericResponse(ctx)
}

// Delete Photo By ID
// @Summary Delete Photo By Specific ID
// @Schemes
// @Description Get Photo By Specific ID
// @Tags comment
// @Produce json
// @Success 200 {object} dto.GenericResponseDTO
// @Failure 400 {object} dto.GenericResponseDTO
// @Failure 401 {object} dto.GenericResponseDTO
// @Failure 404 {object} dto.GenericResponseDTO
// @Router /photo/:id [delete]
func DeletePhotoByID(ctx *gin.Context) {
	id := ctx.Param("id")
	photoID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		helper.AbortGenericResponse(
			ctx,
			http.StatusBadRequest,
			"Please provide a valid ID",
		)
		return
	}

	if err := service.DeletePhotoByID(uint(photoID)); err != nil {
		helper.AbortGenericResponse(
			ctx,
			http.StatusBadRequest,
			"Photo not found",
		)
		return
	}

	helper.SendGenericResponse(ctx)
}
