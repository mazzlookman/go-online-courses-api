package app

import (
	"go-pzn-restful-api/model/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func DBConnection() *gorm.DB {
	dbGorm, err := gorm.Open(mysql.New(mysql.Config{
		DSN: "root:root@tcp(mysql-db:3306)/go_pzn_restful_api?charset=utf8&parseTime=True&loc=Local", // data source name for docker
		//DSN:               "root:@tcp(localhost:3306)/go_pzn_restful_api?charset=utf8&parseTime=True&loc=Local", // data source name for IDE
		DefaultStringSize: 255, // default size for string fields
	}), &gorm.Config{})

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
	)

	if err != nil {
		return err
	}

	log.Println("Migration is successfully")
	return nil
}
