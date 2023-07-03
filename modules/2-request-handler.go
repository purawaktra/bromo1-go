package modules

type Bromo1RequestHandler struct {
	ctrl Bromo1Controller
}

func Create1BromoRequestHandler(ctrl Bromo1Controller) Bromo1RequestHandler {
	return Bromo1RequestHandler{
		ctrl: ctrl,
	}
}
