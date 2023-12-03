package tests

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"testing"
)

type PostBody struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func PostRequest(body PostBody) int {
	postURL := url.URL{
		Host:   "localhost:8081",
		Path:   "/auth/login",
		Scheme: "http",
	}

	bodyBytes, err := json.Marshal(&body)
	if err != nil {
		log.Fatal(err)
	}

	reader := bytes.NewReader(bodyBytes)
	resp, err := http.Post(postURL.String(), "application/json", reader)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	return resp.StatusCode
}

func TestSuccessAuth(t *testing.T) {
	body := PostBody{
		Login:    "ggwp",
		Password: "ggwp",
	}

	result := PostRequest(body)

	if result >= 400 && result <= 500 {
		t.Error("Error response. Status Code: ", result)
		return
	}
}

func TestFailAuth(t *testing.T) {
	body := PostBody{
		Login:    "random",
		Password: "random",
	}

	result := PostRequest(body)

	if result == 200 {
		t.Error("Auth is success: ", result)
		return
	}
}
