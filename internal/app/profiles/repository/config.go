package repository

// Config ...
type Config struct {
	DatabaseURL string `toml:"profiles_db_url"`
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{}
}
