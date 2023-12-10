package http

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/vshigimoto/BookingService/internal/auth/auth"
	redisClient "github.com/vshigimoto/BookingService/internal/auth/redis"
	"log"
	"net/http"
	"time"
)

type EndpointHandler struct {
	authService auth.Service
}

func New(authService auth.Service) *EndpointHandler {
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

		if err := ctx.ShouldBindJSON(&request); err != nil {
			log.Print("cannot unmarshall json")
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		rdb := redisClient.NewRedisClient()
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

func (h *EndpointHandler) RenewToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		request := struct {
			Login    string `json:"login"`
			Password string `json:"password"`
			Token    string `json:"token"`
		}{}

		err := ctx.ShouldBindJSON(&request)
		if err != nil {
			log.Print("cannot unmarshall json")
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		rdb := redisClient.NewRedisClient()

		_, err = rdb.Get(request.Token).Result()
		if err == redis.Nil {
			log.Print("token expired or not exist")
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		} else if err != nil {
			log.Printf("failed get token err: %v", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
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
