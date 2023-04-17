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

func CreateComment(ctx *gin.Context) {
	commentDTO := dto.CommentDTO{}
	contentType := helper.GetContentType(ctx)
	if contentType == "application/json" {
		if err := ctx.ShouldBindJSON(&commentDTO); err != nil {
			helper.AbortGenericResponse(
				ctx,
				http.StatusBadRequest,
				"Please provide a valid data",
			)
			return
		}
	} else {
		if err := ctx.ShouldBind(&commentDTO); err != nil {
			helper.AbortGenericResponse(
				ctx,
				http.StatusBadRequest,
				"Please provide a valid data",
			)
			return
		}
	}

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	commentDTO.UserID = uint(userData["id"].(float64))

	if _, err := valid.ValidateStruct(commentDTO); err != nil {
		helper.AbortGenericResponse(
			ctx,
			http.StatusBadRequest,
			err.Error(),
		)
		return
	}

	if err := service.InsertComment(mapper.CommentDTOToEntity(commentDTO)); err != nil {
		helper.AbortGenericResponse(
			ctx,
			http.StatusBadRequest,
			err.Error(),
		)
		return
	}

	helper.SendGenericResponse(ctx)
}

func GetAllComments(ctx *gin.Context) {
	commentEntities := service.GetAllComments()
	commentsDTO := []dto.CommentDTO{}
	for _, entity := range commentEntities {
		commentsDTO = append(commentsDTO, mapper.CommentEntityToDTO(entity))
	}
	helper.SendGenericResponseWithData(ctx, commentsDTO)
}

func GetCommentByID(ctx *gin.Context) {
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

	commentEntity, err := service.GetCommentByID(uint(photoID))
	if err != nil {
		helper.AbortGenericResponse(
			ctx,
			http.StatusBadRequest,
			"Comment not found",
		)
		return
	}

	helper.SendGenericResponseWithData(ctx, commentEntity)
}

func UpdateCommentByID(ctx *gin.Context) {
	id := ctx.Param("id")
	commentID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		helper.AbortGenericResponse(
			ctx,
			http.StatusBadRequest,
			"Please provide a valid ID",
		)
		return
	}

	commentDTO := dto.CommentDTO{}
	contentType := helper.GetContentType(ctx)
	if contentType == "application/json" {
		if err := ctx.ShouldBindJSON(&commentDTO); err != nil {
			helper.AbortGenericResponse(
				ctx,
				http.StatusBadRequest,
				"Please provide a valid data",
			)
			return
		}
	} else {
		if err := ctx.ShouldBind(&commentDTO); err != nil {
			helper.AbortGenericResponse(
				ctx,
				http.StatusBadRequest,
				"Please provide a valid data",
			)
			return
		}
	}

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	commentDTO.UserID = uint(userData["id"].(float64))
	commentDTO.ID = uint(commentID)

	if _, err := valid.ValidateStruct(commentDTO); err != nil {
		helper.AbortGenericResponse(
			ctx,
			http.StatusBadRequest,
			err.Error(),
		)
		return
	}

	if err := service.UpdateComment(mapper.CommentDTOToEntity(commentDTO)); err != nil {
		helper.AbortGenericResponse(
			ctx,
			http.StatusBadRequest,
			err.Error(),
		)
		return
	}

	helper.SendGenericResponse(ctx)
}

func DeleteCommentByID(ctx *gin.Context) {
	id := ctx.Param("id")
	commentID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		helper.AbortGenericResponse(
			ctx,
			http.StatusBadRequest,
			"Please provide a valid ID",
		)
		return
	}

	if err := service.DeleteCommentByID(uint(commentID)); err != nil {
		helper.AbortGenericResponse(
			ctx,
			http.StatusBadRequest,
			"Comment not found",
		)
		return
	}

	helper.SendGenericResponse(ctx)
}
