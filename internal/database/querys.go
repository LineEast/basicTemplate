package database

import (
	"context"
	"kmfRedirect/internal/models"
)

// CRUD

func (db *Database) Create(ctx context.Context, r *models.User) (err error) {
	err = db.Pool.QueryRow(
		ctx,
		"insert into users (name, phoneNumber) values ($1, $2) returning id",
		r.Name, r.PhoneNumber,
	).Scan(&r.ID)

	return
}

func (db *Database) ReadAllUserList(ctx context.Context) (userList []models.User, err error) {
	rows, err := db.Pool.Query(
		ctx,
		"select id, name, phoneNumber from users",
	)
	if err != nil {
		return
	}

	user := models.User{}
	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Name, &user.PhoneNumber)
		if err != nil {
			return
		}

		userList = append(userList, user)
	}

	return
}

func (db *Database) ReadUser(ctx context.Context, user *models.User) (err error) {
	err = db.Pool.QueryRow(
		ctx,
		"select name, phoneNumber from users where id = $1",
		user.ID,
	).Scan(&user.Name, &user.PhoneNumber)

	return
}
