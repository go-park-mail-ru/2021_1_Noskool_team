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
		MusicServerAddr:         ":8080",
		SessionMicroserviceAddr: "127.0.0.1:8081",
		SessionRedisStore:       "redis://user:@localhost:6379/0",
		MusicPostgresBD:         "host=localhost port=5432 dbname=music_service sslmode=disable",
		ProfilesServerAddr:      "8082",
		ProfileDB:               "",
		LogLevel:                "debug",
		FrontendURL:             "some_url",
		MediaFolder:             "./static",
	}
}
