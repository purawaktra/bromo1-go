package module

type BromoUsecase struct {
	repo BromoRepo
}

func CreateBromoUsecase(repo BromoRepo) BromoUsecase {
	return BromoUsecase{repo: repo}
}
