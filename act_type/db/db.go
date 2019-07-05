package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func ConnectToDb() *gorm.DB{
	db, err := gorm.Open("mysql", "root:251719952@/act1?charset=utf8&parseTime=True&loc=Local")

	if err!= nil {
		log.Fatal("Couldn't connect to database", err.Error())
	}

	return db
}
