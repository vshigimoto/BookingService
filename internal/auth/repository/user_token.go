package repository

import (
	"fmt"
	"github.com/vshigimoto/BookingService/internal/auth/entity"
)

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
