package repository

import (
	"2021_1_Noskool_team/internal/microservices/auth"
	"2021_1_Noskool_team/internal/microservices/auth/models"
	"errors"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/gomodule/redigo/redis"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite
	redisServer *miniredis.Miniredis
	sessions    auth.Repository
}

func (s *Suite) SetupSuite() {
	var err error
	s.redisServer, err = miniredis.Run()
	require.NoError(s.T(), err)

	addr := s.redisServer.Addr()
	redisConn := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", addr)
		},
	}

	s.sessions = NewSessionRepository(redisConn)
}

func (s *Suite) AfterTest(_, _ string) {
	s.SetupSuite()
}

func (s *Suite) TearDownSuite() {
	s.redisServer.Close()
}

func TestSessions(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) TestCreateSession() {
	session := &models.Sessions{
		UserID:     "1",
		Expiration: 86400,
	}

	session, err := s.sessions.CreateSession(session)
	require.NoError(s.T(), err)

	value, err := s.redisServer.Get(session.Hash)
	require.NoError(s.T(), err)
	require.Equal(s.T(), value, session.UserID)

	s.redisServer.FastForward(time.Second * 100000)

	_, err = s.redisServer.Get(session.UserID)
	require.Equal(s.T(), err, errors.New("ERR no such key"))

	s.redisServer.Close()
}

func (s *Suite) TestCheckSession() {
	sessionExpected := &models.Sessions{
		UserID:     "1",
		Hash:       "some hash",
		Expiration: 86400,
	}

	session, err := s.sessions.CreateSession(sessionExpected) //nolint
	require.NoError(s.T(), err)

	value, err := s.redisServer.Get(sessionExpected.Hash)
	require.NoError(s.T(), err)
	require.Equal(s.T(), value, sessionExpected.UserID)

	session, err = s.sessions.CheckSession(sessionExpected)
	require.NoError(s.T(), err)
	require.Equal(s.T(), session, sessionExpected)

	s.redisServer.FastForward(time.Second * 100000)

	_, err = s.sessions.CheckSession(sessionExpected)
	require.Equal(s.T(), err, errors.New("redigo: nil returned"))

	s.redisServer.Close()
}

func (s *Suite) TestDeleteSession() {
	sessionExpected := &models.Sessions{
		UserID:     "1",
		Hash:       "some hash",
		Expiration: 86400,
	}

	_, err := s.sessions.CreateSession(sessionExpected)
	require.NoError(s.T(), err)

	value, err := s.redisServer.Get(sessionExpected.Hash)
	require.NoError(s.T(), err)
	require.Equal(s.T(), value, sessionExpected.UserID)

	err = s.sessions.DeleteSession(sessionExpected)
	require.NoError(s.T(), err)

	_, err = s.redisServer.Get(sessionExpected.Hash)
	require.Equal(s.T(), err, nil)

	s.redisServer.Close()
}
