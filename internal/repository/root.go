package repository

import (
	"goweb1/internal/model"
	"sync"
)

var (
	repositoryInit     sync.Once //라우터 한번만 호출하기 위함
	repositoryInstance *Repository
)

type Repository struct {
	FoodMap []*model.Food
}

func NewRepository() *Repository {
	repositoryInit.Do(func() {
		repositoryInstance = &Repository{}
	})

	return repositoryInstance
}
