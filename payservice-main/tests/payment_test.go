package tests

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"testing"
)

/*
{
    "requisite": "4400430282842811",
    "exp": "06/24",
    "cvc": 111,
    "fullName": "GABDYLGAZIZ ZHAGYPAR"
}
*/

type PaymentPostBody struct {
	Requisite string `json:"requisite"`
	Exp       string `json:"exp"`
	CVC       int    `json:"cvc"`
	FullName  string `json:"fullName"`
}

func Request(body PaymentPostBody, url url.URL, method string, token string) int {
	bodyBytes, err := json.Marshal(&body)
	if err != nil {
		log.Fatal(err)
	}

	reader := bytes.NewReader(bodyBytes)
	req, err := http.NewRequest(method, url.String(), reader)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	return resp.StatusCode
}

func TestSuccessGetCards(t *testing.T) {
	body := PaymentPostBody{
		Requisite: "",
		Exp:       "",
		CVC:       0,
		FullName:  "",
	}

	Url := url.URL{
		Host:   "localhost:8082",
		Path:   "/v2/cards",
		Scheme: "http",
	}

	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjgsImV4cCI6MTcwMDQxMDE3M30.32zlD4vlrsKJ-4nd8Ap7JdkfPq0FgZvCUJ6eD2eTU00"

	result := Request(body, Url, "GET", token)

	if result >= 400 && result <= 500 {
		t.Error("Error response. Status Code: ", result)
		return
	}
}

func TestFailedGetCards(t *testing.T) {
	body := PaymentPostBody{
		Requisite: "",
		Exp:       "",
		CVC:       0,
		FullName:  "",
	}

	Url := url.URL{
		Host:   "localhost:8082",
		Path:   "/v2/cards",
		Scheme: "http",
	}

	token := "random_token"

	result := Request(body, Url, "GET", token)

	if result == 200 {
		t.Error("Error response. Status Code: ", result)
		return
	}
}

func TestSuccessCreateCard(t *testing.T) {
	body := PaymentPostBody{
		Requisite: "test",
		Exp:       "test",
		CVC:       1,
		FullName:  "test",
	}

	Url := url.URL{
		Host:   "localhost:8082",
		Path:   "/v2/card",
		Scheme: "http",
	}

	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjgsImV4cCI6MTcwMDQxMDE3M30.32zlD4vlrsKJ-4nd8Ap7JdkfPq0FgZvCUJ6eD2eTU00"

	result := Request(body, Url, "POST", token)

	if result >= 400 && result <= 500 {
		t.Error("Error response. Status Code: ", result)
		return
	}
}
