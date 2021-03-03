package configs

import "2021_1_Noskool_team/internal/app/profiles/repository"

// Config ...
type Config struct {
	MusicServerAddr         string `toml:"music_server_addr"`
	SessionMicroserviceAddr string `toml:"session_microservice_addr"`
	SessionRedisStore       string `toml:"session_redis_url"`

	ProfilesServerAddr string             `toml:"profiles_server_addr"`
	ProfileDB          *repository.Config `toml:"ProfileDB"`

	LogLevel string `toml:"log_level"`
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		MusicServerAddr:         ":8080",
		SessionMicroserviceAddr: "127.0.0.1:8081",
		SessionRedisStore:       "redis://user:@localhost:6379/0",

		ProfilesServerAddr: "8082",
		ProfileDB:          repository.NewConfig(),

		LogLevel: "debug",
	}
}
