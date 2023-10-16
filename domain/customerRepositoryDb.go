package domain

import (
	"database/sql"
	"log" // You need to import the "log" package for logging.
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sabarinathan07/banking/errs"
)

type CustomerRepositoryDb struct {
	client *sql.DB
}

func (d CustomerRepositoryDb) FindAll() ([]Customer, error) {
	findAllSql := "SELECT customer_id, name, city, zipcode, date_of_birth, status FROM customers"
	rows, err := d.client.Query(findAllSql)
	if err != nil {
		// Check for an error in Query
		log.Println("Error while querying customer table " + err.Error())
		return nil, err
	}
	defer rows.Close() // Close the rows when done.

	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.ZipCode, &c.DateOfBirth, &c.Status) // Correct variable names.
		if err != nil {
			log.Println("Error while scanning customers " + err.Error())
			return nil, err
		}
		customers = append(customers, c)
	}
	return customers, nil
}
func (d CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
	customerSql := "SELECT customer_id, name, city, zipcode, date_of_birth, status FROM customers WHERE customer_id = ?"
	row := d.client.QueryRow(customerSql, id)
	var c Customer
	err := row.Scan(&c.Id, &c.Name, &c.City, &c.ZipCode, &c.DateOfBirth, &c.Status) // Correct variable names.
	if err != nil {
		if err == sql.ErrNoRows{
			return nil, errs.NewNotFoundError("customer not found")
		}else{
			log.Println("Error while scanning customer" + err.Error())
			return nil, errs.NewUnexpectedError("unexpected database error")
		}

		
	}

	return &c, nil
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	client, err := sql.Open("mysql", "root:@tcp(localhost)/banking")
	if err != nil {
		log.Fatal(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return CustomerRepositoryDb{client}
}
