package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Transfer struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	Type        string    `gorm:"type:varchar(20);not null;check:type IN ('cashin','cashout','debt_payment')"`
	Currency    string    `gorm:"size:3;not null;default:'BRL'"`
	Amount      int       `gorm:"not null"`
	Category    string    `gorm:"type:varchar(20);not null"`
	Description string    `gorm:"type:text"`
	Date        time.Time `gorm:"not null"`

	PaymentMethodID *uuid.UUID     `gorm:"type:uuid"`
	PaymentMethod   *PaymentMethod `gorm:"foreignKey:PaymentMethodID"`

	AccountID *uuid.UUID `gorm:"type:uuid;not null"`
	Account   *Account   `gorm:"foreignKey:AccountID"`

	Observations string `gorm:"type:text"`

	CreatedAt time.Time
	UpdatedAt time.Time
	//DeletedAt gorm.DeletedAt `gorm:"index"`
}

type TransferFilter struct {
	Type     string
	Category string
	Limit    int
	Offset   int
}

type TransferOutputDTO struct {
	ID              uuid.UUID  `json:"id"`
	Type            string     `json:"type"`
	Currency        string     `json:"currency"`
	Amount          int        `json:"amount"`
	Description     string     `json:"description"`
	Date            string     `json:"date"`
	Category        string     `json:"category"`
	PaymentMethodID *uuid.UUID `json:"payment_method_id"`
	AccountID       *uuid.UUID `json:"account_id"`
	Observations    string     `json:"observations"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	DeletedAt       time.Time  `json:"deleted_at"`
}

type CreateCashinTransferInputDTO struct {
	Currency     string     `json:"currency"`
	Amount       int        `json:"amount" binding:"required,gt=0"`
	Description  string     `json:"description" binding:"required"`
	Date         string     `json:"date" binding:"required"`
	Category     string     `json:"category" binding:"required"`
	AccountID    *uuid.UUID `json:"account_id" binding:"required"`
	Observations string     `json:"observations"`
}

type CreateCashinTransferOutputDTO struct {
	ID           uuid.UUID  `json:"id"`
	Type         string     `json:"type"`
	Currency     string     `json:"currency"`
	Amount       int        `json:"amount"`
	Description  string     `json:"description"`
	Date         string     `json:"date"`
	Category     string     `json:"category"`
	AccountID    *uuid.UUID `json:"account_id" binding:"required"`
	Observations string     `json:"observations"`
}

type CreateCashoutTransferInputDTO struct {
	Currency        string     `json:"currency"`
	Amount          int        `json:"amount" binding:"required,gt=0"`
	Description     string     `json:"description" binding:"required"`
	Date            string     `json:"date" binding:"required"`
	Category        string     `json:"category" binding:"required"`
	PaymentMethodID *uuid.UUID `json:"payment_method_id" binding:"required"`
	Observations    string     `json:"observations"`
}

type CreateCashoutTransferOutputDTO struct {
	ID              uuid.UUID  `json:"id"`
	Type            string     `json:"type"`
	Currency        string     `json:"currency"`
	Amount          int        `json:"amount"`
	Description     string     `json:"description"`
	Date            string     `json:"date"`
	Category        string     `json:"category"`
	PaymentMethodID *uuid.UUID `json:"payment_method_id"`
	AccountID       *uuid.UUID `json:"account_id"`
	Observations    string     `json:"observations"`
}

func (p *Transfer) BeforeCreate(tx *gorm.DB) (err error) {
	p.ID = uuid.New()
	return
}
