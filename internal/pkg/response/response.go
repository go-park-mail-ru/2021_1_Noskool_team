package response

import (
	"2021_1_Noskool_team/internal/app/musicians/models"
	"encoding/json"
	"net/http"
)

func FailedResponse(w http.ResponseWriter, code int) []byte {
	w.WriteHeader(code)
	response := models.FailedResponse{}
	response.ResultStatus = "failed"
	resp, _ := json.Marshal(response)
	return resp
}
