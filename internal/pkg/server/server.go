package server

import (
	"2021_1_Noskool_team/configs"
	"2021_1_Noskool_team/internal/app/musicians/models"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Server struct {
	handler Handler
	config  *configs.Config
}

func NewServer(config *configs.Config, handler Handler) (*Server, error) {
	serv := &Server{
		config:  config,
		handler: handler,
	}

	return serv, nil
}

func Start(config *configs.Config, handler Handler) error {
	serv, err := NewServer(config, handler)
	if err != nil {
		logrus.Error(err)
		return err
	}
	logrus.Info(serv.config.MusicServerAddr)
	return http.ListenAndServe(serv.config.MusicServerAddr, serv.handler)
}

func FailedResponse() []byte {
	response := models.FailedResponse{}
	response.ResultStatus = "failed"
	resp, _ := json.Marshal(response)
	return resp
}
