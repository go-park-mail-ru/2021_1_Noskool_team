package configs

// Config ...
type Config struct {
	MusicServerAddr         string `toml:"music_server_addr"`
	SessionMicroserviceAddr string `toml:"session_microservice_addr"`
	SessionRedisStore       string `toml:"session_redis_url"`
	LogLevel                string `toml:"log_level"`
	MusicPostgresBD         string `toml:"music_bd"`
	ProfilesServerAddr      string `toml:"profiles_server_addr"`
	ProfileDB               string `toml:"profiles_db_url"`
	FrontendURL             string `toml:"frontend_url"`
	MediaFolder             string `toml:"media_folder"`
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		MusicServerAddr:         ":8888",
		SessionMicroserviceAddr: "sessions-service:8081",
		SessionRedisStore:       "redis://redis/0",
		MusicPostgresBD:         "host=music-bd port=5432 user=andrewkireev dbname=music_service_docker password=password sslmode=disable",
		ProfilesServerAddr:      ":8082",
		ProfileDB:               "host=music-bd port=5432 user=andrewkireev dbname=music_service_docker password=password sslmode=disable",
		LogLevel:                "debug",
		FrontendURL:             "84.201.189.81:9000",
		MediaFolder:             "./static",
	}
}
