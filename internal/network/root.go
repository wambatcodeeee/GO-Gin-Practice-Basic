package network

import "github.com/gin-gonic/gin"

type Network struct {
	engine *gin.Engine
}

func SetupNetwork() *Network {
	r := &Network{
		engine: gin.Default(),
	}

	return r
}

func (n *Network) RunServer(port string) error {
	return n.engine.Run(port)
}
