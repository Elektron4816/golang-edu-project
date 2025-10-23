package apiserver

type Config struct {
	BindAddress string `toml:"bind_address"`
	LogLevel    string `toml:"log_level"`
	databaseUrl string `toml:"database_url"`
}

func NewConfig() *Config {
	return &Config{
		BindAddress: ":8080",
		LogLevel:    "debug",
		// databaseUrl: "postgres://mac:4816@localhost/education_golang_project?sslmode=disable",
	}
}
