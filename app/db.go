package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-pzn-restful-api/model/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DBConnection() *gorm.DB {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:               "root:@tcp(127.0.0.1:3306)/go_pzn_restful_api?charset=utf8&parseTime=True&loc=Local", // data source name
		DefaultStringSize: 255,                                                                                  // default size for string fields
	}), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db
}

func DBMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(
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

	return nil
}

func FindUsersCourses(db *gorm.DB) (string, error) {
	users := []domain.User{}
	err := db.Preload("Courses").Find(&users).Error

	if err != nil {
		return "", err
	}

	marshal, _ := json.Marshal(users)
	usersJson := fmt.Sprintf("%s", marshal)
	return usersJson, nil
}

func FindCoursesUsers(db *gorm.DB) (string, error) {
	courses := []domain.Course{}
	err := db.Preload("Users").Find(&courses).Error

	if err != nil {
		return "", err
	}

	marshal, _ := json.Marshal(courses)
	coursesJson := bytes.NewBuffer(marshal).String()
	return coursesJson, nil
}

func FindAuthorCourses(db *gorm.DB) (string, error) {
	authors := []domain.Author{}
	err := db.Preload("Courses").Find(&authors).Error

	if err != nil {
		return "", err
	}

	marshal, _ := json.Marshal(authors)
	authorsJson := bytes.NewBuffer(marshal).String()
	return authorsJson, nil
}
