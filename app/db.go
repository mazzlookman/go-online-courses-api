package app

import (
	"go-pzn-restful-api/model/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func DBConnection() *gorm.DB {
	config := mysql.Config{}
	config.DefaultStringSize = 255                                                                    // default size for string fields
	config.DSN = "root:@tcp(localhost:3306)/go_pzn_restful_api?charset=utf8&parseTime=True&loc=Local" // data source name for IDE

	if os.Getenv("DB_RUN") == "docker" {
		config.DSN = "root:root@tcp(mysql-db:3306)/go_pzn_restful_api?charset=utf8&parseTime=True&loc=Local" // data source name for docker
	}

	dbGorm, err := gorm.Open(mysql.New(config), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return dbGorm
}

func DBMigrate(DB *gorm.DB) error {
	err := DB.AutoMigrate(
		&domain.User{},
		&domain.Author{},
		&domain.LessonTitle{},
		&domain.LessonContent{},
		&domain.Course{},
		&domain.Category{},
		&domain.Transaction{},
	)

	if err != nil {
		return err
	}

	log.Println("Migration is successfully")
	return nil
}
