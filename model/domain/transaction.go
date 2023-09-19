package domain

import "time"

type Transaction struct {
	Id         int `gorm:"primaryKey"`
	UserId     int
	CourseId   int
	Amount     int
	Status     string
	PaymentUrl string
	CreatedAt  time.Time `gorm:"type:TIMESTAMP;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt  time.Time `gorm:"type:TIMESTAMP;not null;default:CURRENT_TIMESTAMP"`
}
