package service

import (
	"context"
	"errors"
	"strings"

	"github.com/ropehapi/finance-manager-go/internal/model"
	"github.com/ropehapi/finance-manager-go/internal/repository"
)

type PaymentMethodService interface {
	Create(ctx context.Context, input model.PaymentMethod) (*model.PaymentMethod, error)
	GetAll(ctx context.Context) ([]model.PaymentMethod, error)
	GetByID(ctx context.Context, id string) (*model.PaymentMethod, error)
	Update(ctx context.Context, id string, input model.PaymentMethod) (*model.PaymentMethod, error)
	Delete(ctx context.Context, id string) error
}

type paymentMethodService struct {
	repo repository.PaymentMethodRepository
}

func NewPaymentMethodService(repo repository.PaymentMethodRepository) PaymentMethodService {
	return &paymentMethodService{repo}
}

func (s *paymentMethodService) Create(ctx context.Context, input model.PaymentMethod) (*model.PaymentMethod, error) {
	if input.Name == "" || input.Type == "" {
		return nil, errors.New("name and type are required")
	}

	input.Type = strings.ToLower(input.Type)

	err := s.repo.Create(ctx, &input)
	return &input, err
}

func (s *paymentMethodService) GetAll(ctx context.Context) ([]model.PaymentMethod, error) {
	return s.repo.FindAll(ctx)
}

func (s *paymentMethodService) GetByID(ctx context.Context, id string) (*model.PaymentMethod, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *paymentMethodService) Update(ctx context.Context, id string, input model.PaymentMethod) (*model.PaymentMethod, error) {
	method, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	method.Name = input.Name
	method.Type = strings.ToLower(input.Type)

	err = s.repo.Update(ctx, method)
	return method, err
}

func (s *paymentMethodService) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
