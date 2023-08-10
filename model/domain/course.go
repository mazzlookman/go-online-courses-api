package domain

import "time"

type Course struct {
	ID          int `gorm:"primaryKey"`
	AuthorID    int
	Title       string
	Slug        string
	Description string `gorm:"type:text"`
	Perks       string `gorm:"type:text"`
	Price       int    `gorm:"default:0;not null"`
	Banner      string
	CreatedAt   time.Time `gorm:"default:current_timestamp"`
	UpdatedAt   time.Time `gorm:"default:current_timestamp"`
	Users       []User    `gorm:"many2many:user_courses;"`
	Author      Author
}
