package domain

import "time"

type LessonTitle struct {
	Id             int `gorm:"primaryKey"`
	CourseId       int
	Title          string `gorm:"not null;default:untitled"`
	InOrder        int
	CreatedAt      time.Time `gorm:"type:TIMESTAMP;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt      time.Time `gorm:"type:TIMESTAMP;not null;default:CURRENT_TIMESTAMP"`
	LessonContents []LessonContent
}
