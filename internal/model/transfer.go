package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Transfer struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey"`
	Type         string    `gorm:"type:varchar(20);not null"` // cashin, cashout, debt_payment
	Currency     string    `gorm:"size:3;not null"`           // ex: BRL, USD
	Amount       int64     `gorm:"not null"`
	Date         time.Time `gorm:"not null"`
	Description  string    `gorm:"type:text"`
	Observations string    `gorm:"type:text"`

	AccountID uuid.UUID `gorm:"type:uuid;not null"`
	Account   Account   `gorm:"foreignKey:AccountID"`

	PaymentMethodID *uuid.UUID     `gorm:"type:uuid"`
	PaymentMethod   *PaymentMethod `gorm:"foreignKey:PaymentMethodID"`

	CategoryID *uuid.UUID `gorm:"type:uuid"`
	Category   *Category  `gorm:"foreignKey:CategoryID"`

	DebtID *uuid.UUID `gorm:"type:uuid"`
	Debt   *Debt      `gorm:"foreignKey:DebtID"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (t *Transfer) BeforeCreate(tx *gorm.DB) (err error) {
	t.ID = uuid.New()
	return
}
