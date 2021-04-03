package http

import (
	"2021_1_Noskool_team/configs"
	"2021_1_Noskool_team/internal/app/middleware"
	"2021_1_Noskool_team/internal/app/search"
	"2021_1_Noskool_team/internal/pkg/response"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

type SearchHandler struct {
	router        *mux.Router
	searchUsecase search.Usecase
	logger        *logrus.Logger
}

func NewSearchHandler(r *mux.Router, config *configs.Config, usecase search.Usecase) *SearchHandler {
	handler := &SearchHandler{
		router:        r,
		searchUsecase: usecase,
		logger:        logrus.New(),
	}
	err := ConfigLogger(handler, config)
	if err != nil {
		logrus.Error(err)
	}
	handler.router.Use(middleware.ContentTypeJson)

	handler.router.HandleFunc("/", handler.SearchContent)

	return handler
}

func ConfigLogger(handler *SearchHandler, config *configs.Config) error {
	level, err := logrus.ParseLevel(config.LogLevel)
	if err != nil {
		return err
	}

	handler.logger.SetLevel(level)
	return nil
}

func (handler *SearchHandler) SearchContent(w http.ResponseWriter, r *http.Request) {
	searchQuery := r.URL.Query().Get("search")
	search := handler.searchUsecase.SearchContent(searchQuery)
	response.SendCorrectResponse(w, search, http.StatusOK)
}
