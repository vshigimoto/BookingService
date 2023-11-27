package repository

import (
	"booking/internal/user/entity"
	"database/sql"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

// CreateUser Create function add new user to DB
func (r *Repo) CreateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user entity.User
		err := ctx.ShouldBindJSON(&user)
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return
		}
		//user.Name and user.Email is arguments for prepared statement ($1 and $2)
		err = r.main.QueryRow("insert into users(name, email,login, password) values($1, $2, $3, $4) returning ID", user.Name, user.Email, user.Login, string(hashedPassword)).Scan(&user.Id) // $1 and $2 is prepared statement
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": err})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"id": user.Id})
	}
}

func (r *Repo) GetUsers() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		users := make([]entity.User, 0)

		rows, err := r.replica.Query("SELECT * from users")
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
		}
		defer func(rows *sql.Rows) {
			err := rows.Close()
			if err != nil {

			}
		}(rows)
		for rows.Next() {
			var user entity.User
			if err = rows.Scan(&user.Id, &user.Name, &user.Email, &user.Login, &user.Password); err != nil {
				ctx.JSON(http.StatusInternalServerError, err)
				return
			}
			users = append(users, user)
		}
		if err = rows.Err(); err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		ctx.JSON(http.StatusOK, users)
	}

}

func (r *Repo) UpdateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		if id == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "id should not to be empty"})
			return
		}
		var user entity.User
		// Unmarshal json to a new user
		if err := ctx.ShouldBindJSON(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
			return
		}
		_, err := r.main.Exec("UPDATE users SET name=$1, email=$2, login=$3 WHERE id=$4", user.Name, user.Email, user.Login, id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
	}
}

func (r *Repo) DeleteUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		if id == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "id should not to be empty"})
			return
		}
		_, err := r.main.Exec("DELETE from users WHERE id=$1", id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
	}
}

func (r *Repo) GetByLogin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		login := ctx.Param("login")
		if login == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "id should not to be empty"})
			return
		}
		rows, err := r.replica.Query("SELECT * FROM users WHERE login=$1", login)
		var user entity.User
		defer func(rows *sql.Rows) {
			err := rows.Close()
			if err != nil {

			}
		}(rows)
		rows.Next()
		if err = rows.Scan(&user.Id, &user.Name, &user.Email, &user.Login, &user.Password); err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		ctx.JSON(http.StatusOK, user)
	}
}
