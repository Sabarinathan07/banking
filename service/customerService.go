package service

import (
	"github.com/sabarinathan07/banking/domain"
)

type CustomerService interface {
	GetAllCustomer() ([]domain.Customer, error)
}

type defaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s defaultCustomerService) GetAllCustomer() ([]domain.Customer, error) {
	return s.repo.FindAll()
}

func NewCustomerService(repository domain.CustomerRepository) defaultCustomerService {
	// injecting repository inside default customerService
	return defaultCustomerService{repository}

}
