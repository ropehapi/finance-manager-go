package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Account struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey"`
	CurrencyCode string    `gorm:"size:3;not null;default:'BRL'"` // ex: USD, BRL
	Balance      int       `gorm:"not null"`
	Name         string    `gorm:"not null"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type AccountFilter struct {
	Name         string
	CurrencyCode string
	Limit        int
	Offset       int
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

func (p *Account) BeforeCreate(tx *gorm.DB) (err error) {
	p.ID = uuid.New()
	return
}
