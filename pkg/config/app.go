package config

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))

	if err != nil {
		log.Panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
