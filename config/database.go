package config

import (
	"crypto/tls"
	"fmt"
	"log"
	"os"

	sqlDriver "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	// Register TLS for TiDB Cloud
	err := sqlDriver.RegisterTLSConfig(os.Getenv("DB_TLS_NAME"), &tls.Config{
		MinVersion: tls.VersionTLS12,
		ServerName: os.Getenv("DB_HOST"),
	})
	if err != nil {
		log.Fatal("❌ TLS config error:", err)
	}

	// Build DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s:4000)/%s?tls=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_TLS_NAME"),
	)

	// Connect to DB
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Failed to connect to TiDB:", err)
	}
	DB = db
	log.Println("✅ Connected to TiDB")
}
