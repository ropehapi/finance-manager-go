package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ropehapi/finance-manager-go/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDatabase() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"), os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	if err := db.AutoMigrate(&model.Account{}); err != nil {
		log.Fatalf("migration failed: %v", err)
	}

	if err := db.AutoMigrate(&model.Transfer{}); err != nil {
		log.Fatalf("migration failed: %v", err)
	}

	if err := db.AutoMigrate(&model.PaymentMethod{}); err != nil {
		log.Fatalf("migration failed: %v", err)
	}

	if err := db.AutoMigrate(&model.Category{}); err != nil {
		log.Fatalf("migration failed: %v", err)
	}

	if err := db.AutoMigrate(&model.Debt{}); err != nil {
		log.Fatalf("migration failed: %v", err)
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db
}
