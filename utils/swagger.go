package utils

import (
	"github.com/gin-gonic/gin"
	_ "github.com/purawaktra/bromo1-go/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type SwaggerRouter struct {
	engine *gin.Engine
}

func CreateSwaggerRouter(engine *gin.Engine) SwaggerRouter {
	return SwaggerRouter{
		engine: engine,
	}
}

func (r SwaggerRouter) Init(path string) {
	//swaggerSpec, err := loads.JSONDoc("openapi.json")
	//if err != nil {
	//	panic("Failed to load OpenAPI spec")
	//}

	pathGroup := r.engine.Group(path)
	pathGroup.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
