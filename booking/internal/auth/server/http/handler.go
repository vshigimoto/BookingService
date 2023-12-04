package http

import (
	"booking/internal/auth/auth"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"log"
	"net/http"
	"time"
)

type EndpointHandler struct {
	authService auth.Service
}

func NewEndpointHandler(authService auth.Service) *EndpointHandler {
	return &EndpointHandler{
		authService: authService,
	}
}

func newRedisClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return rdb
}

func (h *EndpointHandler) Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		request := struct {
			Login    string `json:"login"`
			Password string `json:"password"`
		}{}

		err := ctx.ShouldBindJSON(&request)
		if err != nil {
			log.Print("cannot unmarshall json")
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		rdb := newRedisClient()
		token, err := rdb.Get(request.Login).Result()
		if err == redis.Nil {
			log.Print("token expired or not exist")
		} else if err != nil {
			log.Printf("error with get from redis %v", err)
		} else {
			refresh, err := rdb.Get(token).Result()
			if err != nil {
				log.Printf("error with get refresh token %v", err)
			} else {
				response := struct {
					Token        string `json:"token"`
					RefreshToken string `json:"refresh_token"`
				}{
					Token:        token,
					RefreshToken: refresh,
				}
				ctx.JSON(http.StatusOK, response)
				return
			}
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
		err = rdb.Set(request.Login, userToken.Token, 15*time.Minute).Err()
		if err != nil {
			log.Printf("error with set token in redis:%v", err)
		}
		err = rdb.Set(userToken.Token, userToken.RefreshToken, 40*time.Minute).Err()
		if err != nil {
			log.Printf("error with set token in redis:%v", err)
		}
		ctx.JSON(http.StatusCreated, response)
	}
}

func (h *EndpointHandler) Confirm() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
