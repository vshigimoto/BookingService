package repository

import (
	"database/sql"
	"fmt"
	"time"

	"payservice/internal/payment/entity"
)

func (r *Repo) GetPayments(id int, params entity.QueryParams) (data []entity.CardTransaction, err error) {
	var q string
	var rows *sql.Rows

	subquery := r.queryBuilder.
		Select("card_id").
		From("user_cards").
		Where("user_id = $1").
		Build()

	if id != 0 {
		q = r.queryBuilder.Select("*").From("card_transaction").Where("fromcard_id").In(subquery).OrderBy(params.SortBy + " " + params.SortOrder).Build()
		rows, err = r.replica.Query(q, id)
	} else {
		//"SELECT * FROM card_transaction;"
		q = r.queryBuilder.Select("*").From("card_transaction").OrderBy(params.SortBy + " " + params.SortOrder).Build()
		rows, err = r.replica.Query(q)
	}

	if err != nil {
		return []entity.CardTransaction{}, err
	}

	for rows.Next() {
		var temp entity.CardTransaction
		err := rows.Scan(&temp.ID, &temp.FromCardID, &temp.ToCardRequsite, &temp.ToCardName, &temp.ToCardSurname, &temp.Money, &temp.Time)
		if err != nil {
			return []entity.CardTransaction{}, err
		}

		data = append(data, temp)
	}

	return data, nil
}

func (r *Repo) DeleteCardPayment(id int) error {
	q := "DELETE FROM card_transaction WHERE id = $1;"

	_, err := r.main.Exec(q, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repo) GetPaymentByID(userId int, id int) (data entity.CardTransaction, err error) {
	q := "SELECT * FROM card_transaction WHERE " +
		"fromcard_id IN (SELECT card_id FROM user_cards WHERE user_id = $1) AND id = $2;"
	rows, err := r.replica.Query(q, userId, id)
	if err != nil {
		return entity.CardTransaction{}, err
	}

	for rows.Next() {
		err := rows.Scan(&data.ID, &data.FromCardID, &data.ToCardRequsite, &data.ToCardName, &data.ToCardSurname, &data.Money, &data.Time)
		if err != nil {
			return entity.CardTransaction{}, err
		}
	}

	return data, nil
}

func (r *Repo) CreateCard(id int, card entity.Card) error {
	q := "INSERT INTO user_cards(user_id) VALUES ($1) RETURNING card_id;"

	var l int
	rows, err := r.replica.Query(q, id)
	if err != nil {
		return err
	}

	for rows.Next() {
		err = rows.Scan(&l)
		if err != nil {
			return err
		}
	}
	fmt.Println(id, card, l)
	q = "INSERT INTO card (id, requisite, exp, cvc, full_name) VALUES ($1, $2, $3, $4, $5) "

	_, err = r.main.Exec(q, l, card.Requisite, card.Exp, card.CVC, card.FullName)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repo) GetCards(id int) ([]entity.Card, error) {

	//q := "SELECT * FROM card WHERE id in (SELECT card_id FROM user_cards WHERE user_id = $1)"
	subquery := r.queryBuilder.Select("card_id").From("user_cards").Where("user_id = $1").Build()
	q := r.queryBuilder.Select("*").From("card").Where("id").In(subquery).Build()

	rows, err := r.replica.Query(q, id)
	if err != nil {
		return []entity.Card{}, nil
	}

	var cards []entity.Card
	for rows.Next() {
		var card entity.Card
		err := rows.Scan(&card.ID, &card.Requisite, &card.Exp, &card.CVC, &card.FullName)
		if err != nil {
			return []entity.Card{}, nil
		}
		cards = append(cards, card)
	}

	return cards, nil
}

func (r *Repo) CreatePayment(transaction entity.CardTransaction) error {
	q := "SELECT COUNT(*) FROM card;"

	var l int
	rows, err := r.replica.Query(q)
	if err != nil {
		return err
	}

	for rows.Next() {
		err = rows.Scan(&l)
		if err != nil {
			return err
		}
	}

	q = "INSERT INTO card_transaction (fromcard_id, tocardrequisite, tocardname, tocardsurname, money, time) VALUES($1, $2, $3, $4, $5, $6);"

	_, err = r.main.Exec(q, transaction.FromCardID, transaction.ToCardRequsite, transaction.ToCardName, transaction.ToCardSurname, transaction.Money, time.Now())
	if err != nil {
		return err
	}

	return nil
}

func (r *Repo) UpdateCard(id int, card entity.Card) error {
	q := "UPDATE card SET requisite = $1, exp = $2, cvc = $3, full_name = $4 WHERE id = $5"
	_, err := r.main.Exec(q, card.Requisite, card.Exp, card.CVC, card.FullName, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repo) DeleteCard(id int) error {
	fmt.Println(id)
	q := "DELETE FROM card WHERE id = $1"

	_, err := r.main.Exec(q, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repo) CreateAccount(account entity.Account) error {
	q := r.queryBuilder.
		Insert("user_balance", "user_id", "currency", "balance").
		Values("$1", "$2", "$3").
		Build()

	_, err := r.main.Exec(q, account.UserID, account.Currency, account.Balance)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repo) GetAccounts(id int) (accounts []entity.Account, err error) {
	q := r.queryBuilder.
		Select("*").
		From("user_balance").
		Where("user_id = $1").
		Build()

	rows, err := r.replica.Query(q, id)
	if err != nil {
		return []entity.Account{}, err
	}

	for rows.Next() {
		var account entity.Account
		err = rows.Scan(&account.Id, &account.UserID, &account.Currency, &account.Balance)
		if err != nil {
			return []entity.Account{}, nil
		}
		accounts = append(accounts, account)
	}

	return accounts, nil
}

func (r *Repo) GetAccount(userid int, id int) (account entity.Account, err error) {
	q := "SELECT * FROM user_balance WHERE user_id = $1 AND id = $2;"

	rows, err := r.replica.Query(q, userid, id)
	if err != nil {
		return entity.Account{}, err
	}

	for rows.Next() {
		err = rows.Scan(&account.Id, &account.UserID, &account.Currency, &account.Balance)
		if err != nil {
			return entity.Account{}, nil
		}
	}

	return account, nil
}

func (r *Repo) DeleteAccount(id int) (err error) {
	q := "DELETE FROM user_balance WHERE id = $1"

	_, err = r.main.Exec(q, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repo) CreateAccountPayment(transaction entity.AccountTransaction) error {
	q := "SELECT balance FROM user_balance WHERE user_id = $1;"
	rows, err := r.replica.Query(q, transaction.FromID)
	var balance int
	for rows.Next() {
		err = rows.Scan(&balance)
		if err != nil {
			return err
		}
	}

	if balance-transaction.Money < 0 {
		return err
	}

	fmt.Printf("USER BALANCE id %d", balance)

	q = "UPDATE user_balance SET balance = balance - $1 WHERE user_id = $2 AND currency = $3;"

	_, err = r.main.Exec(q, transaction.Money, transaction.FromID, transaction.Currency)
	if err != nil {
		return err
	}

	q = "UPDATE user_balance SET balance = balance + $1 WHERE user_id = $2 AND currency = $3;"

	_, err = r.main.Exec(q, transaction.Money, transaction.ToID, transaction.Currency)
	if err != nil {
		return err
	}

	q = "INSERT INTO p2p_transaction (from_id, to_id, currency, money, message, time) VALUES ($1, $2, $3, $4, $5, $6);"

	_, err = r.main.Exec(q, transaction.FromID, transaction.ToID, transaction.Currency, transaction.Money, transaction.Message, time.Now())
	if err != nil {
		return err
	}

	return nil
}

func (r *Repo) GetAccountPayments(id int) (transactions []entity.AccountTransaction, err error) {
	q := "SELECT * FROM p2p_transaction WHERE from_id = $1 OR to_id = $1"

	rows, err := r.replica.Query(q, id)
	if err != nil {
		return []entity.AccountTransaction{}, err
	}

	for rows.Next() {
		var transaction entity.AccountTransaction
		err = rows.Scan(&transaction.ID, &transaction.FromID, &transaction.ToID, &transaction.Currency, &transaction.Money, &transaction.Message, &transaction.Time)
		if err != nil {
			return []entity.AccountTransaction{}, err
		}
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

func (r *Repo) GetAccountPayment(userid int, id int) (transaction entity.AccountTransaction, err error) {
	q := "SELECT * FROM p2p_transaction WHERE (from_id = $1 OR to_id = $1) AND id = $2;"

	rows, err := r.replica.Query(q, userid, id)
	if err != nil {
		return entity.AccountTransaction{}, err
	}

	for rows.Next() {
		err := rows.Scan(&transaction.ID, &transaction.FromID, &transaction.ToID, &transaction.Currency, &transaction.Money, &transaction.Message, &transaction.Time)
		if err != nil {
			return entity.AccountTransaction{}, err
		}
	}

	return transaction, nil
}
