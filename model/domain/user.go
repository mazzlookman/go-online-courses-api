package domain

import "time"

type User struct {
	ID        int `gorm:"primaryKey"`
	Name      string
	Email     string
	Password  string
	Avatar    string
	Token     string
	CreatedAt time.Time `gorm:"default:current_timestamp;"`
	UpdatedAt time.Time `gorm:"default:current_timestamp;"`
	Courses   []Course  `gorm:"many2many:user_courses;"`
}
