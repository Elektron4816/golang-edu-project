package model

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id          int
	Email       string
	TmpPassword string
	Password    string
}

func (u *User) Validate() error {
	return validation.ValidateStruct(
		u,
		validation.Field(
			&u.Email,
			validation.Required,
			is.Email,
		),
		validation.Field(
			&u.TmpPassword,
			validation.By(requiredIf(u.Password == "")),
			validation.Length(8, 100),
		),
	)
}

func (u *User) BeforeCreate() error {
	if len(u.TmpPassword) > 0 {
		enc, err := encryptSitring(u.TmpPassword)
		if err != nil {
			return err
		}

		u.Password = enc
	}
	return nil
}

func encryptSitring(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)

	if err != nil {
		return "", err
	}

	return string(b), nil
}
