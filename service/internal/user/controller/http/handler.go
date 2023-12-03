package http

import (
	"encoding/json"
	"errors"
	"io/fs"
	"mime"
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"

	"service/internal/user/repository"
	"service/internal/user/user"
	"service/swagger"
)

type EndpointHandler struct {
	logger      *zap.SugaredLogger
	userService user.UseCase
}

func NewEndpointHandler(
	logger *zap.SugaredLogger,
	userService user.UseCase,
) *EndpointHandler {
	return &EndpointHandler{
		logger:      logger,
		userService: userService,
	}
}

// swagger:model UserResponse
type UserResponse struct {
	// example: 1
	Id int `json:"id"`
	// example: mytest
	Login string `json:"login"`
	// example: fkdslkf
	FirstName string `json:"first_name"`
	// example: mytest3
	LastName string `json:"last_name"`
	// example: true
	IsConfirmed bool `json:"is_confirmed"`
	// example: mytest4
	Password string `json:"password"`
}

// swagger:route GET /v1/user/{login} GetUser
//
//	Parameters:
//	 + name: login
//	   in: path
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Schemes: http, https
//
//	Security:
//	  Bearer:
//
//	Responses:
//	  200: UserResponse
func (h *EndpointHandler) GetUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		login := chi.URLParam(r, "login")

		userEntity, err := h.userService.GetUserByLogin(r.Context(), login)
		if err != nil {
			if errors.Is(err, repository.ErrNotFound) {
				h.logger.Warnf("user with login: %s not found", login)
				w.WriteHeader(http.StatusNotFound)
				json.NewEncoder(w).Encode(nil)
				return
			}

			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		response := UserResponse{
			Id:          userEntity.Id,
			Login:       userEntity.Login,
			FirstName:   userEntity.FirstName,
			LastName:    userEntity.LastName,
			IsConfirmed: userEntity.IsConfirmed,
			Password:    userEntity.Password,
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}

//func (h *EndpointHandler) GetUser(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//
//	login := chi.URLParam(r, "login")
//
//	userEntity, err := h.userService.GetUserByLogin(r.Context(), login)
//	if err != nil {
//		if errors.Is(err, repository.ErrNotFound) {
//			h.logger.Warnf("user with login: %s not found", login)
//			w.WriteHeader(http.StatusNotFound)
//			json.NewEncoder(w).Encode(nil)
//			return
//		}
//
//		w.WriteHeader(http.StatusInternalServerError)
//		return
//	}
//
//	// swagger:model userResponse
//	userResponse := struct {
//		// example: credit-card
//		Id int `json:"id"`
//		// example: credit-card
//		Login string `json:"login"`
//		// example: credit-card
//		FirstName string `json:"first_name"`
//		// example: credit-card
//		LastName string `json:"last_name"`
//		// example: credit-card
//		IsConfirmed bool `json:"is_confirmed"`
//		// example: true
//		Password string `json:"password"`
//		// example: credit-card
//	}{
//		Id:          userEntity.Id,
//		Login:       userEntity.Login,
//		FirstName:   userEntity.FirstName,
//		LastName:    userEntity.LastName,
//		IsConfirmed: userEntity.IsConfirmed,
//		Password:    userEntity.Password,
//	}
//
//	w.WriteHeader(http.StatusOK)
//	json.NewEncoder(w).Encode(userResponse)
//}

type swaggerServer struct {
	openApi http.Handler
}

func (h *EndpointHandler) Swagger() http.Handler {
	if err := mime.AddExtensionType(".svg", "image/svg+xml"); err != nil {
		h.logger.Error("AddExtensionType mimetype error: %w", zap.Error(err))
	}

	openApi, err := fs.Sub(swagger.OpenAPI, "OpenAPI")
	if err != nil {
		panic("couldn't create sub filesystem: " + err.Error())
	}

	return &swaggerServer{
		openApi: http.StripPrefix("/swagger/", http.FileServer(http.FS(openApi))),
	}
}

func (sws *swaggerServer) ServeHTTP(w http.ResponseWriter, rq *http.Request) {
	sws.openApi.ServeHTTP(w, rq)
}
