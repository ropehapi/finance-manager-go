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
	Cashout(ctx context.Context, input model.CreateCashoutTransferInputDTO) (*model.CreateCashoutTransferOutputDTO, error)
	GetAll(ctx context.Context, filter model.TransferFilter) ([]model.TransferOutputDTO, error)
	GetByID(ctx context.Context, id string) (*model.TransferOutputDTO, error)
	Delete(ctx context.Context, id string) error
}

type transferService struct {
	repo              repository.TransferRepository
	accountRepo       repository.AccountRepository
	paymentMethodRepo repository.PaymentMethodRepository
	debtRepo          repository.DebtRepository
}

func NewTransferService(repo repository.TransferRepository, accountRepo repository.AccountRepository, paymentMethodRepo repository.PaymentMethodRepository, debtRepo repository.DebtRepository) TransferService {
	return &transferService{repo, accountRepo, paymentMethodRepo, debtRepo}
}

func (s *transferService) Cashin(ctx context.Context, input model.CreateCashinTransferInputDTO) (*model.CreateCashinTransferOutputDTO, error) {
	if input.Amount <= 0 {
		return nil, errors.New("amount must be positive")
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
		Category:     input.Category,
		AccountID:    input.AccountID,
		Observations: input.Observations,
	}

	err = s.repo.Create(ctx, &transfer)
	if err != nil {
		return nil, err
	}

	output := model.CreateCashinTransferOutputDTO{
		ID:           transfer.ID,
		Type:         transfer.Type,
		Currency:     transfer.Currency,
		Amount:       transfer.Amount,
		Description:  transfer.Description,
		Date:         transfer.Date.String(),
		Category:     transfer.Category,
		AccountID:    transfer.AccountID,
		Observations: transfer.Observations,
	}
	return &output, err
}

func (s *transferService) Cashout(ctx context.Context, input model.CreateCashoutTransferInputDTO) (*model.CreateCashoutTransferOutputDTO, error) {
	if input.Amount <= 0 {
		return nil, errors.New("amount must be positive")
	}

	paymentMethod, err := s.paymentMethodRepo.FindByID(ctx, input.PaymentMethodID.String())
	if err != nil {
		return nil, err
	}

	account, err := s.accountRepo.FindByID(ctx, paymentMethod.AccountID.String())
	if paymentMethod.Type == "debit" {
		if err != nil {
			return nil, err
		}
		account.Balance -= input.Amount
		if account.Balance < 0 {
			return nil, errors.New("insuficient funds")
		}

		if err = s.accountRepo.Update(ctx, account); err != nil {
			return nil, err
		}
	} else if paymentMethod.Type == "credit" {
		debt, err := s.debtRepo.GetUnpaidAccountForPaymentMethod(ctx, input.PaymentMethodID.String())
		if err != nil {
			return nil, err
		}
		if debt == nil {
			debt = &model.Debt{
				Currency:        input.Currency,
				Amount:          input.Amount,
				PaymentMethodID: *input.PaymentMethodID,
			}

			if err = s.debtRepo.Create(ctx, debt); err != nil {
				return nil, err
			}
		} else {
			debt.Amount += input.Amount
			if err = s.debtRepo.Update(ctx, debt); err != nil {
				return nil, err
			}
		}
	}

	parsedDate, err := time.Parse("2006-01-02", input.Date)
	if err != nil {
		return nil, fmt.Errorf("invalid date format: %v", err)
	}

	transfer := model.Transfer{
		Type:            "cashout",
		Currency:        input.Currency,
		Amount:          input.Amount,
		Description:     input.Description,
		Date:            parsedDate,
		Category:        input.Category,
		PaymentMethodID: input.PaymentMethodID,
		AccountID:       &paymentMethod.AccountID,
		Observations:    input.Observations,
	}

	err = s.repo.Create(ctx, &transfer)
	if err != nil {
		return nil, err
	}

	output := model.CreateCashoutTransferOutputDTO{
		ID:              transfer.ID,
		Type:            transfer.Type,
		Currency:        transfer.Currency,
		Amount:          transfer.Amount,
		Description:     transfer.Description,
		Date:            transfer.Date.String(),
		Category:        transfer.Category,
		PaymentMethodID: transfer.PaymentMethodID,
		AccountID:       &paymentMethod.AccountID, //TODO: Reavaliar posteriormente
		Observations:    transfer.Observations,
	}
	return &output, err
}

func (s *transferService) GetAll(ctx context.Context, filter model.TransferFilter) ([]model.TransferOutputDTO, error) {
	transactions, err := s.repo.GetAll(ctx, filter)
	if err != nil {
		return nil, err
	}

	output := make([]model.TransferOutputDTO, len(transactions))
	for i, transaction := range transactions {
		output[i] = model.TransferOutputDTO{
			ID:           transaction.ID,
			Type:         transaction.Type,
			Currency:     transaction.Currency,
			Amount:       transaction.Amount,
			Description:  transaction.Description,
			Date:         transaction.Date.String(),
			Category:     transaction.Category,
			AccountID:    transaction.AccountID,
			Observations: transaction.Observations,
			CreatedAt:    transaction.CreatedAt,
			UpdatedAt:    transaction.UpdatedAt,
		}
	}
	return output, nil
}

func (s *transferService) GetByID(ctx context.Context, id string) (*model.TransferOutputDTO, error) {
	transfer, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &model.TransferOutputDTO{
		ID:              transfer.ID,
		Type:            transfer.Type,
		Currency:        transfer.Currency,
		Amount:          transfer.Amount,
		Description:     transfer.Description,
		Date:            transfer.Date.String(),
		Category:        transfer.Category,
		PaymentMethodID: transfer.PaymentMethodID,
		AccountID:       transfer.AccountID,
		Observations:    transfer.Observations,
		CreatedAt:       transfer.CreatedAt,
		UpdatedAt:       transfer.UpdatedAt,
	}, nil
}

func (s *transferService) Delete(ctx context.Context, id string) error {
	transfer, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	account, err := s.accountRepo.FindByID(ctx, transfer.AccountID.String())
	if err != nil {
		return err
	}

	if transfer.Type == "cashin" {
		account.Balance -= transfer.Amount
	} else if transfer.Type == "cashout" {
		paymentMethod, err := s.paymentMethodRepo.FindByID(ctx, transfer.PaymentMethodID.String())
		if err != nil {
			return err
		}

		if paymentMethod.Type == "debit" {
			account.Balance += transfer.Amount
		} else if paymentMethod.Type == "credit" {
			debt, err := s.debtRepo.GetUnpaidAccountForPaymentMethod(ctx, transfer.PaymentMethodID.String())
			if err != nil {
				return err
			}
			if debt == nil {
				return errors.New("Debt is already payed")
			}
			debt.Amount -= transfer.Amount
			if err = s.debtRepo.Update(ctx, debt); err != nil {
				return err
			}
		}
	}

	if err = s.accountRepo.Update(ctx, account); err != nil {
		return err
	}
	return s.repo.Delete(ctx, id)
}
