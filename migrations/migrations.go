package migrations

import (
	"github.com/ropehapi/finance-manager-go/internal/model"
	"gorm.io/gorm"
	"log"
	"time"
)

func Migrate(db *gorm.DB) {
	if err := db.AutoMigrate(&model.Account{}); err != nil {
		log.Fatalf("migration failed: %v", err)
	}

	if err := db.AutoMigrate(&model.Transfer{}); err != nil {
		log.Fatalf("migration failed: %v", err)
	}

	if err := db.AutoMigrate(&model.PaymentMethod{}); err != nil {
		log.Fatalf("migration failed: %v", err)
	}

	if err := db.AutoMigrate(&model.Debt{}); err != nil {
		log.Fatalf("migration failed: %v", err)
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(time.Hour)

}
