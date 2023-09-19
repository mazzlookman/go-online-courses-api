package domain

import "time"

type User struct {
	Id          int `gorm:"primaryKey"`
	Name        string
	Email       string
	Password    string
	Avatar      string
	Token       string
	CreatedAt   time.Time `gorm:"type:TIMESTAMP;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `gorm:"type:TIMESTAMP;not null;default:CURRENT_TIMESTAMP"`
	Courses     []Course  `gorm:"many2many:user_courses;"`
	Transaction []Transaction
}
