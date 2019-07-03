package common

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

func InitDB() *gorm.DB{
	db, err := gorm.Open("postgres", "host=68.183.179.182 port=5432 user=postgres dbname=udemy password=welcome1,. sslmode=disable")

	if err!= nil{
		fmt.Println("error while connecting db ", err)
	}
	DB = db
	return DB
}

func GetDB() *gorm.DB  {
	return DB
}