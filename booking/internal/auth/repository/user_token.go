package repository

import (
	"booking/internal/auth/entity"
	"fmt"
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

func (r *Repo) GetUserRole(userId int) (string, error) {
	rows, err := r.replica.Query("SELECT * FROM user_role WHERE user_id=$1", userId)
	if err != nil {
		return "", err
	}
	ok := rows.Next()
	if !ok {
		return "", fmt.Errorf("user with id %d does not exist", userId)
	}
	var userRole entity.UserRole
	if err := rows.Scan(&userRole.Id, &userRole.UserId, &userRole.Role); err != nil {
		return "", err
	}
	return userRole.Role, nil
}
