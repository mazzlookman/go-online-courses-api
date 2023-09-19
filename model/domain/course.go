package domain

import "time"

type Course struct {
	Id          int `gorm:"primaryKey"`
	AuthorId    int
	Title       string
	Slug        string
	Description string `gorm:"type:text"`
	Perks       string `gorm:"type:text"`
	Price       int    `gorm:"default:0;not null"`
	Banner      string
	CreatedAt   time.Time `gorm:"type:TIMESTAMP;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `gorm:"type:TIMESTAMP;not null;default:CURRENT_TIMESTAMP"`
	Users       []User    `gorm:"many2many:user_courses;"`
	Author      Author
	Transaction []Transaction
}
