package domain

import "time"

type Category struct {
	ID        int `gorm:"primaryKey"`
	Name      string
	CreatedAt time.Time `gorm:"default:current_timestamp"`
	UpdatedAt time.Time `gorm:"default:current_timestamp"`
	Courses   []Course  `gorm:"many2many:category_courses"`
}
