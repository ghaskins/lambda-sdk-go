package lambda

import (
	"encoding/json"
	"github.com/manetu/lambda-sdk-go/internal"
)

type Headers map[string]string
type Params map[string]any

type Request struct {
	Headers  Headers `json:"headers"`
	PathInfo string  `json:"path-info"`
	Params   Params  `json:"params"`
}

type Response struct {
	Status  int     `json:"status"`
	Headers Headers `json:"headers"`
	Body    any     `json:"body"`
}

type Lambda interface {
	Handler(Request) Response
}

type context struct {
	lambda Lambda
}

func (c context) Handler(request string) string {
	var req Request
	err := json.Unmarshal([]byte(request), &req)
	if err != nil {
		v, err := json.Marshal(&Response{Status: 500})
		if err != nil {
			panic(err)
		}
		return string(v)
	}

	resp := c.lambda.Handler(req)

	v, err := json.Marshal(resp)
	if err != nil {
		panic(err)
	}

	return string(v)
}

func Init(lambda Lambda) {
	internal.SetLambda(&context{lambda: lambda})
}
