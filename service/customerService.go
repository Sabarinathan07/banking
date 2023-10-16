package service

import (
	"github.com/sabarinathan07/banking/domain"
	"github.com/sabarinathan07/banking/errs"
)

type CustomerService interface {
	GetAllCustomer() ([]domain.Customer, error)
	GetCustomer(string) (*domain.Customer, *errs.AppError)

}

type defaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s defaultCustomerService) GetAllCustomer() ([]domain.Customer, error) {
	return s.repo.FindAll()
}

func (s defaultCustomerService) GetCustomer(id string) (*domain.Customer, *errs.AppError) {
	return s.repo.ById(id)
}

func NewCustomerService(repository domain.CustomerRepository) defaultCustomerService {
	// injecting repository inside default customerService
	return defaultCustomerService{repository}

}
