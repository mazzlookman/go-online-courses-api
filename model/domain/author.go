package domain

import "time"

type Author struct {
	ID        int `gorm:"primaryKey"`
	Name      string
	Email     string
	Password  string
	Profile   string `gorm:"type:text"`
	Avatar    string
	Token     string
	CreatedAt time.Time `gorm:"type:TIMESTAMP;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"type:TIMESTAMP;not null;default:CURRENT_TIMESTAMP"`
	Courses   []Course
}
