package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// const (
// 	dsn = "host=localhost user=postgres password=postgres dbname=go_rest_api port=5432 sslmode=disable TimeZone=Asia/Shanghai"
// )

func GetDB(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	SeedDB(db)

	return db
}
