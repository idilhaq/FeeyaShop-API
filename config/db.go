package config

import (
	"database/sql"
	"feeyashop/models"
	"feeyashop/utils"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
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
	environment := utils.Getenv("ENVIRONMENT", "development")

	if environment == "production" {
		username := os.Getenv("DATABASE_USERNAME")
		password := os.Getenv("DATABASE_PASSWORD")
		host := os.Getenv("DATABASE_HOST")
		port := os.Getenv("DATABASE_PORT")
		database := os.Getenv("DATABASE_NAME")
		// production
		dsn := "host=" + host + " user=" + username + " password=" + password + " dbname=" + database + " port=" + port + " sslmode=require"
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err.Error())
		}

		db.AutoMigrate(
			&models.Category{},
			&models.Product{},
			&models.User{},
			&models.Access{},
			&models.Cart{},
			&models.Comment{},
			&models.Like{},
			&models.Purchase{},
			&models.Rating{})

		return db
	} else {
		// development
		host := "tcp(127.0.0.1:3306)"

		dsn := fmt.Sprintf("%v:%v@%v/%v?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, database)
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

		if err != nil {
			panic(err.Error())
		}

		db.AutoMigrate(
			&models.Category{},
			&models.Product{},
			&models.User{},
			&models.Access{},
			&models.Cart{},
			&models.Comment{},
			&models.Like{},
			&models.Purchase{},
			&models.Rating{})

		return db
	}
}
