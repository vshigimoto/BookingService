package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"net/http"
	"strings"
)

func AdminVerify() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var tokenString string
		tokenHeader := ctx.Request.Header.Get("Authorization")
		tokenFields := strings.Fields(tokenHeader)
		if len(tokenFields) == 2 && tokenFields[0] == "Bearer" {
			tokenString = tokenFields[1]
		} else {
			ctx.AbortWithStatus(http.StatusForbidden)
			return
		}
		claims := jwt.MapClaims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method %v", token.Header["alg"])
			}
			return []byte("AllYourBase"), nil
		})
		if err != nil {
			ctx.AbortWithStatus(http.StatusForbidden)
			return
		}
		if !token.Valid {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		userRole, ok := claims["user_role"]
		if !ok {
			log.Printf("user id could not be parsed from JWT")
			return
		}
		if userRole != "admin" {
			ctx.AbortWithStatus(http.StatusForbidden)
		}
		ctx.Set("user_role", userRole)
		ctx.Next()
	}
}
