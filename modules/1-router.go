package modules

import "github.com/gin-gonic/gin"

type Bromo1Router struct {
	engine *gin.Engine
	rh     Bromo1RequestHandler
}

func NewBromo1Router(engine *gin.Engine, rh Bromo1RequestHandler) Bromo1Router {
	return Bromo1Router{
		engine: engine,
		rh:     rh,
	}
}
