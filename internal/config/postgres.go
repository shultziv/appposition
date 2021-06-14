package config

type Postgres struct {
	Url string `env:"POSTGRES_URL"`
}
