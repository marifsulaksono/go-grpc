package config

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/ecommerce"))
	if err != nil {
		log.Fatalf("Connect database failed : %v", err)
		return nil
	}

	return db
}
