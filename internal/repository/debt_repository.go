package repository

import (
	"context"
	"errors"

	"github.com/ropehapi/finance-manager-go/internal/model"
	"gorm.io/gorm"
)

type DebtRepository interface {
	Create(ctx context.Context, debt *model.Debt) error
	GetAll(ctx context.Context) ([]model.Debt, error)
	FindByID(ctx context.Context, id string) (*model.Debt, error)
	Update(ctx context.Context, debt *model.Debt) error
	Delete(ctx context.Context, id string) error
	GetUnpaidAccountForPaymentMethod(ctx context.Context, paymentMethodID string) (*model.Debt, error)
}

type debtRepository struct {
	db *gorm.DB
}

func NewDebtRepository(db *gorm.DB) DebtRepository {
	return &debtRepository{db}
}

func (r *debtRepository) Create(ctx context.Context, debt *model.Debt) error {
	return r.db.WithContext(ctx).Create(debt).Error
}

func (r *debtRepository) GetAll(ctx context.Context) ([]model.Debt, error) {
	var debts []model.Debt
	query := r.db.WithContext(ctx)

	if err := query.Find(&debts).Error; err != nil {
		return nil, err
	}

	return debts, nil
}

func (r *debtRepository) FindByID(ctx context.Context, id string) (*model.Debt, error) {
	var debt model.Debt
	if err := r.db.WithContext(ctx).First(&debt, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &debt, nil
}

func (r *debtRepository) Update(ctx context.Context, debt *model.Debt) error {
	return r.db.WithContext(ctx).Save(debt).Error
}

func (r *debtRepository) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Delete(&model.Debt{}, "id = ?", id).Error
}

func (r *debtRepository) GetUnpaidAccountForPaymentMethod(ctx context.Context, paymentMethodId string) (*model.Debt, error) {
	var debt model.Debt
	if err := r.db.WithContext(ctx).First(&debt, "payment_method_id = ? AND paid = false", paymentMethodId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &debt, nil
}
