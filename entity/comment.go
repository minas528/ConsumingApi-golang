package entity

import "time"

// Comment represents comments forwarded by application users
type Comment struct {
	ID uint
	FullName string `gorm:"type:varchar(255)"`
	Message string
	Phone string `gorm:"type:varchar(100);not null; unique"`
	Email string `gorm:"type:varchar(255)'not null; unique'"`
	PlacedAt time.Time
}
