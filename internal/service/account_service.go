package service

import (
	"context"
	"github.com/ropehapi/finance-manager-go/internal/model"
	"github.com/ropehapi/finance-manager-go/internal/repository"
)

type AccountService interface {
	Create(ctx context.Context, input model.CreateAccountInputDTO) (*model.CreateAccountOutputDTO, error)
	GetAll(ctx context.Context, filter model.AccountFilter) ([]model.CreateAccountOutputDTO, error)
	GetByID(ctx context.Context, id string) (*model.CreateAccountOutputDTO, error)
	Update(ctx context.Context, id string, input model.CreateAccountInputDTO) (*model.CreateAccountOutputDTO, error)
	Delete(ctx context.Context, id string) error
}

type accountService struct {
	repo repository.AccountRepository
}

func NewAccountService(repo repository.AccountRepository) AccountService {
	return &accountService{repo}
}

func (s *accountService) Create(ctx context.Context, input model.CreateAccountInputDTO) (*model.CreateAccountOutputDTO, error) {
	account := &model.Account{
		Balance:      input.Balance,
		CurrencyCode: input.CurrencyCode,
		Name:         input.Name,
	}

	err := s.repo.Create(ctx, account)
	if err != nil {
		return nil, err
	}

	output := &model.CreateAccountOutputDTO{
		Id:           account.ID,
		Balance:      account.Balance,
		CurrencyCode: account.CurrencyCode,
		Name:         account.Name,
	}

	return output, nil
}

func (s *accountService) GetAll(ctx context.Context, filter model.AccountFilter) ([]model.CreateAccountOutputDTO, error) {
	accounts, err := s.repo.GetAll(ctx, filter)
	if err != nil {
		return nil, err
	}

	output := make([]model.CreateAccountOutputDTO, len(accounts))
	for i := range accounts {
		output[i] = model.CreateAccountOutputDTO{
			Id:           accounts[i].ID,
			Balance:      accounts[i].Balance,
			CurrencyCode: accounts[i].CurrencyCode,
			Name:         accounts[i].Name,
		}
	}

	return output, nil
}

func (s *accountService) GetByID(ctx context.Context, id string) (*model.CreateAccountOutputDTO, error) {
	account, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	output := &model.CreateAccountOutputDTO{
		Id:           account.ID,
		Balance:      account.Balance,
		CurrencyCode: account.CurrencyCode,
		Name:         account.Name,
	}
	return output, nil
}

func (s *accountService) Update(ctx context.Context, id string, input model.CreateAccountInputDTO) (*model.CreateAccountOutputDTO, error) {
	account, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	account.Name = input.Name
	account.CurrencyCode = input.CurrencyCode
	account.Balance = input.Balance

	err = s.repo.Update(ctx, account)
	if err != nil {
		return nil, err
	}

	output := &model.CreateAccountOutputDTO{
		Id:           account.ID,
		Balance:      account.Balance,
		CurrencyCode: account.CurrencyCode,
		Name:         account.Name,
	}

	return output, err
}

func (s *accountService) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
