package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/ropehapi/finance-manager-go/internal/model"
	"github.com/ropehapi/finance-manager-go/internal/repository"
)

type TransferService interface {
	Cashin(ctx context.Context, input model.CreateCashinTransferInputDTO) (*model.CreateCashinTransferOutputDTO, error)
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

func (s *transferService) Cashin(ctx context.Context, input model.CreateCashinTransferInputDTO) (*model.CreateCashinTransferOutputDTO, error) {
	if input.Amount <= 0 {
		return nil, errors.New("amount must be positive")
	}
	if input.Currency == "" || len(input.Currency) != 3 {
		return nil, errors.New("invalid currency")
	}

	account, err := s.accountRepo.FindByID(ctx, input.AccountID.String())
	if err != nil {
		return nil, err
	}
	account.Balance += input.Amount

	if err = s.accountRepo.Update(ctx, account); err != nil {
		return nil, err
	}

	parsedDate, err := time.Parse("2006-01-02", input.Date)
	if err != nil {
		return nil, fmt.Errorf("invalid date format: %v", err)
	}

	transfer := model.Transfer{
		Type:         "cashin",
		Currency:     input.Currency,
		Amount:       input.Amount,
		Description:  input.Description,
		Date:         parsedDate,
		CategoryID:   input.CategoryID,
		AccountID:    input.AccountID,
		Observations: input.Observations,
	}

	err = s.repo.Create(ctx, &transfer)
	if err != nil {
		return nil, err
	}

	output := model.CreateCashinTransferOutputDTO{
		ID:           transfer.ID,
		Currency:     transfer.Currency,
		Amount:       transfer.Amount,
		Description:  transfer.Description,
		Date:         transfer.Date.String(),
		CategoryID:   transfer.CategoryID,
		AccountID:    transfer.AccountID,
		Observations: transfer.Observations,
	}
	return &output, err
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
