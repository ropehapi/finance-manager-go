package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Debt struct {
	ID              uuid.UUID  `gorm:"type:uuid;primaryKey"`
	Currency        string     `gorm:"type:text;default:'BRL';not null"`
	Amount          int        `gorm:"not null"`
	PaymentMethodID uuid.UUID  `gorm:"type:uuid;not null"`
	PayerAccountID  *uuid.UUID `gorm:"type:uuid"`
	Paid            bool       `gorm:"not null;default:false"`

	PayerAccount  Account       `gorm:"foreignKey:PayerAccountID"`
	PaymentMethod PaymentMethod `gorm:"foreignKey:PaymentMethodID"`

	CreatedAt time.Time
	UpdatedAt time.Time
	//DeletedAt gorm.DeletedAt `gorm:"index"`
}

type DebtOutputDTO struct {
	ID              uuid.UUID  `json:"id"`
	Currency        string     `json:"currency"`
	Amount          int        `json:"amount"`
	PaymentMethodID uuid.UUID  `json:"payment_method_id"`
	PayerAccountID  *uuid.UUID `json:"payer_account_id"`
	Paid            bool       `json:"paid"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
}

func (p *Debt) BeforeCreate(tx *gorm.DB) (err error) {
	p.ID = uuid.New()
	return
}
