package modules

import (
	"github.com/gin-gonic/gin"
	"github.com/purawaktra/bromo1-go/utils"
)

type Bromo1Router struct {
	engine *gin.Engine
	rh     Bromo1RequestHandler
}

func CreateBromo1Router(engine *gin.Engine, rh Bromo1RequestHandler) Bromo1Router {
	return Bromo1Router{
		engine: engine,
		rh:     rh,
	}
}

func (r Bromo1Router) Init(path string) {
	pathGroup := r.engine.Group(path, utils.CheckBasicAuth)
	pathGroup.GET("/select/profile_picture", r.rh.RetrieveProfilePicture)
	pathGroup.POST("/insert/profile_picture", r.rh.StoreProfilePicture)
	pathGroup.DELETE("/delete/profile_picture", r.rh.RemoveProfilePicture)
}
