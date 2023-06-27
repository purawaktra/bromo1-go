package module

import "github.com/gin-gonic/gin"

type BromoRouter struct {
	engine *gin.Engine
	rh     BromoRequestHandler
}

func NewBromoRouter(engine *gin.Engine, rh BromoRequestHandler) BromoRouter {
	return BromoRouter{
		engine: engine,
		rh:     rh,
	}
}
