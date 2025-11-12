package model_test

import (
	"testing"

	"github.com/golang-edu-project/internal/app/model"
	"github.com/stretchr/testify/assert"
)

func TestUser_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		u       func() *model.User
		isValid bool
	}{
		{
			name: "valid",
			u: func() *model.User {
				return model.TestUser(t)
			},
			isValid: true,
		},
		{
			name: "with password",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Password = "TmpPassword"
				u.EncryptedPassword = "TmpPassword"
				return u
			},
			isValid: true,
		},
		{
			name: "empty email",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Email = ""
				return u
			},
			isValid: false,
		},
		{
			name: "invalid email",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Email = "invalid email"
				return u
			},
			isValid: false,
		},
		{
			name: "invalid tmpPassword",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Password = "12345"
				return u
			},
			isValid: false,
		},
		{
			name: "empty tmpPassword",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Password = ""
				return u
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.u().Validate())
			} else {
				assert.Error(t, tc.u().Validate())
			}
		})
	}
}

func TestUser_BeforeCreate(t *testing.T) {
	u := model.TestUser(t)

	assert.NoError(t, u.BeforeCreate())
	assert.NotNil(t, u.Password)
}
