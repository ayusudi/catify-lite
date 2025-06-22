package models

import "time"

type CatFact struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	ApiID     string    `json:"api_id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}
