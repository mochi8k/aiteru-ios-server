package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
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
	Post(url string, queries url.Values, body io.Reader) (APIStatus, interface{})
	Get(url string, queries url.Values, body io.Reader) (APIStatus, interface{})
	Put(url string, queries url.Values, body io.Reader) (APIStatus, interface{})
	Delete(url string, queries url.Values, body io.Reader) (APIStatus, interface{})
	Patch(url string, queries url.Values, body io.Reader) (APIStatus, interface{})
	Options(url string, queries url.Values, body io.Reader) (APIStatus, interface{})
}

type APIResourceBase struct{}

func (APIResourceBase) Post(url string, queries url.Values, body io.Reader) (APIStatus, interface{}) {
	return FailByCode(http.StatusMethodNotAllowed), nil
}

func (APIResourceBase) Get(url string, queries url.Values, body io.Reader) (APIStatus, interface{}) {
	return FailByCode(http.StatusMethodNotAllowed), nil
}

func (APIResourceBase) Put(url string, queries url.Values, body io.Reader) (APIStatus, interface{}) {
	return FailByCode(http.StatusMethodNotAllowed), nil
}
func (APIResourceBase) Delete(url string, queries url.Values, body io.Reader) (APIStatus, interface{}) {
	return FailByCode(http.StatusMethodNotAllowed), nil
}
func (APIResourceBase) Patch(url string, queries url.Values, body io.Reader) (APIStatus, interface{}) {
	return FailByCode(http.StatusMethodNotAllowed), nil
}

func (APIResourceBase) Options(url string, queries url.Values, body io.Reader) (APIStatus, interface{}) {
	return FailByCode(http.StatusMethodNotAllowed), nil
}

func Success(code int) APIStatus {
	return APIStatus{isSuccess: true, code: code, message: ""}
}

func Fail(code int, message string) APIStatus {
	return APIStatus{isSuccess: false, code: code, message: message}
}

func FailByCode(code int) APIStatus {
	return APIStatus{isSuccess: false, code: code, message: strconv.Itoa(code) + " " + http.StatusText(code)}
}

func APIResourceHandler(apiResource APIResource) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		b := bytes.NewBuffer(make([]byte, 0))
		reader := io.TeeReader(r.Body, b)

		r.Body = ioutil.NopCloser(b)
		defer r.Body.Close()

		r.ParseForm()

		var status APIStatus
		var response interface{}

		fmt.Printf("%s: %s\n", r.Method, r.URL.Path)
		fmt.Printf("Queries: %v\n", r.Form)
		fmt.Printf("Body: %v\n", reader)

		switch r.Method {
		case post:
			status, response = apiResource.Post(r.URL.Path, r.Form, reader)
		case get:
			status, response = apiResource.Get(r.URL.Path, r.Form, reader)
		case put:
			status, response = apiResource.Put(r.URL.Path, r.Form, reader)
		case delete:
			status, response = apiResource.Delete(r.URL.Path, r.Form, reader)
		case patch:
			status, response = apiResource.Patch(r.URL.Path, r.Form, reader)
		case options:
			status, response = apiResource.Options(r.URL.Path, r.Form, reader)

		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status.code)

		responseBytes, _ := json.Marshal(response)
		w.Write(responseBytes)
	}
}
