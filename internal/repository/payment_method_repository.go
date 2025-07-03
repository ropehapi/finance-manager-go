package repository

import (
	"context"

	"github.com/ropehapi/finance-manager-go/internal/model"
	"gorm.io/gorm"
)

type PaymentMethodRepository interface {
	Create(ctx context.Context, method *model.PaymentMethod) error
	FindAll(ctx context.Context) ([]model.PaymentMethod, error)
	FindByID(ctx context.Context, id string) (*model.PaymentMethod, error)
	Update(ctx context.Context, method *model.PaymentMethod) error
	Delete(ctx context.Context, id string) error
}

type paymentMethodRepository struct {
	db *gorm.DB
}

func NewPaymentMethodRepository(db *gorm.DB) PaymentMethodRepository {
	return &paymentMethodRepository{db}
}

func (r *paymentMethodRepository) Create(ctx context.Context, method *model.PaymentMethod) error {
	return r.db.WithContext(ctx).Create(method).Error
}

func (r *paymentMethodRepository) FindAll(ctx context.Context) ([]model.PaymentMethod, error) {
	var methods []model.PaymentMethod
	err := r.db.WithContext(ctx).Find(&methods).Error
	return methods, err
}

func (r *paymentMethodRepository) FindByID(ctx context.Context, id string) (*model.PaymentMethod, error) {
	var method model.PaymentMethod
	if err := r.db.WithContext(ctx).First(&method, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &method, nil
}

func (r *paymentMethodRepository) Update(ctx context.Context, method *model.PaymentMethod) error {
	return r.db.WithContext(ctx).Save(method).Error
}

func (r *paymentMethodRepository) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Delete(&model.PaymentMethod{}, "id = ?", id).Error
}
