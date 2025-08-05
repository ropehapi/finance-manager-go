package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PaymentMethod struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name      string    `gorm:"not null"`
	Type      string    `gorm:"not null"` // ex: credit_card, debit_card, pix
	AccountID uuid.UUID `gorm:"type:uuid;not null"`

	Transfers []Transfer
	Debts     []Debt

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type PaymentMethodOutputDTO struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Type      string    `json:"type"`
	AccountID uuid.UUID `json:"accountId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type CreatePaymentMethodInputDTO struct {
	Name      string    `json:"name" binding:"required"`
	Type      string    `json:"type" binding:"required"`
	AccountId uuid.UUID `json:"accountId" binding:"required"`
}

type CreatePaymentMethodOutputDTO struct { //TODO: Avaliar duplicidade
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Type      string    `json:"type"`
	AccountID uuid.UUID `json:"accountId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type UpdatePaymentMethodInputDTO struct {
	Name string `json:"name" binding:"required"`
}

func (p *PaymentMethod) BeforeCreate(tx *gorm.DB) (err error) {
	p.ID = uuid.New()
	return
}
