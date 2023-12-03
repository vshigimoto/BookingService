package repository

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"payservice/internal/user/entity"
)

func (r *Repo) CreateUser(user entity.User) (id int, err error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}
	fmt.Println(user.Password, hashedPassword)

	q := "INSERT INTO users(first_name, last_name, phone, username, password, isconfirmed) VALUES($1, $2, $3, $4, $5, $6) RETURNING id;"
	row, err := r.main.Query(q, user.FirstName, user.LastName, user.Phone, user.Login, string(hashedPassword), false)
	if err != nil {
		return 0, err
	}

	for row.Next() {
		err = row.Scan(&id)
		if err != nil {
			return 0, err
		}
	}

	return id, nil
}

func (r *Repo) GetAllUsers() []entity.User {
	var users []entity.User
	q := "SELECT * FROM users;"
	rows, err := r.replica.Query(q)
	if err != nil {
		return []entity.User{}
	}

	for rows.Next() {
		var user entity.User
		err = rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Phone, &user.Login, &user.Password, &user.IsConfirmed)
		if err != nil {
			panic(err)
		}
		users = append(users, user)
	}

	return users
}

func (r *Repo) UpdateUser(id string, user entity.User) (err error) {
	q := "UPDATE users SET first_name = $2, last_name = $3, password = $4 WHERE id = $1;"
	_, err = r.main.Query(q, id, user.FirstName, user.LastName, user.Password)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repo) DeleteUser(id string) (err error) {
	q := "DELETE FROM users WHERE id = $1;"

	_, err = r.main.Query(q, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repo) GetUserByID(id string) entity.User {
	var user entity.User
	q := "SELECT * FROM users WHERE id = $1;"
	rows, err := r.replica.Query(q, id)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		err = rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Phone, &user.Login, &user.Password, &user.IsConfirmed)
		if err != nil {
			panic(err)
		}
	}

	return user
}

func (r *Repo) GetUserByLogin(login string) entity.User {
	var user entity.User
	q := "SELECT * FROM users WHERE username = $1;"
	rows, err := r.replica.Query(q, login)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		err = rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Phone, &user.Login, &user.Password, &user.IsConfirmed)
		if err != nil {
			panic(err)
		}
	}

	return user
}
