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

	Transfers []Transfer
	// PaymentMethods []PaymentMethod
	Debts []Debt

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type CreateAccountInputDTO struct {
	CurrencyCode string `json:"currencyCode" binding:"required,min=2,max=3"`
	Balance      int    `json:"balance" binding:"required"`
	Name         string `json:"name" binding:"required,max=36"`
}

type CreateAccountOutputDTO struct {
	Id           uuid.UUID `json:"id"`
	CurrencyCode string    `json:"currencyCode"`
	Balance      int       `json:"balance"`
	Name         string    `json:"name"`
}

func (a *Account) BeforeCreate(tx *gorm.DB) (err error) {
	a.ID = uuid.New()
	return
}
