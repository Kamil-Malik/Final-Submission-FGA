package dto

type GenericResponseDTO struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type GenericResponseWithDataDTO struct {
	Status GenericResponseDTO `json:"status"`
	Data   interface{}        `json:"data"`
}
