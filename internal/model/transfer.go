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
	Amount       int       `gorm:"not null"`
	Date         time.Time `gorm:"not null"`
	Description  string    `gorm:"type:text"`
	Observations string    `gorm:"type:text"`

	AccountID *uuid.UUID `gorm:"type:uuid;not null"`
	Account   *Account   `gorm:"foreignKey:AccountID"`

	PaymentMethodID *uuid.UUID     `gorm:"type:uuid"`
	PaymentMethod   *PaymentMethod `gorm:"foreignKey:PaymentMethodID"`

	CategoryID *uuid.UUID `gorm:"type:uuid"`
	Category   *Category  `gorm:"foreignKey:CategoryID"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type CreateCashinTransferInputDTO struct {
	Currency     string     `json:"currency" binding:"required"`
	Amount       int        `json:"amount" binding:"required"`
	Description  string     `json:"description" binding:"required"`
	Date         string     `json:"date" binding:"required"`
	CategoryID   *uuid.UUID `json:"category_id"`
	AccountID    *uuid.UUID `json:"account_id" binding:"required"`
	Observations string     `json:"observations"`
}

type CreateCashinTransferOutputDTO struct {
	ID           uuid.UUID  `json:"id"`
	Currency     string     `json:"currency"`
	Amount       int        `json:"amount"`
	Description  string     `json:"description"`
	Date         string     `json:"date"`
	CategoryID   *uuid.UUID `json:"category_id"`
	AccountID    *uuid.UUID `json:"account_id" binding:"required"`
	Observations string     `json:"observations"`
}

func (t *Transfer) BeforeCreate(tx *gorm.DB) (err error) {
	t.ID = uuid.New()
	return
}
