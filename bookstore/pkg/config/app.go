package config

import (
	"log"

	"github.com/rafaelmgr12/bookstore/pkg/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func DatabaseConnection() {
	connectionString := "host=database port=3306 user=rafael password=secret dbname=bookstore sslmode=disable"
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}
	DB.AutoMigrate(&models.Book{})
}
