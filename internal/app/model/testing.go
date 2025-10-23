package model

import "testing"

func TestUser(t *testing.T) *User {
	return &User{
		Email:       "a.golubev@qsoft.ru",
		TmpPassword: "12345678",
	}
}
