package domain

import "time"

type LessonTitle struct {
	ID             int `gorm:"primaryKey"`
	CourseID       int
	Title          string `gorm:"not null;default:untitled"`
	InOrder        int
	CreatedAt      time.Time `gorm:"default:current_timestamp"`
	UpdatedAt      time.Time `gorm:"default:current_timestamp"`
	LessonContents []LessonContent
}
