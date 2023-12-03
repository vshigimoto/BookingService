package http

import (
	"booking/internal/auth/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

type EndpointHandler struct {
	authService auth.Service
}

func NewEndpointHandler(authService auth.Service) *EndpointHandler {
	return &EndpointHandler{
		authService: authService,
	}
}

func (h *EndpointHandler) Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		request := struct {
			Login    string `json:"login"`
			Password string `json:"password"`
		}{}

		err := ctx.ShouldBindJSON(&request)
		if err != nil {
			// log it
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "bad json"})
			return
		}

		tokenRequest := auth.GenerateTokenRequest{
			Login:    request.Login,
			Password: request.Password,
		}

		userToken, err := h.authService.GenerateToken(ctx, tokenRequest)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "bad json"})
			return
		}

		response := struct {
			Token        string `json:"token"`
			RefreshToken string `json:"refresh_token"`
		}{
			Token:        userToken.Token,
			RefreshToken: userToken.RefreshToken,
		}

		ctx.JSON(http.StatusCreated, response)
	}
}

func (h *EndpointHandler) Confirm() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
