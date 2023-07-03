package modules

import (
	"gorm.io/gorm"
)

type Bromo1Repo struct {
	db *gorm.DB
}

func CreateBromo1Repo(gorm *gorm.DB) Bromo1Repo {
	return Bromo1Repo{
		db: gorm,
	}
}
