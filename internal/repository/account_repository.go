package repository

import (
	"context"

	"github.com/ropehapi/finance-manager-go/internal/model"
	"gorm.io/gorm"
)

type AccountRepository interface {
	Create(ctx context.Context, account *model.Account) error
	FindByID(ctx context.Context, id string) (*model.Account, error)
	GetAll(ctx context.Context, filter model.AccountFilter) ([]model.Account, error)
	Update(ctx context.Context, account *model.Account) error
	Delete(ctx context.Context, id string) error
}

type accountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) AccountRepository {
	return &accountRepository{db}
}

func (r *accountRepository) Create(ctx context.Context, account *model.Account) error {
	return r.db.WithContext(ctx).Create(account).Error
}

func (r *accountRepository) FindByID(ctx context.Context, id string) (*model.Account, error) {
	var account model.Account
	if err := r.db.WithContext(ctx).First(&account, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &account, nil
}

func (r *accountRepository) GetAll(ctx context.Context, filter model.AccountFilter) ([]model.Account, error) {
	var accounts []model.Account
	query := r.db.WithContext(ctx)

	if filter.Name != "" {
		query = query.Where("name ILIKE ?", "%"+filter.Name+"%")
	}

	if filter.CurrencyCode != "" {
		query = query.Where("currency_code = ?", filter.CurrencyCode)
	}

	err := query.
		Limit(filter.Limit).
		Offset(filter.Offset).
		Find(&accounts).Error

	return accounts, err
}

func (r *accountRepository) Update(ctx context.Context, account *model.Account) error {
	return r.db.WithContext(ctx).Save(account).Error
}

func (r *accountRepository) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Delete(&model.Account{}, "id = ?", id).Error
}
