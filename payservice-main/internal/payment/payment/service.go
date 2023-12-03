package payment

import (
	"fmt"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"

	"payservice/internal/payment/entity"
	"payservice/internal/payment/repository"
)

type Service struct {
	repo              repository.Repo
	jwtSecretKey      string
	passwordSecretKey string
	logger            *zap.SugaredLogger
	mu                sync.Mutex
}

func NewService(repo repository.Repo, jwtSecretKey string, passwordSecretKey string, l *zap.SugaredLogger) *Service {
	return &Service{
		repo:              repo,
		jwtSecretKey:      jwtSecretKey,
		passwordSecretKey: passwordSecretKey,
		logger:            l,
	}
}

func (s *Service) VerifyToken(accessToken string) (int, error) {
	type Claims struct {
		UserID int
		jwt.RegisteredClaims
	}

	claims := &Claims{}
	_, err := jwt.ParseWithClaims(accessToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.jwtSecretKey), nil
	})

	if err != nil {
		s.logger.Error("Error with token verifying")
		return 0, err
	}

	if time.Now().After(claims.ExpiresAt.Time) {
		s.logger.Error("AccessToken is expired")
		return 0, err
	}

	s.logger.Infof("Usertoken is Verified. User id is %d", claims.UserID)
	return claims.UserID, nil
}

func (s *Service) UserPayments(id int, params entity.QueryParams) ([]entity.CardTransaction, error) {
	//fmt.Println(id)

	data, err := s.repo.GetPayments(id, params)
	if err != nil {
		s.logger.Errorf("Error with GetPayments. UserId is %d", id)
		return []entity.CardTransaction{}, err
	}

	return data, nil
}

func (s *Service) PaymentByID(userId int, id int) (entity.CardTransaction, error) {
	data, err := s.repo.GetPaymentByID(userId, id)
	if err != nil {
		s.logger.Errorf("Error with GetPayments. UserId is %d", id)
		return entity.CardTransaction{}, err
	}

	return data, nil
}

func (s *Service) ValidateCard(id int, card entity.Card) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	fmt.Println("WORKING")

	err := s.repo.CreateCard(id, card)
	if err != nil {
		s.logger.Errorf("Error with create card. User id is %d", id)
		return err
	}

	return nil
}

func (s *Service) UserCards(id int) ([]entity.Card, error) {
	//fmt.Println(id)

	cards, err := s.repo.GetCards(id)
	if err != nil {
		s.logger.Errorf("Error with getting cards. User id is %d", id)
		return []entity.Card{}, err
	}

	return cards, nil
}

func (s *Service) ValidateTransaction(transaction entity.CardTransaction, id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	cards, err := s.repo.GetCards(id)
	if err != nil {
		s.logger.Errorf("Error with getting card. User id is %d", id)
		return err
	}

	cardIsExist := false
	for i := 0; i < len(cards); i++ {
		if cards[i].ID == transaction.FromCardID {
			cardIsExist = true
		}
	}

	if !cardIsExist {
		s.logger.Error("Card not exist")
		return fmt.Errorf("Card not exist")
	}

	err = s.repo.CreatePayment(transaction)
	if err != nil {
		s.logger.Errorf("Error with create payment. User id is %d", id)
		return err
	}

	s.logger.Info("Transaction is GOOD")
	return nil
}

func (s *Service) UpdateCard(id int, card entity.Card) error {
	err := s.repo.UpdateCard(id, card)
	if err != nil {
		s.logger.Errorf("Error with updating card. User id is %d", id)
		return err
	}

	return nil
}

func (s *Service) DeleteCard(id int) error {
	err := s.repo.DeleteCard(id)
	if err != nil {
		s.logger.Errorf("Error with delete card. User id is %d", id)
		return err
	}

	return nil
}

func (s *Service) CreateAccount(id int, account entity.Account) error {
	account.UserID = id

	err := s.repo.CreateAccount(account)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) GetAccounts(id int) (accounts []entity.Account, err error) {

	accounts, err = s.repo.GetAccounts(id)
	if err != nil {
		return []entity.Account{}, err
	}

	return accounts, nil
}

func (s *Service) GetAccount(userid int, id int) (account entity.Account, err error) {
	account, err = s.repo.GetAccount(userid, id)
	if err != nil {
		return entity.Account{}, nil
	}

	return account, err
}

func (s *Service) DeleteAccount(id int) (err error) {

	err = s.repo.DeleteAccount(id)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) ValidateAccountTransaction(id int, transaction entity.AccountTransaction, wg *sync.WaitGroup) error {
	defer wg.Done()

	s.mu.Lock()
	defer s.mu.Unlock()

	account, err := s.repo.GetAccounts(transaction.ToID)

	if err != nil {
		s.logger.Error("Account not exist")
		return err
	}

	ok := false

	for i := 0; i < len(account); i++ {
		if account[i].Currency == transaction.Currency {
			ok = true
		}
	}

	if !ok {
		s.logger.Error("Account Currency not exist")
		return fmt.Errorf("Account Currency not exist")
	}

	err = s.repo.CreateAccountPayment(transaction)
	if err != nil {
		s.logger.Error("Create Payment is failed")
		return err
	}

	s.logger.Info("Transaction is GOOD")
	return nil
}

func (s *Service) GetAccountPayments(id int) (transactions []entity.AccountTransaction, err error) {
	transactions, err = s.repo.GetAccountPayments(id)
	if err != nil {
		return []entity.AccountTransaction{}, err
	}

	return transactions, nil
}

func (s *Service) GetAccountPayment(userid int, id int) (transactions entity.AccountTransaction, err error) {
	transactions, err = s.repo.GetAccountPayment(userid, id)
	if err != nil {
		return entity.AccountTransaction{}, err
	}

	return transactions, nil
}

func (s *Service) RemoveCardPayment(id int) error {
	err := s.repo.DeleteCardPayment(id)
	if err != nil {
		return err
	}

	return nil
}
