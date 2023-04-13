package domain

import (
	errs "banking-auth/error"
	"strconv"

	"github.com/jmoiron/sqlx"
)

type DefaultAuthRepository struct {
	client *sqlx.DB
}

func (r DefaultAuthRepository) FindById(userId string) (*User, *errs.AppError) {
	var user User
	err := r.client.Get(&user, "select * from users where user_id = ?", userId)
	if err != nil {
		return nil, errs.NewNotFoundError("user not found")
	}
	return &user, nil
}

func (r DefaultAuthRepository) CreateUser(user User) (*User, *errs.AppError) {
	insertUserSQL := "insert into users (username, password, customer_id, role) values (?, ?, ?, ?)"

	result, err := r.client.Exec(insertUserSQL, user.Username, user.Password, user.CustomerId, user.Role)
	if err != nil {
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	userId, err := result.LastInsertId()
	if err != nil {
		return nil, errs.NewUnexpectedError("unexpected database error")
	}
	stringifiedUserId := strconv.FormatInt(userId, 10)
	user.Id = stringifiedUserId

	return &user, nil
}

func NewAuthRepostoryDB(client *sqlx.DB) DefaultAuthRepository {
	return DefaultAuthRepository{client: client}
}