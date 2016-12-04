package http

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	post    = "POST"
	get     = "GET"
	put     = "PUT"
	delete  = "DELETE"
	patch   = "PATCH"
	options = "OPTIONS"
)

type APIStatus struct {
	isSuccess bool
	code      int
	message   string
}

type APIResource interface {
	POST(url string, queries url.Values, body io.Reader) (APIStatus, interface{})
	GET(url string, queries url.Values, body io.Reader) (APIStatus, interface{})
	PUT(url string, queries url.Values, body io.Reader) (APIStatus, interface{})
	DELETE(url string, queries url.Values, body io.Reader) (APIStatus, interface{})
	PATCH(url string, queries url.Values, body io.Reader) (APIStatus, interface{})
	OPTIONS(url string, queries url.Values, body io.Reader) (APIStatus, interface{})
}

func APIRsourceHandler(apiResource APIResource) http.HandleFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		b := bytes.NewBuffer(make([]byte, 0))
		reader := io.TeeReader(r.Body, b)

		r.Body = ioutil.NopCloser(b)
		defer r.Body.Close()

		r.ParseForm()

		var status APIStatus
		var response interface{}

		switch r.Method {
		case post:
			status, response = apiResource.POST(r.URL.Path, r.Form, reader)
		case get:
			status, response = apiResource.GET(r.URL.Path, r.Form, reader)
		case put:
			status, response = apiResource.PUT(r.URL.Path, r.Form, reader)
		case delete:
			status, response = apiResource.DELETE(r.URL.Path, r.Form, reader)
		case patch:
			status, response = apiResource.PATCH(r.URL.Path, r.Form, reader)
		case options:
			status, response = apiResource.OPTIONS(r.URL.Path, r.Form, reader)

		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status.code)

		responseBytes, _ := json.Marshal(response)
		w.Write(responseBytes)
	}
}
