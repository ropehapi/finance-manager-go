package repository

import (
	"context"

	"github.com/ropehapi/finance-manager-go/internal/model"
	"gorm.io/gorm"
)

type TransferRepository interface {
	Create(ctx context.Context, transfer *model.Transfer) error
	FindByID(ctx context.Context, id string) (*model.Transfer, error)
	GetAll(ctx context.Context, filter model.TransferFilter) ([]model.Transfer, error)
	Update(ctx context.Context, transfer *model.Transfer) error
	Delete(ctx context.Context, id string) error
}

type transferRepository struct {
	db *gorm.DB
}

func NewTransferRepository(db *gorm.DB) TransferRepository {
	return &transferRepository{db}
}

func (r *transferRepository) Create(ctx context.Context, transfer *model.Transfer) error {
	return r.db.WithContext(ctx).Create(transfer).Error
}

func (r *transferRepository) FindByID(ctx context.Context, id string) (*model.Transfer, error) {
	var transfer model.Transfer
	if err := r.db.WithContext(ctx).Preload("Account").Preload("PaymentMethod").
		First(&transfer, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &transfer, nil
}

func (r *transferRepository) GetAll(ctx context.Context, filter model.TransferFilter) ([]model.Transfer, error) {
	var transfers []model.Transfer
	query := r.db.WithContext(ctx)

	if filter.Type != "" {
		query = query.Where("type = ?", filter.Type)
	}

	if filter.Category != "" {
		query = query.Where("category ILIKE ?", "%"+filter.Category+"%")
	}

	err := query.
		Limit(filter.Limit).
		Offset(filter.Offset).
		Preload("Account").
		Preload("PaymentMethod").
		Find(&transfers).Error

	if err != nil {
		return nil, err
	}

	return transfers, nil
}

func (r *transferRepository) Update(ctx context.Context, transfer *model.Transfer) error {
	return r.db.WithContext(ctx).Save(transfer).Error
}

func (r *transferRepository) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Delete(&model.Transfer{}, "id = ?", id).Error
}
