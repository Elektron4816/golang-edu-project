package apiserver

import (
	"database/sql"
	"net/http"

	"github.com/golang-edu-project/internal/app/store/sqlstore"
	"github.com/gorilla/sessions"
)

func Start(config *Config) error {

	db, err := newDb(config.databaseUrl)

	if err != nil {
		return err
	}

	defer db.Close()

	store := sqlstore.New(db)
	sessionStore := sessions.NewCookieStore([]byte(config.SessionKey))

	s := newServer(store, sessionStore)

	return http.ListenAndServe(config.BindAddress, s)
}

func newDb(databaseUrl string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseUrl)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
