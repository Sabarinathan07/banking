package domain

import "github.com/sabarinathan07/banking/errs"

type Customer struct{
	Id string
	Name string
	City string
	ZipCode string
	DateOfBirth string
	Status string
}

//secondary port
type CustomerRepository interface{
	FindAll() ([]Customer,error)
	ById(string) (*Customer, *errs.AppError)
}
