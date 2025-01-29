package network

import (
	"github.com/gin-gonic/gin"
	"goweb1/internal/model"
	"goweb1/internal/util"
	"sync"
)

var (
	foodRouterInit     sync.Once //라우터 한번만 호출하기 위함
	foodRouterInstance *foodRouter
)

type foodRouter struct {
	router *Network
}

func newFoodRouter(router *Network) *foodRouter {
	foodRouterInit.Do(func() {
		foodRouterInstance = &foodRouter{router: router}
		router.registerGET("/", foodRouterInstance.get)
		router.registerPOST("/", foodRouterInstance.create)
		router.registerUPDATE("/", foodRouterInstance.update)
		router.registerDELETE("/", foodRouterInstance.delete)
	})

	return foodRouterInstance
}

func (f *foodRouter) create(c *gin.Context) {

}

func (f *foodRouter) get(c *gin.Context) {
	f.router.okResponse(c, &model.FoodResponse{
		ApiResponse: &util.ApiResponse{
			Result:      1,
			Description: "성공입니다.",
		},
		Food: nil,
	})
}

func (f *foodRouter) update(c *gin.Context) {

}

func (f *foodRouter) delete(c *gin.Context) {

}
