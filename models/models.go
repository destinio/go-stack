package models

import "time"

type Todo struct {
	Id        uint `gorm:"primaryKey"`
	Title     string
	Completed bool `gorm:"default:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
