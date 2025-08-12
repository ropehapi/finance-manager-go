package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/ropehapi/finance-manager-go/internal/model"
	"github.com/ropehapi/finance-manager-go/internal/repository"
	"time"
)

type DebtService interface {
	GetAll(ctx context.Context) ([]model.DebtOutputDTO, error)
	Pay(ctx context.Context, id, payerAccountId string) (*model.Message, error)
	Delete(ctx context.Context, id string) error
}

type debtService struct {
	debtRepo     repository.DebtRepository
	accountRepo  repository.AccountRepository
	transferRepo repository.TransferRepository
}

func NewDebtService(debtRepo repository.DebtRepository, accountRepo repository.AccountRepository, transferRepo repository.TransferRepository) DebtService {
	return &debtService{
		debtRepo:     debtRepo,
		accountRepo:  accountRepo,
		transferRepo: transferRepo,
	}
}

func (s *debtService) GetAll(ctx context.Context) ([]model.DebtOutputDTO, error) {
	debts, err := s.debtRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	output := make([]model.DebtOutputDTO, len(debts))
	for i, debt := range debts {
		output[i] = model.DebtOutputDTO{
			ID:              debt.ID,
			Currency:        debt.Currency,
			Amount:          debt.Amount,
			PaymentMethodID: debt.PaymentMethodID,
			PayerAccountID:  debt.PayerAccountID,
			Paid:            debt.Paid,
			CreatedAt:       debt.CreatedAt,
			UpdatedAt:       debt.UpdatedAt,
		}
	}

	return output, nil
}

func (s *debtService) Pay(ctx context.Context, id, payerAccountId string) (*model.Message, error) {
	debt, err := s.debtRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	debt.Paid = true

	if err = s.debtRepo.Update(ctx, debt); err != nil {
		return nil, err
	}

	account, err := s.accountRepo.FindByID(ctx, payerAccountId)
	if err != nil {
		return nil, err
	}

	account.Balance -= debt.Amount
	if err = s.accountRepo.Update(ctx, account); err != nil {
		return nil, err
	}

	parsedAccountId, err := uuid.Parse(payerAccountId)
	if err != nil {
		return nil, err
	}

	transfer := &model.Transfer{
		Type:      "debt_payment",
		Currency:  debt.Currency,
		Amount:    debt.Amount,
		Category:  "despesas",
		AccountID: &parsedAccountId,
		Date:      time.Now(),
	}
	if err = s.transferRepo.Create(ctx, transfer); err != nil {
		return nil, err
	}

	output := &model.Message{
		Message: "Debt has been paid",
	}

	return output, nil
}

func (s *debtService) Delete(ctx context.Context, id string) error {
	return s.debtRepo.Delete(ctx, id)
}
