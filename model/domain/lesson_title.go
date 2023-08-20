package domain

import "time"

type LessonTitle struct {
	ID             int `gorm:"primaryKey"`
	CourseID       int
	Title          string `gorm:"not null;default:untitled"`
	InOrder        int
	CreatedAt      time.Time `gorm:"type:TIMESTAMP;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt      time.Time `gorm:"type:TIMESTAMP;not null;default:CURRENT_TIMESTAMP"`
	LessonContents []LessonContent
}
