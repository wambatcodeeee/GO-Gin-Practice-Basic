package service

import (
	"goweb1/internal/repository"
	"sync"
)

var (
	serviceInit     sync.Once //라우터 한번만 호출하기 위함
	serviceInstance *Service
)

type Service struct {
	repository *repository.Repository
	Food       *Food
}

func NewService(rep *repository.Repository) *Service {
	serviceInit.Do(func() {
		serviceInstance = &Service{
			repository: rep,
		}
		serviceInstance.Food = newFoodService(rep.Food)
	})

	return serviceInstance
}
