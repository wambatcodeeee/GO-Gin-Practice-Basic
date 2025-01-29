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
