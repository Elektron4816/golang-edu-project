package apiserver

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang-edu-project/internal/app/model"
	"github.com/golang-edu-project/internal/app/store/teststore"
	"github.com/gorilla/sessions"
	"github.com/stretchr/testify/assert"
)

func TestServer_HandleUsersCreate(t *testing.T) {
	s := newServer(teststore.New(), sessions.NewCookieStore([]byte("secret")))
	testCases := []struct {
		name         string
		payload      interface{}
		expextedCode int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"email":    "a.golubev@qsoft.ru",
				"password": "123456345",
			},
			expextedCode: http.StatusCreated,
		},
		{
			name:         "invalid payload",
			payload:      "invalid payload",
			expextedCode: http.StatusBadRequest,
		},
		{
			name: "invalid params",
			payload: map[string]string{
				"email":    "a.golubev",
				"password": "123456345",
			},
			expextedCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()

			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req := httptest.NewRequest(http.MethodPost, "/users", b)

			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expextedCode, rec.Code)

		})
	}
}

func TestServer_HandleSessionsCreate(t *testing.T) {
	u := model.TestUser(t)

	store := teststore.New()
	store.User().Create(u)

	s := newServer(store, sessions.NewCookieStore([]byte("secret")))

	testCases := []struct {
		name         string
		payload      interface{}
		expextedCode int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"email":    u.Email,
				"password": u.Password,
			},
			expextedCode: http.StatusOK,
		},
		{
			name:         "invalid payload",
			payload:      "invalid",
			expextedCode: http.StatusBadRequest,
		},
		{
			name: "invalid Email",
			payload: map[string]string{
				"email":    "string",
				"password": u.Password,
			},
			expextedCode: http.StatusUnauthorized,
		},
		{
			name: "invalid Password",
			payload: map[string]string{
				"email":    u.Email,
				"password": "pasword",
			},
			expextedCode: http.StatusUnauthorized,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()

			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req := httptest.NewRequest(http.MethodPost, "/sessions", b)

			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expextedCode, rec.Code)

		})
	}
}
