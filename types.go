package main

import (
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Number		int64	`json:"number"`
	Password	string	`josn:"password"`
}

type LoginResponse struct {
	Number 	int64 	`json:"number"`
	Token	string 	`json:"token"`
}

type CreateAccountRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Password  string `json:"password"`
}

type TransferRequest struct {
	ToAccount	int	`json:"toAccount"`
	Amount		int `json:"amount"`
}

type Account struct {
	ID        			int			`json:"id"`
	FirstName 			string    	`json:"firstName"`
	LastName  			string    	`json:"lastName"`
	EncryptedPassword 	string		`json:"-"`
	Number    			int64		`json:"number"`
	Balance				int64	    `json:"balance"`
	CreatedAt			time.Time	`json:"createdAt"`
}

func NewAccount(firstName, lastName , password string) (*Account, error) {
	encpwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return &Account{
		FirstName: 			firstName,
		LastName:  			lastName,
		EncryptedPassword:	string(encpwd),
		Number:    			int64(rand.Intn(10000000)),
		CreatedAt: 			time.Now().UTC(),
	}, nil
}

func (a *Account) validPassword(pw string) bool {
	return bcrypt.CompareHashAndPassword(
		[]byte(a.EncryptedPassword),
		[]byte(pw),
	) == nil
}