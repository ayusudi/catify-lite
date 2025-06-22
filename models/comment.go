package models

type Comment struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	FactID    uint      `json:"fact_id"`
	Name      string    `json:"name"`
	Comment   string    `json:"comment"`
}
