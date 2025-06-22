package config

import (
	"catify-lite/models"
	"log"
)

func Migrate() {
	err := DB.AutoMigrate(&models.CatFact{}, &models.Comment{})
	if err != nil {
		log.Fatal("❌ Migration failed:", err)
	}
	log.Println("✅ Database migrated successfully")
}
