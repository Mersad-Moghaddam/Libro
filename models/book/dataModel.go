package book

import "time"

type Book struct {
	ID          uint       `gorm:"primaryKey"`
	UserID      uint       `gorm:"index;not null"`
	Title       string     `gorm:"size:255;not null"`
	Author      string     `gorm:"size:255;not null"`
	TotalPages  int        `gorm:"not null"`
	Status      string     `gorm:"size:40;not null"`
	CurrentPage int        `gorm:"default:0;not null"`
	CompletedAt *time.Time `gorm:"default:null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (Book) TableName() string { return "books" }
