package domain

import errs "banking-auth/error"

/*
This is a stub implementation of the UserRepository interface.

This can be used for testing purposes, or for a quick and dirty implementation of the UserRepository interface.
*/

type DefaultAuthRepositoryStub struct {
	users map[string]*User
}

func (r DefaultAuthRepositoryStub) FindById(userId string) (*User, *errs.AppError) {
	user := r.users[userId]
	if user == nil {
		return nil, errs.NewNotFoundError("user not found")
	}
	return user, nil
}

func (r DefaultAuthRepositoryStub) CreateUser(user User) (*User, *errs.AppError) {
	existingUser := r.users[user.Username]
	if existingUser != nil {
		return nil, errs.NewBadRequestError("username already exists")
	}
	r.users[user.Username] = &user
	return &user, nil
}


func NewAutRepositoryStub() AuthRepository {
	users := make(map[string]*User)
	users["1"] = &User{Id: "1", Username: "username1", Password: "password1", CustomerId: 1}
	users["2"] = &User{Id: "2", Username: "username2", Password: "password2", CustomerId: 2}
	return &DefaultAuthRepositoryStub{users: users}
}