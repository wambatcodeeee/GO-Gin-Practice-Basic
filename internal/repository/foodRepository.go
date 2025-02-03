package repository

import "goweb1/internal/model"

type FoodRepository struct {
	foodMap []*model.Food
}

func NewFoodRepository() *FoodRepository {
	return &FoodRepository{
		foodMap: []*model.Food{},
	}
}

func (f *FoodRepository) Create(newFood *model.Food) error {
	return nil
}

func (f *FoodRepository) Get() []*model.Food {
	return f.foodMap
}

func (f *FoodRepository) Update(oldFood, newFood *model.Food) error {
	return nil
}

func (f *FoodRepository) Delete(newFood *model.Food) error {
	return nil
}
