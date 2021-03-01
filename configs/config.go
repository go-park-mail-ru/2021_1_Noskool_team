package configs

type Config struct {
	MusicServerAddr         string `toml:"music_server_addr"`
	SessionMicroserviceAddr string `toml:"session_microservice_addr"`
	SessionRedisStore       string `toml:"session_redis_url"`
	LogLevel                string `toml:"log_level"`
	MusicPostgresBD         string `toml:"music_bd"`
}

func NewConfig() *Config {
	return &Config{
		MusicServerAddr:         ":8080",
		SessionMicroserviceAddr: ":8081",
		SessionRedisStore:       "redis://user:@localhost:6379/0",
		LogLevel:                "debug",
	}
}
