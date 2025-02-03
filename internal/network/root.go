package network

import (
	"github.com/gin-gonic/gin"
	"goweb1/internal/service"
	"net/http"
)

type Network struct {
	engine  *gin.Engine
	service *service.Service
}

func SetupNetwork(service *service.Service) *Network {
	r := &Network{
		engine: gin.Default(),
	}

	newFoodRouter(r)

	return r
}

func (n *Network) RunServer(port string) error {
	return n.engine.Run(port)
}

func (n *Network) registerGET(path string, handler ...gin.HandlerFunc) gin.IRoutes {
	return n.engine.GET(path, handler...)
}

func (n *Network) registerPOST(path string, handler ...gin.HandlerFunc) gin.IRoutes {
	return n.engine.POST(path, handler...)
}

func (n *Network) registerUPDATE(path string, handler ...gin.HandlerFunc) gin.IRoutes {
	return n.engine.PUT(path, handler...)
}

func (n *Network) registerDELETE(path string, handler ...gin.HandlerFunc) gin.IRoutes {
	return n.engine.DELETE(path, handler...)
}

func (n *Network) okResponse(c *gin.Context, result interface{}) {
	c.JSON(http.StatusOK, result)
}

func (n *Network) badRequestResponse(c *gin.Context, result interface{}) {
	c.JSON(http.StatusBadRequest, result)
}
