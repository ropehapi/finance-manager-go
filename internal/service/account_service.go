package service

import (
	"context"
	"errors"

	"github.com/ropehapi/finance-manager-go/internal/model"
	"github.com/ropehapi/finance-manager-go/internal/repository"
)

type AccountService interface {
	Create(ctx context.Context, input model.Account) (*model.Account, error)
	GetAll(ctx context.Context) ([]model.Account, error)
	GetByID(ctx context.Context, id string) (*model.Account, error)
	Update(ctx context.Context, id string, input model.Account) (*model.Account, error)
	Delete(ctx context.Context, id string) error
}

type accountService struct {
	repo repository.AccountRepository
}

func NewAccountService(repo repository.AccountRepository) AccountService {
	return &accountService{repo}
}

func (s *accountService) Create(ctx context.Context, input model.Account) (*model.Account, error) {
	if input.Balance < 0 {
		return nil, errors.New("balance must not be negative")
	}
	if input.CurrencyCode == "" || len(input.CurrencyCode) != 3 {
		return nil, errors.New("invalid currency code")
	}

	err := s.repo.Create(ctx, &input)
	return &input, err
}

func (s *accountService) GetAll(ctx context.Context) ([]model.Account, error) {
	return s.repo.FindAll(ctx)
}

func (s *accountService) GetByID(ctx context.Context, id string) (*model.Account, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *accountService) Update(ctx context.Context, id string, input model.Account) (*model.Account, error) {
	account, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	account.Name = input.Name
	account.Kind = input.Kind
	account.CurrencyCode = input.CurrencyCode
	account.Balance = input.Balance

	err = s.repo.Update(ctx, account)
	return account, err
}

func (s *accountService) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
