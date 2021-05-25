package utility

import (
	sessionModels "2021_1_Noskool_team/internal/microservices/auth/models"
	"2021_1_Noskool_team/internal/models"
	commonModels "2021_1_Noskool_team/internal/models"
	"2021_1_Noskool_team/internal/pkg/response"
	"crypto/sha256"
	"database/sql"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	_ "github.com/lib/pq" //goland:noinspection
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
	_, _ = io.Copy(f, file)
	return &newFileName, nil
}

func ParsePagination(r *http.Request) *models.Pagination {
	pagination := &models.Pagination{}
	limit := r.URL.Query().Get("limit")
	limitInt, err := strconv.Atoi(limit)
	if err == nil {
		pagination.Limit = limitInt
	}
	offset := r.URL.Query().Get("offset")
	offsetInt, err := strconv.Atoi(offset)
	if err == nil {
		pagination.Offset = offsetInt
	}
	return pagination
}

var characterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func RandomString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = characterRunes[rand.Intn(len(characterRunes))]
	}
	return string(b)
}

func CreateCSRFToken(secret string) string {
	b := secret + RandomString(5)
	return fmt.Sprintf("%x", sha256.Sum256([]byte(b)))
}

func CreatePlaylistUID(playlistID string) string {
	return CreateCSRFToken(playlistID)[:10]
}

func CheckUserID(w http.ResponseWriter, r *http.Request, logger *logrus.Logger) (int, error) {
	session, ok := r.Context().Value("user_id").(sessionModels.Result)
	if !ok {
		logger.Error("Не получилось достать из конекста")
		response.SendErrorResponse(w, &commonModels.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "Not correct user id",
		})
		return 0, errors.New("Not correct user id")
	}
	userID, err := strconv.Atoi(session.ID)
	if err != nil {
		logger.Error(err)
		response.SendErrorResponse(w, &commonModels.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: "Error converting userID to int",
		})
		return 0, err
	}
	return userID, nil
}
