package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/purawaktra/bromo1-go/modules"
	"github.com/purawaktra/bromo1-go/utils"
)

func main() {
	utils.InitConfig()
	utils.InitLog()

	utils.Log("============= BROMO1 RUNTIME STARTED =============")
	// create gin engine
	engine := gin.New()

	// create dsn
	dsn := fmt.Sprintf("mongodb://%s:%s@%s:%s/?authSource=%s",
		utils.MongoDBUser,
		utils.MongoDBPassword,
		utils.MongoDBHost,
		utils.MongoDBPort,
		utils.MongoDBDatabase)

	// create mongodb instance
	mongoDbInstance, err := utils.CreateMongoDb(dsn)
	if err != nil {
		utils.Error(err, "main", "")
		panic(err)
	}

	// declare architecture
	repository := modules.CreateBromo1Repo(mongoDbInstance)
	usecase := modules.CreateBromo1Usecase(repository)
	controller := modules.CreateBromo1Controller(usecase)
	requestHandler := modules.CreateBromo1RequestHandler(controller)
	router := modules.CreateBromo1Router(engine, requestHandler)

	// init routing
	router.Init("/bromo1/api/v1")

	utils.Log("============= BROMO1 ENGINE STARTING =============")
	utils.Log(fmt.Sprintf("App Address = %s:%s", utils.AppHost, utils.AppPort))

	// init routing to swagger
	swaggerRouter := utils.CreateSwaggerRouter(engine)
	swaggerRouter.InitRouter("/bromo1/api/v1/swagger")

	// start http api engine
	err = engine.Run(fmt.Sprintf("%s:%s", utils.AppHost, utils.AppPort))
	if err != nil {
		utils.Fatal(err, "main", "")
		panic(err)
	}

	utils.Log("============= BROMO1 ENGINE FAILED =============")

}
