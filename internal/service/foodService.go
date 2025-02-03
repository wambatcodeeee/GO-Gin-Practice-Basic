package service

import (
	"goweb1/internal/model"
	"goweb1/internal/repository"
)

type Food struct {
	foodRepository *repository.FoodRepository
}

func newFoodService(foodRepository *repository.FoodRepository) *Food {
	return &Food{foodRepository: foodRepository}
}

func (f *Food) Create(newFood *model.Food) error {
	return f.foodRepository.Create(newFood)
}

func (f *Food) Get() []*model.Food {
	return f.foodRepository.Get()
}

func (f *Food) Update(oldFood, newFood *model.Food) error {
	return f.foodRepository.Update(oldFood, newFood)
}

func (f *Food) Delete(newFood *model.Food) error {
	return f.foodRepository.Delete(newFood)
}
