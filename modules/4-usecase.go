package modules

type Bromo1Usecase struct {
	repo Bromo1Repo
}

func CreateBromo1Usecase(repo Bromo1Repo) Bromo1Usecase {
	return Bromo1Usecase{
		repo: repo,
	}
}
