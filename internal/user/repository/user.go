package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/vshigimoto/BookingService/internal/user/entity"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func isEqual[T comparable](user T, dbUser T) bool {
	return user == dbUser
}

func (r *Repo) isExist(login string) bool {
	user, err := r.GetByLogin(context.Background(), login)
	if err != nil {
		if err.Error() == "cannot scan query with error: sql: Rows are closed" {
			return false
		}
		log.Printf("error with get user by login: %v", err)
		return true
	}
	ok := isEqual(user.Login, login)
	return ok
}

func (r *Repo) CreateUser(ctx context.Context, user *entity.User) (id int, err error) {
	if exist := r.isExist(user.Login); exist {
		return 0, fmt.Errorf("user with login %s is exist", user.Login)
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, fmt.Errorf("cannot hash password with error: %v", err)
	}
	err = r.main.QueryRow("insert into users(name, email,login, password) values($1, $2, $3, $4) returning ID", user.Name, user.Email, user.Login, string(hashedPassword)).Scan(&user.Id) // $1 and $2 is prepared statement
	if err != nil {
		return 0, fmt.Errorf("cannot query with error: %v", err)
	}
	return user.Id, nil
}

func (r *Repo) GetUsers(ctx context.Context, sortKey, sortBy string) ([]entity.User, error) {
	users := make([]entity.User, 0)
	rows, err := r.replica.Query("SELECT * from users")
	if err != nil {
		return nil, fmt.Errorf("cannot query with error: %v", err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			return
		}
	}(rows)
	for rows.Next() {
		var user entity.User
		if err = rows.Scan(&user.Id, &user.Name, &user.Email, &user.Login, &user.Password); err != nil {
			return nil, fmt.Errorf("cannot scan query with error: %v", err)
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows with error: %v", err)
	}
	return users, nil
}

func (r *Repo) UpdateUser(ctx context.Context, login string, user *entity.User) error {
	_, err := r.GetByLogin(context.TODO(), login)
	if err != nil {
		return fmt.Errorf("cannot get user with login %s", login)
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("cannot hash password with err:%v", err)
	}
	_, err = r.main.Exec("UPDATE users SET name=$1, email=$2, password=$3 WHERE login=$4", user.Name, user.Email, hashedPassword, login)
	if err != nil {
		return fmt.Errorf("cannot query with err:%v", err)
	}
	return nil
}

func (r *Repo) DeleteUser(ctx context.Context, login string) error {
	_, err := r.GetByLogin(context.TODO(), login)
	if err != nil {
		return fmt.Errorf("cannot get user with login %s", login)
	}
	_, err = r.main.Exec("DELETE from users WHERE login=$1", login)
	if err != nil {
		return fmt.Errorf("cannot delete user with err:%v", err)
	}
	return nil
}

func (r *Repo) GetByLogin(ctx context.Context, login string) (*entity.User, error) {
	rows, err := r.replica.Query("SELECT * FROM users WHERE login=$1", login)
	if err != nil {
		return nil, fmt.Errorf("error with query: %v", err)
	}
	var user entity.User
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			return
		}
	}(rows)
	rows.Next()
	if err = rows.Scan(&user.Id, &user.Name, &user.Email, &user.Login, &user.Password); err != nil {
		return nil, fmt.Errorf("cannot scan query with error: %v", err)
	}
	return &user, nil
}
