package domain

import "time"

type Author struct {
	ID           int `gorm:"primaryKey"`
	Name         string
	Introduction string    `gorm:"type:text"`
	CreatedAt    time.Time `gorm:"default:current_timestamp"`
	UpdatedAt    time.Time `gorm:"default:current_timestamp"`
	Courses      []Course
}
