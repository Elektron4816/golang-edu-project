package apiserver

type Config struct {
	BindAddress string `toml:"bind_address"`
	LogLevel    string `toml:"log_level"`
	databaseUrl string `toml:"database_url"`
	SessionKey  string `toml:"session_key"`
}

func NewConfig() *Config {
	return &Config{
		BindAddress: ":3000",
		LogLevel:    "debug",
		databaseUrl: "postgres://agolubev:new_password@localhost/education_golang_project?sslmode=disable",
	}
}
