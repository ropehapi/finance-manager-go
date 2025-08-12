package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Category struct {
	ID   uuid.UUID `gorm:"type:uuid;primaryKey"`
	Type string    `gorm:"type:varchar(20) CHECK(type IN ('earning', 'expense', 'investment'));not null"` // earning, expense, investment
	Name string    `gorm:"not null"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
