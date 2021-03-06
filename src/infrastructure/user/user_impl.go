package user

import (
	"github.com/rikiya/go-clean/src/entity"
	"github.com/rikiya/go-clean/src/infrastructure/database"
)

// UserImpl ...
type UserImpl struct {
	database.SQLHandler
}

// Store ...
func (ui *UserImpl) Store(u entity.User) error {
	ctx := database.NewSQLHandler()
	res, err := ctx.Prepare(
		`INSERT INTO users (first_name, last_name) VALUES(?, ?)`,
	)
	if err != nil {
		return err
	}
	defer res.Close()

	_, err = res.Exec(u.FirstName, u.LastName)
	return err
}

// Index ...
func (ui *UserImpl) Index() (users []entity.User, err error) {
	ctx := database.NewSQLHandler()
	if err = ctx.Select(&users, `SELECT * FROM users`); err != nil {
		return nil, err
	}
	return users, nil
}

// Update ...
func (ui *UserImpl) Update(id int, u entity.User) error {
	ctx := database.NewSQLHandler()
	res, err := ctx.Prepare(
		`UPDATE users SET first_name = ?, last_name = ? WHERE id = ?`,
	)
	if err != nil {
		return err
	}
	defer res.Close()

	_, err = res.Exec(u.FirstName, u.LastName, id)
	return err
}

// Delete ...
func (ui *UserImpl) Delete(id int) error {
	ctx := database.NewSQLHandler()
	res, err := ctx.Prepare(
		`DELETE FROM users where id = ?`,
	)
	if err != nil {
		return err
	}
	defer res.Close()

	_, err = res.Exec(id)
	return err
}
