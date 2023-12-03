package repository

import (
	"database/sql"

	"payservice/internal/payment/entity"
)

type Repository interface {
	PaymentRepository
}

type PaymentRepository interface {
	GetPayments(id int) ([]entity.CardTransaction, error)
	DeleteCardPayment(id int) error
	GetPaymentByID(userId int, id int) (entity.CardTransaction, error)
	CreateCard(id int, card entity.Card) error
	GetCards(id int) ([]entity.Card, error)
	CreatePayment(transaction entity.CardTransaction) error
	UpdateCard(id int, card entity.Card) error
	DeleteCard(id int) error
	CreateAccount(account entity.Account) error
	GetAccount(id int) ([]entity.Account, error)
	CreateAccountPayment(transaction entity.AccountTransaction) error
	GetAccountPayments(id int) ([]entity.AccountTransaction, error)
	GetAccountPayment(userid int, id int) (entity.AccountTransaction, error)
}

type Repo struct {
	main         *sql.DB
	replica      *sql.DB
	queryBuilder SQLQueryBuilder
}

func NewRepository(main *sql.DB, replica *sql.DB, queryBuilder SQLQueryBuilder) *Repo {
	return &Repo{
		main:         main,
		replica:      replica,
		queryBuilder: queryBuilder,
	}
}
