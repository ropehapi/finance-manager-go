package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Debt struct {
	ID              uuid.UUID `gorm:"type:uuid;primaryKey"`
	Currency        string    `gorm:"type:text;default:'BRL';not null"`
	Amount          int       `gorm:"not null"`
	PaymentMethodID uuid.UUID `gorm:"type:uuid;not null"`
	PayerAccountID  uuid.UUID `gorm:"type:uuid"`
	Paid            bool      `gorm:"not null;default:false"`

	PayerAccount  Account       `gorm:"foreignKey:PayerAccountID"`
	PaymentMethod PaymentMethod `gorm:"foreignKey:PaymentMethodID"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
