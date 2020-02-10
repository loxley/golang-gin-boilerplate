package db

import (
	_ "github.com/go-sql-driver/mysql"

	"basic_server/server/model"
	"fmt"
	"github.com/jinzhu/gorm"
	"os"
)

func InitDB() *gorm.DB {
	dataSourceName := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"))
	db, err := gorm.Open(os.Getenv("DB_DRIVER"), dataSourceName)
	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&model.Post{})

	return db
}