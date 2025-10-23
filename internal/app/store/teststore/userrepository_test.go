package teststore_test

import (
	"testing"

	"github.com/golang-edu-project/internal/app/model"
	"github.com/golang-edu-project/internal/app/store"
	"github.com/golang-edu-project/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	s := teststore.New()

	u := model.TestUser(t)

	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepository_GetByEmail(t *testing.T) {
	s := teststore.New()
	email := "a.golubev@qsoft.ru"

	_, err := s.User().GetByEmail(email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	u := model.TestUser(t)

	s.User().Create(u)
	u, err = s.User().GetByEmail(email)

	assert.NoError(t, err)
	assert.NotNil(t, u)
}
