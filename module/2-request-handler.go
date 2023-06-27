package module

type BromoRequestHandler struct {
	ctrl BromoController
}

func CreateBromoRequestHandler(ctrl BromoController) BromoRequestHandler {
	return BromoRequestHandler{ctrl: ctrl}
}
