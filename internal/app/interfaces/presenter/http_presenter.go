package presenter

import (
	"encoding/json"
	"net/http"
)

type httpPresenter struct {
	writer  http.ResponseWriter
	status  int
	encoder *json.Encoder
}

func NewHttpPresenter(res http.ResponseWriter) *httpPresenter {
	return &httpPresenter{
		encoder: json.NewEncoder(res),
		writer:  res,
	}
}

func (res *httpPresenter) Status(status int) *httpPresenter {
	res.status = status
	res.writer.WriteHeader(status)
	return res
}

func (res *httpPresenter) Json(data interface{}) {
	res.writer.Header().Set("Content-Type", "application/json")
	res.encoder.Encode(data)
}

func (res *httpPresenter) Error(err error) {
	res.writer.Header().Set("Content-Type", "application/json")

	response := make(map[string]interface{})
	erro := make(map[string]interface{})
	erro["message"] = err.Error()

	response["code"] = res.status
	response["error"] = erro

	res.encoder.Encode(response)
}

func (res *httpPresenter) Success(message string, data interface{}) {
	res.writer.Header().Set("Content-Type", "application/json")

	response := make(map[string]interface{})
	
	if len(message) > 0 {
		success := make(map[string]interface{})
		success["message"] = message
		response["success"] = success
	}

	response["code"] = res.status
	response["data"] = data

	res.encoder.Encode(response)
}

func (res *httpPresenter) Send(data interface{}) {
	res.writer.Header().Set("Content-Type", "text/html")
	res.encoder.Encode(data)
}
