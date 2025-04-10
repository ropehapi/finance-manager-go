package service

import (
	"context"
	"errors"
	"strings"

	"github.com/ropehapi/finance-manager-go/internal/model"
	"github.com/ropehapi/finance-manager-go/internal/repository"
)

type TransferService interface {
	Create(ctx context.Context, input model.Transfer) (*model.Transfer, error)
	GetAll(ctx context.Context) ([]model.Transfer, error)
	GetByID(ctx context.Context, id string) (*model.Transfer, error)
	Update(ctx context.Context, id string, input model.Transfer) (*model.Transfer, error)
	Delete(ctx context.Context, id string) error
}

type transferService struct {
	repo        repository.TransferRepository
	accountRepo repository.AccountRepository
}

func NewTransferService(repo repository.TransferRepository, accountRepo repository.AccountRepository) TransferService {
	return &transferService{repo, accountRepo}
}

func (s *transferService) Create(ctx context.Context, input model.Transfer) (*model.Transfer, error) {
	if input.Amount <= 0 {
		return nil, errors.New("amount must be positive")
	}
	if input.Currency == "" || len(input.Currency) != 3 {
		return nil, errors.New("invalid currency")
	}

	// Ajusta saldo apenas em cashout com payment method válido
	if strings.ToLower(input.Type) == "cashout" && input.PaymentMethod != nil {
		pm := strings.ToLower(input.PaymentMethod.Type)
		if pm == "debit" || pm == "pix" {
			account, err := s.accountRepo.FindByID(ctx, input.AccountID.String())
			if err != nil {
				return nil, err
			}
			account.Balance -= int(input.Amount)
			if err := s.accountRepo.Update(ctx, account); err != nil {
				return nil, err
			}
		}
	}

	err := s.repo.Create(ctx, &input)
	return &input, err
}

func (s *transferService) GetAll(ctx context.Context) ([]model.Transfer, error) {
	return s.repo.FindAll(ctx)
}

func (s *transferService) GetByID(ctx context.Context, id string) (*model.Transfer, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *transferService) Update(ctx context.Context, id string, input model.Transfer) (*model.Transfer, error) {
	transfer, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	transfer.Type = input.Type
	transfer.Currency = input.Currency
	transfer.Amount = input.Amount
	transfer.Date = input.Date
	transfer.Description = input.Description
	transfer.Observations = input.Observations
	transfer.AccountID = input.AccountID
	transfer.PaymentMethodID = input.PaymentMethodID
	transfer.CategoryID = input.CategoryID

	if err := s.repo.Update(ctx, transfer); err != nil {
		return nil, err
	}
	return transfer, nil
}

func (s *transferService) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
