package models

import (
	"time"

	"gorm.io/gorm"
)

type Tickets struct {
	gorm.Model
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" gorm:"text;null;default:null"`
	Content   *string   `json:"content" gorm:"text;not null;default:null"`
	CreatedAt time.Time `json:"created_at" gorm:"text;not null;default:null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"text;null;default:null"`
	Assignee  *string   `json:"assignee" gorm:"text;null;default:null"`
}
