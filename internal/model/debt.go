package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Debt struct {
	ID              uuid.UUID `gorm:"type:uuid;primaryKey"`
	AccountID       uuid.UUID `gorm:"type:uuid;not null"`
	Account         Account   `gorm:"foreignKey:AccountID"`
	PaymentMethodID uuid.UUID `gorm:"type:uuid;not null"`
	PaymentMethod   PaymentMethod

	Amount    int  `gorm:"not null"`
	Paid      bool `gorm:"not null;default:false"`
	Transfers []Transfer

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (d *Debt) BeforeCreate(tx *gorm.DB) (err error) {
	d.ID = uuid.New()
	return
}
