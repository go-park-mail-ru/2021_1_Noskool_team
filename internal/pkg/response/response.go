package response

import (
	"2021_1_Noskool_team/internal/app/musicians/models"
	commonModels "2021_1_Noskool_team/internal/models"
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

func FailedResponse(w http.ResponseWriter, code int) []byte {
	w.WriteHeader(code)
	response := models.FailedResponse{}
	response.ResultStatus = "failed"
	resp, _ := response.MarshalJSON()
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
	_, _ = w.Write(body)
}

func SendCorrectResponse(w http.ResponseWriter, data interface{}, HTTPStatus int,
	marshalingFunc func(data interface{}) ([]byte, error)) {
	//body, err := commonModels.Response{Body: data}.MarshalJSON()
	body, err := marshalingFunc(data)
	if err != nil {
		SendErrorResponse(w, &commonModels.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: "Error encoding json",
		})
		return
	}

	w.WriteHeader(HTTPStatus)
	_, _ = w.Write(body)
}

func SendEmptyBody(w http.ResponseWriter, HTTPStatusCode int) {
	w.WriteHeader(HTTPStatusCode)
}
