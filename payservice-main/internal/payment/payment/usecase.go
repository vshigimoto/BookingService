package payment

import (
	"payservice/internal/payment/entity"
	"sync"
)

type UseCase interface {
	VerifyToken(accessToken string) (int, error)
	UserPayments(id int, params entity.QueryParams) ([]entity.CardTransaction, error)
	RemoveCardPayment(id int) error
	PaymentByID(userId int, id int) (entity.CardTransaction, error)
	ValidateCard(id int, card entity.Card) error
	UserCards(id int) ([]entity.Card, error)
	ValidateTransaction(transaction entity.CardTransaction, id int) error
	ValidateAccountTransaction(id int, transaction entity.AccountTransaction, wg *sync.WaitGroup) error
	UpdateCard(id int, card entity.Card) error
	DeleteCard(id int) error
	CreateAccount(id int, account entity.Account) error
	GetAccounts(id int) ([]entity.Account, error)
	GetAccount(userid int, id int) (entity.Account, error)
	DeleteAccount(id int) error
	GetAccountPayments(id int) ([]entity.AccountTransaction, error)
	GetAccountPayment(userid int, id int) (entity.AccountTransaction, error)
}
