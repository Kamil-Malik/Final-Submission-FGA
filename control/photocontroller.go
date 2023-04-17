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

func CreatePhoto(ctx *gin.Context) {
	photoDTO := dto.PhotoDTO{}
	contentType := helper.GetContentType(ctx)
	if contentType == "application/json" {
		if err := ctx.ShouldBindJSON(&photoDTO); err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest,
				dto.GenericResponseDTO{
					Code:    http.StatusBadRequest,
					Message: "Please provide a valid data",
				})
			return
		}
	} else {
		if err := ctx.ShouldBind(&photoDTO); err != nil {
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
	photoDTO.UserID = uint(userData["id"].(float64))

	if _, err := valid.ValidateStruct(photoDTO); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dto.GenericResponseDTO{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	if err := service.InsertPhoto(mapper.PhotoDTOToEntity(photoDTO)); err != nil {
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

func GetAllPhoto(ctx *gin.Context) {
	photoEntities := service.GetAllPhoto()
	photosDto := []dto.PhotoDTO{}
	for _, entity := range photoEntities {
		photosDto = append(photosDto, mapper.PhotoEntityToDTO(entity))
	}

	ctx.JSON(
		http.StatusOK,
		gin.H{
			"status": dto.GenericResponseDTO{
				Code:    http.StatusOK,
				Message: "OK",
			},
			"data": photosDto,
		},
	)
}

func GetPhotoByID(ctx *gin.Context) {
	id := ctx.Param("id")
	photoID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			dto.GenericResponseDTO{
				Code:    http.StatusNotFound,
				Message: "Please provide a valid ID",
			})
		return
	}

	photoEntity, err := service.GetPhotoByID(uint(photoID))
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			dto.GenericResponseDTO{
				Code:    http.StatusNotFound,
				Message: "Photo not found",
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
			"data": mapper.PhotoEntityToDTO(photoEntity),
		},
	)
}

func UpdatePhotoByID(ctx *gin.Context) {
	id := ctx.Param("id")
	photoID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			dto.GenericResponseDTO{
				Code:    http.StatusNotFound,
				Message: "Please provide a valid ID",
			})
		return
	}

	photoDTO := dto.PhotoDTO{}
	contentType := helper.GetContentType(ctx)
	if contentType == "application/json" {
		if err := ctx.ShouldBindJSON(&photoDTO); err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest,
				dto.GenericResponseDTO{
					Code:    http.StatusBadRequest,
					Message: "Please provide a valid data",
				})
			return
		}
	} else {
		if err := ctx.ShouldBind(&photoDTO); err != nil {
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
	photoDTO.UserID = uint(userData["id"].(float64))
	photoDTO.ID = uint(photoID)

	if _, err := valid.ValidateStruct(photoDTO); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dto.GenericResponseDTO{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	if err := service.UpdatePhoto(mapper.PhotoDTOToEntity(photoDTO)); err != nil {
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

func DeletePhotoByID(ctx *gin.Context) {
	id := ctx.Param("id")
	photoID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			dto.GenericResponseDTO{
				Code:    http.StatusNotFound,
				Message: "Please provide a valid ID",
			})
		return
	}

	if err := service.DeletePhotoByID(uint(photoID)); err != nil {
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
