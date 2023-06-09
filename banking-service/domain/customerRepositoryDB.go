package domain

import (
	"go-bank-api/errs"
	"go-bank-api/logger"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type CustomerRepositoryDB struct {
	client *sqlx.DB
}

func (d CustomerRepositoryDB) FindAll(status string) ([]Customer, *errs.AppError) {
	customers := make([]Customer, 0)
	var err error

	findAllSql := "select * from customers"
	if status == "" {
		err = d.client.Select(&customers, findAllSql)
	} else {
		err = d.client.Select(&customers, findAllSql, status)
	}

	if err != nil {
		logger.Error("Error while querying customer table" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	// sqlx removes this kind of boilerplate code
	// for rows.Next() {
	// 	var customer Customer
	// 	err := rows.Scan(&customer.Id, &customer.Name, &customer.City, &customer.Zipcode, &customer.DateofBirth, &customer.Status)
	// 	if err != nil {
	// 		logger.Error("Error while scanning customer rows" + err.Error())
	// 		return nil, errs.NewUnexpectedError("Unexpected database error")
	// 	}
	// 	customers = append(customers, customer)
	// }

	return customers, nil
}

func (cr CustomerRepositoryDB) CreateCustomer(customer Customer) (*Customer, *errs.AppError) {
	insertSQL := "insert into customers (name, city, zipcode, date_of_birth) values (?, ?, ?, ?)"
	createdCustomer, err := cr.client.Exec(insertSQL, customer.Name, customer.City, customer.Zipcode, customer.DateofBirth)
	if err != nil {
		logger.Error("Error while creating customer: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	//now get the newly created customer id
	id, err := createdCustomer.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last insert id for new customer: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	//convert the last inserted id to a string with Ioc
	customer.Id = strconv.FormatInt(id, 10)
	return &customer, nil
}

func (cr CustomerRepositoryDB) FindById(id string) (*Customer, *errs.AppError) {
	var c Customer

	customerSQL := "select id, name, date_of_birth, city, zipcode, status from customers where id = ?"
	err := cr.client.Get(&c, customerSQL, id)

	if err != nil {
		logger.Error("Error while getting the customer" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	// err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)
	// if err != nil {
	// 	if err == sql.ErrNoRows {
	// 		return nil, errs.NewNotFoundError("Customer not found")
	// 	} else {
	// 		return nil, errs.NewUnexpectedError("Unexpected database error")
	// 	}
	// }

	return &c, nil
}

func NewCustomerRepositoryDB(client *sqlx.DB) CustomerRepositoryDB {
	return CustomerRepositoryDB{
		client: client,
	}
}
