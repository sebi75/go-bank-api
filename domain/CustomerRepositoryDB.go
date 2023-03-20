package domain

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type CustomerRepositoryDB struct {
	client *sql.DB
}

func (d CustomerRepositoryDB) FindAll() ([]Customer, error) {
	customers := make([]Customer, 0)

	findAllSql := "select * from customers"
	rows, err := d.client.Query(findAllSql)
	if err != nil {
		log.Println("Error whie querying customer table", err.Error())
		return nil, err
	}

	for rows.Next() {
		var customer Customer
		err := rows.Scan(&customer.Id, &customer.Name, &customer.City, &customer.Zipcode, &customer.DateofBirth, &customer.Status)
		if err != nil {
			log.Println("Error while scanning customer rows into struct", err.Error())
			return nil, err
		}
		customers = append(customers, customer)
	}

	return customers, nil
}

func (d CustomerRepositoryDB) FindById(customerId string) (Customer, error) {
	var customer Customer
	findByIdSql := "select * from customers where customer_id = ?"
	row := d.client.QueryRow(findByIdSql, customerId)
	err := row.Scan(&customer.Id, &customer.Name, &customer.City, &customer.Zipcode, &customer.DateofBirth, &customer.Status)
	if err != nil {
		log.Println("Error while scanning customer row into struct", err.Error())
		return Customer{}, err
	}

	return customer, nil
}

func NewCustomerRepositoryDB() CustomerRepositoryDB {
	client, err := sql.Open("mysql", "root:qWeR1@1`@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err.Error())
	}
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return CustomerRepositoryDB{
		client: client,
	}
}