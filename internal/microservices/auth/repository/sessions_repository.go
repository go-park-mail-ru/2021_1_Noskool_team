package repository

import (
	"errors"
	"github.com/gomodule/redigo/redis"
	"testWorkWithAuth/internal/microservices/auth/models"
)

type SessionsRepository struct {
	con redis.Conn
}

func NewSessionRepository(redisURL string) *SessionsRepository {
	rep := &SessionsRepository{}
	rep.con, _ = redis.DialURL(redisURL)
	return rep
}

func (sessionRep *SessionsRepository) CreateSession(session *models.Sessions) (*models.Sessions, error) {
	result, err := redis.String(sessionRep.con.Do("SET", session.UserID, session.UserID,
		"EX", session.Expiration))
	if result != "OK" {
		return session, errors.New("status not OK")
	}
	return session, err
}

func (sessionRep *SessionsRepository) CheckSession(session *models.Sessions) (*models.Sessions, error) {
	userID, err := redis.Int(sessionRep.con.Do("GET", session.UserID))
	session.UserID = userID
	return  session, err
}

func (sessionRep *SessionsRepository) DeleteSession(session *models.Sessions) error {
	_, err := redis.Int(sessionRep.con.Do("DEL", session.UserID))
	return err
}
