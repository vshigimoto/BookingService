package entity

import "time"

type Card struct {
	ID        int
	Requisite string `json:"requisite"`
	Exp       string `json:"exp"`
	CVC       int    `json:"cvc"`
	FullName  string `json:"fullName"`
}

type TokenResponse struct {
	UserID       int    `json:"UserID"`
	Token        string `json:"AccessToken"`
	RefreshToken string `json:"RefreshToken"`
}

type UserCards struct {
	UserID int
	CardID int
}

type CardTransaction struct {
	ID             int
	ToCardRequsite string `json:"to_card_requisite"`
	ToCardName     string `json:"to_card_name"`
	ToCardSurname  string `json:"to_card_surname"`
	FromCardID     int    `json:"from_card_id"`
	Money          int    `json:"money"`
	Time           time.Time
}

type Account struct {
	Id       int
	UserID   int
	Currency string `json:"currency"`
	Balance  int    `json:"balance"`
}

type AccountTransaction struct {
	ID       int
	FromID   int    `json:"from_id"`
	ToID     int    `json:"to_id"`
	Currency string `json:"currency"`
	Money    int    `json:"money"`
	Message  string `json:"message"`
	Time     time.Time
}

type QueryParams struct {
	SortBy      string
	SortOrder   string
	Search      string
	SearchQuery string
}
