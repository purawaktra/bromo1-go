package modules

type Bromo1Controller struct {
	uc Bromo1Usecase
}

func CreateBromo1Controller(uc Bromo1Usecase) Bromo1Controller {
	return Bromo1Controller{uc: uc}
}
