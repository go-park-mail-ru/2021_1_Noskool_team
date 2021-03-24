package response

import (
	"2021_1_Noskool_team/internal/app/musicians/models"
	commonModels "2021_1_Noskool_team/internal/models"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
)

func FailedResponse(w http.ResponseWriter, code int) []byte {
	w.WriteHeader(code)
	response := models.FailedResponse{}
	response.ResultStatus = "failed"
	resp, _ := json.Marshal(response)
	return resp
}

func SendErrorResponse(w http.ResponseWriter, error *commonModels.HTTPError) {
	logrus.Error(error.Message)
	w.WriteHeader(error.Code)
	body, err := json.Marshal(error)
	if err != nil {
		SendErrorResponse(w, &commonModels.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: "Error encoding json",
		})
		return
	}
	w.Write(body)
}

func SendCorrectResponse(w http.ResponseWriter, data interface{}, HTTPStatus int) {
	body, err := json.Marshal(data)
	if err != nil {
		SendErrorResponse(w, &commonModels.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: "Error encoding json",
		})
		return
	}

	w.WriteHeader(HTTPStatus)
	w.Write(body)
}

func SendEmptyBody(w http.ResponseWriter, HTTPStatusCode int) {
	w.WriteHeader(HTTPStatusCode)
}
