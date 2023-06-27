package utils

import (
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func openMysql(dsn string) gorm.Dialector {
	return mysql.Open(dsn)
}

func CreateGorm(dsn string) (*gorm.DB, error) {
	var db *gorm.DB
	var err error
	db, err = gorm.Open(openMysql(dsn), &gorm.Config{})
	if err != nil {
		log.Println("gorm.open", err)
		return nil, err
	}

	var healthCheck *sql.DB
	healthCheck, err = db.DB()
	if err != nil {
		return nil, err
	}

	err = healthCheck.Ping()
	if err != nil {
		return nil, err
	}

	return db, err
}
