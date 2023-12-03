package transport

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"service/internal/auth/config"
)

type UserTransport struct {
	config config.UserTransport
}

func NewTransport(config config.UserTransport) *UserTransport {
	return &UserTransport{
		config: config,
	}
}

type GetUserResponse struct {
	Id          int    `json:"id"`
	Login       string `json:"login"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	IsConfirmed string `yaml:"is_confirmed"`
	Password    string `json:"password"`
}

func (ut *UserTransport) GetUser(ctx context.Context, login string) (*GetUserResponse, error) {
	var response *GetUserResponse

	responseBody, err := ut.makeRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/user/v1/user/%s", login),
		ut.config.Timeout,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to makeRequest err: %w", err)
	}

	if err := json.Unmarshal(responseBody, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshall response err: %w", err)
	}

	return response, nil
}

func (ut *UserTransport) makeRequest(
	ctx context.Context,
	httpMethod string,
	endpoint string,
	timeout time.Duration,
) (b []byte, err error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	requestURL := ut.config.Host + endpoint

	req, err := http.NewRequestWithContext(ctx, httpMethod, requestURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to NewRequestWithContext err: %w", err)
	}

	httpClient := &http.Client{}

	res, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("client making http request err: %w", err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body err: %w", err)
	}

	return body, nil
}
