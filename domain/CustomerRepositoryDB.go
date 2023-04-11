package domain

import (
	"database/sql"
	"go-bank-api/errs"
	"log"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type CustomerRepositoryDB struct {
	client *sql.DB
}

func (d CustomerRepositoryDB) FindAll() ([]Customer, *errs.AppError) {
	customers := make([]Customer, 0)

	findAllSql := "select * from customers"
	rows, err := d.client.Query(findAllSql)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("No customers found")
		} else {
			log.Println("Error while querying customer table", err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}

	for rows.Next() {
		var customer Customer
		err := rows.Scan(&customer.Id, &customer.Name, &customer.City, &customer.Zipcode, &customer.DateofBirth, &customer.Status)
		if err != nil {
			log.Println("Error while scanning customer rows", err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
		customers = append(customers, customer)
	}

	return customers, nil
}

func (cr CustomerRepositoryDB) CreateCustomer(customer Customer) (*Customer, *errs.AppError) {
	insertSQL := "insert into customers (name, city, zipcode, date_of_birth) values (?, ?, ?, ?)"
	createdCustomer, err := cr.client.Exec(insertSQL, customer.Name, customer.City, customer.Zipcode, customer.DateofBirth)
	if err != nil {
		log.Println("Error while creating customer: ", err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	//now get the newly created customer id
	id, err := createdCustomer.LastInsertId()
	if err != nil {
		log.Println("Error while getting last insert id for new customer: ", err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	//convert the last inserted id to a string with Ioc
	customer.Id = strconv.FormatInt(id, 10)
	return &customer, nil
}

func (cr CustomerRepositoryDB) FindById(id string) (*Customer, *errs.AppError) {
	var c Customer

	customerSQL := "select id, name, date_of_birth, city, zipcode, status from customers where id = ?"
	row := cr.client.QueryRow(customerSQL, id)

	err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		} else {
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}

	return &c, nil
}

func NewCustomerRepositoryDB() CustomerRepositoryDB {
	client, err := sql.Open("mysql", "root:qwerty123432@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err.Error())
	}
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	log.Println("Database connection successful")

	return CustomerRepositoryDB{
		client: client,
	}
}
