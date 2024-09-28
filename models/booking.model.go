package models

import "gorm.io/gorm"

type Train struct {
	gorm.Model
	TrainNumber    string `gorm:"unique;not null" json:"train_number" validate:"required"`
	Source         string `gorm:"not null" json:"source" validate:"required"`
	Destination    string `gorm:"not null" json:"destination" validate:"required"`
	TotalSeats     int    `gorm:"not null" json:"total_seats" validate:"required"`
	AvailableSeats int    `gorm:"not null" json:"available_seats" validate:"required"`
}

type Booking struct {
	gorm.Model
	UserID     uint   `gorm:"not null" json:"user_id"`
	TrainID    uint   `gorm:"not null" json:"train_id"`
	Status     string `gorm:"not null" json:"status"`
	SeatNumber int    `gorm:"not null" json:"seat_number"`
}
