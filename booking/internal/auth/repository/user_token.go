package repository

import (
	"booking/internal/auth/entity"
)

func (r *Repo) CreateUserToken(userToken entity.UserToken) error {
	_, err := r.main.Exec("insert into user_token(user_id, token, refresh_token) VALUES($1,$2,$3)", userToken.UserId, userToken.Token, userToken.RefreshToken)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repo) UpdateUserToken(userToken entity.UserToken) error {
	_, err := r.main.Exec("UPDATE user_token SET token=$1, refresh_token=$2 WHERE user_id=$3", userToken.Token, userToken.RefreshToken, userToken.UserId)
	if err != nil {
		return err
	}
	return nil
}
