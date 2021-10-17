package config

import (
	"database/sql"
	"feeyashop/models"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	username string = "feeyaadmin"
	password string = "feeyashop"
	database string = "feeyashop"
)

var dsn = fmt.Sprintf("%v:%v@/%v", username, password, database)

func MySQL() (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, err
	}

	return db, nil
}

func ConnectDataBase() *gorm.DB {
	host := "tcp(127.0.0.1:3306)"

	dsn := fmt.Sprintf("%v:%v@%v/%v?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, database)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&models.Category{}, &models.Product{}, &models.User{})

	return db
}
