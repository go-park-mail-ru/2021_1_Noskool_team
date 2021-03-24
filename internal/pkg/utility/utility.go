package utility

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" //goland:noinspection
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"path/filepath"
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

func SaveFile(r *http.Request, formKey string, directory string, fileName string) (*string, error) {
	err := r.ParseMultipartForm(5 * 1024 * 1025)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	file, handler, err := r.FormFile(formKey)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	path, _ := os.Getwd()
	newFileName := fileName + filepath.Ext(handler.Filename)
	filePath := path + directory + newFileName

	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer f.Close()
	io.Copy(f, file)
	return &newFileName, nil
}
