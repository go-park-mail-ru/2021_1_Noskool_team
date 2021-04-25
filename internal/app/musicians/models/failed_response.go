package models

//easyjson:json
type FailedResponse struct {
	ResultStatus string `json:"status"`
}

func test() {
	r := FailedResponse{}
	r.MarshalJSON()
}
