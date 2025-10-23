package sqlstore

import (
	"database/sql"

	"github.com/golang-edu-project/internal/app/model"
	"github.com/golang-edu-project/internal/app/store"
)

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}

	return r.store.db.QueryRow(
		"INSERT INTO users (email, password) VALUES ($1, $2) RETURNING id",
		u.Email,
		u.Password,
	).Scan(&u.Id)
}

func (r *UserRepository) GetByEmail(email string) (*model.User, error) {
	u := &model.User{}

	if err := r.store.db.QueryRow(
		"SELECT id, email, password FROM users WHERE email = $1",
		email,
	).Scan(&u.Id, &u.Email, &u.Password); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}

	return u, nil
}
