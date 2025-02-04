package model

import "goweb1/internal/util"

type Food struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type FoodResponse struct {
	*util.ApiResponse
	*Food
}

type CreateRequest struct {
}

type CreateFoodResponse struct {
	*util.ApiResponse `json:"apiResponse"`
	*Food             `json:"food"`
}

type GetFoodResponse struct {
	*util.ApiResponse `json:"apiResponse"`
	Foods             []*Food `json:"result"`
}

type CreateRequest struct {
}

type UpdateFoodResponse struct {
	*util.ApiResponse `json:"apiResponse"`
	*Food             `json:"food"`
}

type CreateRequest struct {
}

type DeleteFoodResponse struct {
	*util.ApiResponse `json:"apiResponse"`
	*Food             `json:"food"`
}
