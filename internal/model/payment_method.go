package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PaymentMethod struct {
	ID   uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name string    `gorm:"not null"`
	Type string    `gorm:"not null"` // ex: credit_card, debit_card, pix

	AccountID uuid.UUID `gorm:"type:uuid;not null"`
	Account   Account   `gorm:"foreignKey:AccountID"`

	Transfers []Transfer
	Debts     []Debt

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (p *PaymentMethod) BeforeCreate(tx *gorm.DB) (err error) {
	p.ID = uuid.New()
	return
}
