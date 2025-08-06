package service

import (
	"context"
	"errors"
	"strings"

	"github.com/ropehapi/finance-manager-go/internal/model"
	"github.com/ropehapi/finance-manager-go/internal/repository"
)

type PaymentMethodService interface {
	Create(ctx context.Context, input model.CreatePaymentMethodInputDTO) (*model.CreatePaymentMethodOutputDTO, error)
	GetAll(ctx context.Context, filter model.PaymentMethodFilter) ([]model.PaymentMethodOutputDTO, error)
	GetByID(ctx context.Context, id string) (*model.PaymentMethodOutputDTO, error)
	Update(ctx context.Context, id string, input model.UpdatePaymentMethodInputDTO) (*model.PaymentMethodOutputDTO, error)
	Delete(ctx context.Context, id string) error
}

type paymentMethodService struct {
	repo repository.PaymentMethodRepository
}

func NewPaymentMethodService(repo repository.PaymentMethodRepository) PaymentMethodService {
	return &paymentMethodService{repo}
}

func (s *paymentMethodService) Create(ctx context.Context, input model.CreatePaymentMethodInputDTO) (*model.CreatePaymentMethodOutputDTO, error) {
	if input.Type != "credit" && input.Type != "debit" {
		return nil, errors.New("invalid type")
	} //TODO: Verificar se é necessário validar aqui ou colocar na struct

	paymentMethod := &model.PaymentMethod{
		Name:      input.Name,
		Type:      strings.ToLower(input.Type),
		AccountID: input.AccountId,
	}

	err := s.repo.Create(ctx, paymentMethod)
	if err != nil {
		return nil, err
	}

	output := &model.CreatePaymentMethodOutputDTO{
		ID:        paymentMethod.ID,
		Name:      paymentMethod.Name,
		Type:      paymentMethod.Type,
		CreatedAt: paymentMethod.CreatedAt,
		UpdatedAt: paymentMethod.UpdatedAt,
	}

	return output, err
}

func (s *paymentMethodService) GetAll(ctx context.Context, filter model.PaymentMethodFilter) ([]model.PaymentMethodOutputDTO, error) {
	paymentMethods, err := s.repo.GetAll(ctx, filter)
	if err != nil {
		return nil, err
	}

	output := make([]model.PaymentMethodOutputDTO, len(paymentMethods))
	for i, paymentMethod := range paymentMethods {
		output[i] = model.PaymentMethodOutputDTO{
			ID:        paymentMethod.ID,
			Name:      paymentMethod.Name,
			Type:      paymentMethod.Type,
			AccountID: paymentMethod.AccountID,
			CreatedAt: paymentMethod.CreatedAt,
			UpdatedAt: paymentMethod.UpdatedAt,
		}
	}

	return output, nil
}

func (s *paymentMethodService) GetByID(ctx context.Context, id string) (*model.PaymentMethodOutputDTO, error) {
	paymentMethod, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	output := &model.PaymentMethodOutputDTO{
		ID:        paymentMethod.ID,
		Name:      paymentMethod.Name,
		Type:      paymentMethod.Type,
		AccountID: paymentMethod.AccountID,
		CreatedAt: paymentMethod.CreatedAt,
		UpdatedAt: paymentMethod.UpdatedAt,
	}

	return output, nil
}

func (s *paymentMethodService) Update(ctx context.Context, id string, input model.UpdatePaymentMethodInputDTO) (*model.PaymentMethodOutputDTO, error) {
	method, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	method.Name = input.Name

	err = s.repo.Update(ctx, method)
	if err != nil {
		return nil, err
	}

	output := &model.PaymentMethodOutputDTO{
		ID:        method.ID,
		Name:      method.Name,
		Type:      method.Type,
		AccountID: method.AccountID,
		CreatedAt: method.CreatedAt,
		UpdatedAt: method.UpdatedAt,
	}

	return output, err
}

func (s *paymentMethodService) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
