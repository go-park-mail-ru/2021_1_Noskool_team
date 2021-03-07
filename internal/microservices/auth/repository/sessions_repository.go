package repository

import (
	"2021_1_Noskool_team/internal/microservices/auth/models"
	"errors"
	"fmt"
	"github.com/gomodule/redigo/redis"
)

type SessionsRepository struct {
	redisPool *redis.Pool
}

func NewSessionRepository(conn *redis.Pool) *SessionsRepository {
	return &SessionsRepository{
		redisPool: conn,
	}
}

func (sessionRep *SessionsRepository) CreateSession(session *models.Sessions) (*models.Sessions, error) {
	con := sessionRep.redisPool.Get()
	defer con.Close()
	result, err := redis.String(con.Do("SET", session.UserID, session.UserID,
		"EX", session.Expiration))
	fmt.Println(result)
	if result != "OK" {
		return session, errors.New("status not OK")
	}
	return session, err
}

func (sessionRep *SessionsRepository) CheckSession(session *models.Sessions) (*models.Sessions, error) {
	con := sessionRep.redisPool.Get()
	defer con.Close()
	result, err := redis.String(con.Do("GET", session.UserID))
	fmt.Println(result)
	return session, err
}

func (sessionRep *SessionsRepository) DeleteSession(session *models.Sessions) error {
	con := sessionRep.redisPool.Get()
	defer con.Close()
	result, err := redis.Int(con.Do("DEL", session.UserID))
	fmt.Println(result)
	return err
}
