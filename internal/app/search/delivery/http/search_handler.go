package http

import (
	"2021_1_Noskool_team/configs"
	"2021_1_Noskool_team/internal/app/middleware"
	"2021_1_Noskool_team/internal/app/search"
	"2021_1_Noskool_team/internal/app/search/models"
	"2021_1_Noskool_team/internal/pkg/response"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/microcosm-cc/bluemonday"
	"github.com/sirupsen/logrus"
	"net/http"
)

type SearchHandler struct {
	router        *mux.Router
	searchUsecase search.Usecase
	logger        *logrus.Logger
	sanitiser     *bluemonday.Policy
}

func NewSearchHandler(r *mux.Router, config *configs.Config, usecase search.Usecase,
	sanitizer *bluemonday.Policy) *SearchHandler {
	handler := &SearchHandler{
		router:        r,
		searchUsecase: usecase,
		logger:        logrus.New(),
		sanitiser:     sanitizer,
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
	searchQuery = handler.sanitiser.Sanitize(searchQuery)

	fmt.Println(searchQuery)
	search := handler.searchUsecase.SearchContent(searchQuery)
	response.SendCorrectResponse(w, search, http.StatusOK, models.MarshalSearch)
}
