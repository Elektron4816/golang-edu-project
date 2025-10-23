package store

import "github.com/golang-edu-project/internal/app/model"

type UserRepository interface {
	Create(*model.User) error
	GetByEmail(string) (*model.User, error)
}
