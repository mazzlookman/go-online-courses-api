package domain

import "time"

type Author struct {
	ID        int `gorm:"primaryKey"`
	Name      string
	Email     string
	Password  string
	Profile   string `gorm:"type:text"`
	Avatar    string
	CreatedAt time.Time `gorm:"default:current_timestamp"`
	UpdatedAt time.Time `gorm:"default:current_timestamp"`
	Courses   []Course
}
