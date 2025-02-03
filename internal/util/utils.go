package util

import "goweb1/internal/model"

type ApiResponse struct {
	Result      int64  `json:"result"`
	Description string `json:"description"`
}

type CreateFoodResponse struct {
	*ApiResponse `json:"apiResponse"`
	*model.Food  `json:"food"`
}

type GetFoodResponse struct {
	*ApiResponse `json:"apiResponse"`
	*model.Food  `json:"food"`
}

type UpdateFoodResponse struct {
	*ApiResponse `json:"apiResponse"`
	*model.Food  `json:"food"`
}

type DeleteFoodResponse struct {
	*ApiResponse `json:"apiResponse"`
	*model.Food  `json:"food"`
}

func NewApiResponse(description string, result int64) *ApiResponse {
	return &ApiResponse{
		Result:      result,
		Description: description,
	}
}
