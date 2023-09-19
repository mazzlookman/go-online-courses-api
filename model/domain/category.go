package domain

import "time"

type Category struct {
	Id        int `gorm:"primaryKey"`
	Name      string
	CreatedAt time.Time `gorm:"type:TIMESTAMP;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"type:TIMESTAMP;not null;default:CURRENT_TIMESTAMP"`
	Courses   []Course  `gorm:"many2many:category_courses"`
}
