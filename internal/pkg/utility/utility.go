package utility

import (
	"database/sql"
	"github.com/sirupsen/logrus"
)

func CreatePostgresConnection(dbSettings string) (*sql.DB, error) {
	logrus.Info(dbSettings)
	dbCon, err := sql.Open("postgres",
		dbSettings,
	)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	err = dbCon.Ping()
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return dbCon, nil
}
