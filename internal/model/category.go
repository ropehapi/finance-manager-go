package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Category struct {
	ID   uuid.UUID `gorm:"type:uuid;primaryKey"`
	Type string    `gorm:"type:varchar(20);not null"` // earning, expense, investment
	Name string    `gorm:"not null"`

	Transfers []Transfer

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (c *Category) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.New()
	return
}
