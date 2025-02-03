package repository

import (
	"sync"
)

var (
	repositoryInit     sync.Once //라우터 한번만 호출하기 위함
	repositoryInstance *Repository
)

type Repository struct {
	Food *FoodRepository
}

func NewRepository() *Repository {
	repositoryInit.Do(func() {
		repositoryInstance = &Repository{
			Food: NewFoodRepository(),
		}
	})

	return repositoryInstance
}
