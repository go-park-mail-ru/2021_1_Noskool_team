package models

//easyjson:json
type Response struct {
	Body interface{} `json:"body"`
}

func test() {
	r := Response{}
	r.MarshalJSON()
}