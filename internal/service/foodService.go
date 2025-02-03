package service

import "goweb1/internal/repository"

type Food struct {
	foodRepository *repository.FoodRepository
}

func newFoodService(foodRepository *repository.FoodRepository) *Food {
	return &Food{foodRepository: foodRepository}
}
