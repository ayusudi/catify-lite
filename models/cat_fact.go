package models

type CatFact struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Fact      string    `gorm:"type:varchar(255);uniqueIndex" json:"fact"`
}