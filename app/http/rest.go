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

type apiHeader struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type apiEnvelope struct {
	Header   apiHeader   `json:"header"`
	Response interface{} `json:"response"`
}

func APIRsourceHandler(apiResource APIResource) http.HandleFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		b := bytes.NewBuffer(make([]byte, 0))
		reader := io.TeeReader(r.Body, b)

		r.Body = ioutil.NopCloser(b)
		defer r.Body.Close()

		r.ParseForm()

		var status APIStatus
		var data interface{}

		switch r.Method {
		case post:
			status, data = apiResource.POST(r.URL.Path, r.Form, reader)
		case get:
			status, data = apiResource.GET(r.URL.Path, r.Form, reader)
		case put:
			status, data = apiResource.PUT(r.URL.Path, r.Form, reader)
		case delete:
			status, data = apiResource.DELETE(r.URL.Path, r.Form, reader)
		case patch:
			status, data = apiResource.PATCH(r.URL.Path, r.Form, reader)
		case options:
			status, data = apiResource.OPTIONS(r.URL.Path, r.Form, reader)

		}

		var content []byte
		var e error

		if status.isSuccess {
			content, e = json.Marshal(apiEnvelope{
				Header:   apiHeader{Status: "success"},
				Response: data,
			})
		} else {
			content, e = json.Marshal(apiEnvelope{
				Header: apiHeader{Status: "fail", Message: status.message},
			})
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status.code)
		w.Write(content)

	}
}
