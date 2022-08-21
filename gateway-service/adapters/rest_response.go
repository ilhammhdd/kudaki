package adapters

import (
	"encoding/json"
	"net/http"

	"github.com/ilhammhdd/go-toolkit/errorkit"
)

type Response struct {
	HttpCode int
	JSONBody []byte
}

func (res *Response) WriteResponse(w *http.ResponseWriter) {
	(*w).Header().Set("Content-Type", "application/json")
	(*w).WriteHeader(res.HttpCode)
	(*w).Write(res.JSONBody)
}

type ResponseBody struct {
	Errs *[]string   `json:"errors,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

func (r *ResponseBody) Parse() []byte {
	resBodyJSON, err := json.Marshal(r)
	errorkit.ErrorHandled(err)

	return resBodyJSON
}

func NewResponse(httpCode int, resBody *ResponseBody) *Response {
	resBodyJSON, err := json.Marshal(resBody)
	errorkit.ErrorHandled(err)

	return &Response{
		HttpCode: httpCode,
		JSONBody: resBodyJSON}
}

type DataMap map[string]interface{}
