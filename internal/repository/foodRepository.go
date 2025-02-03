package repository

import "goweb1/internal/model"

type FoodRepository struct {
	foodMap []*model.Food
}

func NewFoodRepository() *FoodRepository {
	return &FoodRepository{}
}
