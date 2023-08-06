package domain

import "time"

type LessonContent struct {
	ID            int `gorm:"primaryKey"`
	LessonTitleID int
	Content       string `gorm:"not null;default:unknown"`
	Duration      string
	InOrder       int       `gorm:"type:int"`
	CreatedAt     time.Time `gorm:"default:current_timestamp"`
	UpdatedAt     time.Time `gorm:"default:current_timestamp"`
}
