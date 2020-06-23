package db

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql" //nolint

	"github.com/jinzhu/gorm"
)

func InitDB() *gorm.DB {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	fmt.Println(dataSourceName)

	db, err := gorm.Open(os.Getenv("DB_DRIVER"), dataSourceName)
	if err != nil {
		panic(err.Error())
	}

	return db
}
