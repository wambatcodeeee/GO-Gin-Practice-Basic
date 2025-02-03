package network

import (
	"github.com/gin-gonic/gin"
	"goweb1/internal/service"
	"goweb1/internal/util"
	"sync"
)

var (
	foodRouterInit     sync.Once //라우터 한번만 호출하기 위함
	foodRouterInstance *foodRouter
)

type foodRouter struct {
	router      *Network
	foodService *service.Food
}

func newFoodRouter(router *Network, foodService *service.Food) *foodRouter {
	foodRouterInit.Do(func() {
		foodRouterInstance = &foodRouter{
			router:      router,
			foodService: foodService,
		}
		router.registerGET("/", foodRouterInstance.get)
		router.registerPOST("/", foodRouterInstance.create)
		router.registerUPDATE("/", foodRouterInstance.update)
		router.registerDELETE("/", foodRouterInstance.delete)
	})

	return foodRouterInstance
}

func (f *foodRouter) create(c *gin.Context) {
	f.foodService.Create(nil)
	f.router.okResponse(c, &util.CreateFoodResponse{
		ApiResponse: util.NewApiResponse("성공입니다.", 1),
	})
}

func (f *foodRouter) get(c *gin.Context) {
	f.foodService.Get()
	f.router.okResponse(c, &util.GetFoodResponse{
		ApiResponse: util.NewApiResponse("성공입니다.", 1),
		Foods:       f.foodService.Get(),
	})
}

func (f *foodRouter) update(c *gin.Context) {
	f.foodService.Update(nil, nil)
	f.router.okResponse(c, &util.UpdateFoodResponse{
		ApiResponse: util.NewApiResponse("성공입니다.", 1),
	})
}

func (f *foodRouter) delete(c *gin.Context) {
	f.foodService.Delete(nil)
	f.router.okResponse(c, &util.DeleteFoodResponse{
		ApiResponse: util.NewApiResponse("성공입니다.", 1),
	})
}
