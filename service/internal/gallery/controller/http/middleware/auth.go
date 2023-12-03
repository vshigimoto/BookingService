package middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"go.uber.org/zap"

	"service/internal/gallery/auth"
)

const (
	AuthorizationHeaderKey = "Authorization"
)

type JwtV1 struct {
	authService auth.UseCase
	logger      *zap.SugaredLogger
}

func NewJwtV1Middleware(authService auth.UseCase, logger *zap.SugaredLogger) *JwtV1 {
	return &JwtV1{
		authService: authService,
		logger:      logger,
	}
}

func (j *JwtV1) AuthV1(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, ok := r.Header[AuthorizationHeaderKey]; !ok {
			j.logger.Warn("'Authorization' key missing from headers")

			w.WriteHeader(http.StatusBadRequest)
			return
		}

		jwtToken, ok := j.getTokenFromHeader(r)
		if !ok {
			j.logger.Warn(fmt.Sprintf(
				"failed to getTokenFromHeader invalidToken: %s",
				r.Header.Get(AuthorizationHeaderKey),
			))

			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		authService := j.authService

		contextUser, err := authService.GetJwtUser(jwtToken)
		if err != nil {
			if !errors.Is(err, auth.ErrExpiredToken) {
				j.logger.Errorf("failed to GetJwtUser err: %v", err)
			}

			w.WriteHeader(http.StatusUnauthorized)
		} else {
			newCtx := context.WithValue(r.Context(), authService.GetContextUserKey(), contextUser)
			next.ServeHTTP(w, r.WithContext(newCtx))
		}
	})
}

func (j *JwtV1) getTokenFromHeader(r *http.Request) (string, bool) {
	bearer := r.Header.Get(AuthorizationHeaderKey)
	if len(bearer) > 7 && strings.ToUpper(bearer[0:6]) == "BEARER" {
		return bearer[7:], true
	}

	return "", false
}

func GetContextUser(ctx context.Context) (*auth.ContextUser, error) {
	if user, ok := ctx.Value(auth.ContextUserKey{}).(*auth.ContextUser); ok {
		return user, nil
	}

	return nil, errors.New("could not find user by contextUserKey{}")
}
