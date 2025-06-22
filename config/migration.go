package config

import (
	"log"

	"github.com/ayusudi/catify-lite/models"
)

func Migrate() {
	err := DB.AutoMigrate(&models.CatFact{}, &models.Comment{})
	if err != nil {
		log.Fatal("❌ Migration failed:", err)
	}
	log.Println("✅ Database migrated successfully")
}
