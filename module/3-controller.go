package module

type BromoController struct {
	uc BromoUsecase
}

func CreateBromoController(uc BromoUsecase) BromoController {
	return BromoController{uc: uc}
}
