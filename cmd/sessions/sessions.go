package main

import (
	"2021_1_Noskool_team/configs"
	grpcSerc "2021_1_Noskool_team/internal/microservices/auth/delivery/grpc/server"
	"2021_1_Noskool_team/internal/microservices/auth/repository"
	sesUsecase "2021_1_Noskool_team/internal/microservices/auth/usecase"
	"time"

	"github.com/BurntSushi/toml"

	"github.com/gomodule/redigo/redis"
	"github.com/sirupsen/logrus"
)

const (
	configPath = "configs/config.toml"
)

func main() {
	time.Sleep(20 * time.Second)
	config := configs.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		logrus.Error(err)
	}

	redisPool := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			pool, err := redis.DialURL(config.SessionRedisStore)
			if err != nil {
				logrus.Error(err)
			}
			return pool, nil
		},
	}

	sessionRep := repository.NewSessionRepository(redisPool)

	sessionsUsecase := sesUsecase.NewSessionsUsecase(sessionRep)
	grpcSerc.StartSessionsGRPCServer(&sessionsUsecase, config.SessionMicroserviceAddr)
}
