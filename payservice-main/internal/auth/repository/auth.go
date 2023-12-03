package repository

import (
	"payservice/internal/auth/entity"
	"payservice/internal/auth/metrics"
)

func (r *Repo) CreateToken(u entity.UserToken) error {
	ok, fail := metrics.DatabaseQueryTime("CreateUserToken")
	defer fail()

	q := `
	INSERT INTO user_token (token, refresh_token, user_id)
	VALUES ($1, $2, $3)
	ON CONFLICT (user_id)
	DO UPDATE SET
		token = EXCLUDED.token,
		refresh_token = EXCLUDED.refresh_token;
	`

	_, err := r.main.Exec(q, u.AccessToken, u.RefreshToken, u.UserID)

	if err != nil {
		return err
	}

	ok()

	return nil
}

func (r *Repo) UpdateToken() (err error) {
	q := "UPDATE user_token SET token = $1, refresh_token = $2 WHERE user_id = $3;"

	_, err = r.main.Exec(q)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repo) GetToken(token string) (entity.UserToken, error) {
	q := "SELECT user_id, token, refresh_token FROM user_token WHERE token = $1"
	var userToken entity.UserToken

	rows, err := r.main.Query(q, token)
	if err != nil {
		return userToken, err
	}

	for rows.Next() {
		err = rows.Scan(&userToken.UserID, &userToken.AccessToken, &userToken.RefreshToken)
		if err != nil {
			return userToken, err
		}
	}

	return userToken, nil
}

func (r *Repo) CreateUserCode(id int, Code string) error {
	ok, fail := metrics.DatabaseQueryTime("CreateUserCode")
	defer fail()

	q := "INSERT INTO usercode(user_id, code) VALUES ($1, $2);"

	_, err := r.main.Exec(q, id, Code)
	if err != nil {
		return err
	}

	ok()
	return nil
}

func (r *Repo) ConfirmUserCode(Code string) error {
	q := "DELETE FROM usercode WHERE code = $1 RETURNING user_id;"

	rows, err := r.main.Query(q, Code)
	if err != nil {
		return err
	}

	var id int
	for rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			return err
		}
	}

	q = "UPDATE users SET isconfirmed = true WHERE id = $1;"

	_, err = r.main.Query(q, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repo) DeleteToken(id int) (err error) {
	q := "DELETE FROM user_token WHERE user_id = $1;"

	_, err = r.main.Exec(q, id)
	if err != nil {
		return err
	}

	return nil
}
