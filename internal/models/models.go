package models

import (
	"gorm.io/gorm"
	"time"
)

type Song struct {
	ID        uint           `json:"id" example:"1" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at" example:"2025-03-22T10:00:00Z"`
	UpdatedAt time.Time      `json:"updated_at" example:"2025-03-22T10:01:00Z"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	Group       string `json:"group" example:"Imagine Dragons"`
	Song        string `json:"song" example:"Believer"`
	ReleaseDate string `json:"release_date" example:"2017-02-01"`
	Text        string `json:"text" example:"First things first..."`
	Link        string `json:"link" example:"https://youtube.com/believer"`
}

type SongDetail struct {
	ReleaseDate string `json:"release_date" example:"2017-02-01"`
	Text        string `json:"text" example:"First things first..."`
	Link        string `json:"link" example:"https://youtube.com/believer"`
}
