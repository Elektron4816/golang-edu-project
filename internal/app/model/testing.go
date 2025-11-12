package model

import "testing"

func TestUser(t *testing.T) *User {
	return &User{
		Email:    "a.golubev@qsoft.ru",
		Password: "12345678",
	}
}
