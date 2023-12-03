package http

import (
	"encoding/json"
	"io"
	"net/http"

	"go.uber.org/zap"

	"service/internal/auth/auth"
)

type EndpointHandler struct {
	authService auth.UseCase
	logger      *zap.SugaredLogger
}

func NewEndpointHandler(
	authService auth.UseCase,
	logger *zap.SugaredLogger,
) *EndpointHandler {
	return &EndpointHandler{
		authService: authService,
		logger:      logger,
	}
}

func (eh *EndpointHandler) Register(w http.ResponseWriter, r *http.Request) {
	err := eh.authService.Register(r.Context())
	if err != nil {
		eh.logger.Errorf("failed to Register err: %v", err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	eh.logger.Info("user successfully is registered")
}

func (eh *EndpointHandler) ConfirmUser(w http.ResponseWriter, r *http.Request) {
	// check code in database

	// if ok
	// then update user through user-service by grpc
	// set is_confirmed = true
	// update set user is_confirmed = true where id = ?
}

func (eh *EndpointHandler) Login(w http.ResponseWriter, r *http.Request) {
	logger := eh.logger.With(
		zap.String("endpoint", "login"),
		zap.String("params", r.URL.String()),
	)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Errorf("failed to read body err: %v", err)
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	request := struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}{}

	err = json.Unmarshal(body, &request)
	if err != nil {
		logger.Errorf("failed to unmarshal body err: %v", err)
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	tokenRequest := auth.GenerateTokenRequest{
		Login:    request.Login,
		Password: request.Password,
	}

	userToken, err := eh.authService.GenerateToken(r.Context(), tokenRequest)
	if err != nil {
		logger.Errorf("failed to GenerateToken err: %v", err)
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	response := struct {
		Token        string `json:"token"`
		RefreshToken string `json:"refresh_token"`
	}{
		Token:        userToken.Token,
		RefreshToken: userToken.RefreshToken,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		//your err handle code
	}
}

func (eh *EndpointHandler) RenewToken(w http.ResponseWriter, r *http.Request) {

}
