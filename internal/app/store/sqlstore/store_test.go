package sqlstore_test

import (
	"os"
	"testing"
)

var (
	databaseUrl string
)

func TestMain(m *testing.M) {
	databaseUrl = os.Getenv("DATABASE_URL")
	if databaseUrl == "" {
		databaseUrl = "postgres://mac:4816@localhost/rest_api_test?sslmode=disable"
	}

	os.Exit(m.Run())
}
