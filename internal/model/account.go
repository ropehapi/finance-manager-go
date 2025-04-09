package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Account struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey"`
	Kind         string    `gorm:"not null"`
	CurrencyCode string    `gorm:"size:3;not null"` // ex: USD, BRL
	Name         string    `gorm:"not null"`
	Balance      int       `gorm:"not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

// Hook para gerar UUID automaticamente
func (a *Account) BeforeCreate(tx *gorm.DB) (err error) {
	a.ID = uuid.New()
	return
}
