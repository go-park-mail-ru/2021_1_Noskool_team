package repository

import (
	"2021_1_Noskool_team/internal/microservices/auth/models"
	"errors"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/sirupsen/logrus"
)

type SessionsRepository struct {
	con redis.Conn
}

func NewSessionRepository(redisURL string) *SessionsRepository {
	rep := &SessionsRepository{}
	var err error
	logrus.Info(redisURL)
	rep.con, err = redis.DialURL(redisURL)
	logrus.Error(err)
	return rep
}

func (sessionRep *SessionsRepository) CreateSession(session *models.Sessions) (*models.Sessions, error) {
	result, err := redis.String(sessionRep.con.Do("SET", session.UserID, session.UserID,
		"EX", session.Expiration))
	fmt.Println(result)
	if result != "OK" {
		return session, errors.New("status not OK")
	}
	return session, err
}

func (sessionRep *SessionsRepository) CheckSession(session *models.Sessions) (*models.Sessions, error) {
	result, err := redis.String(sessionRep.con.Do("GET", session.UserID))
	fmt.Println(result)
	return session, err
}

func (sessionRep *SessionsRepository) DeleteSession(session *models.Sessions) error {
	result, err := redis.Int(sessionRep.con.Do("DEL", session.UserID))
	fmt.Println(result)
	return err
}
