package domain

import "time"

type LessonContent struct {
	Id            int `gorm:"primaryKey"`
	LessonTitleId int
	Content       string `gorm:"not null;default:unknown"`
	Duration      string
	InOrder       int       `gorm:"type:int"`
	CreatedAt     time.Time `gorm:"type:TIMESTAMP;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt     time.Time `gorm:"type:TIMESTAMP;not null;default:CURRENT_TIMESTAMP"`
}
