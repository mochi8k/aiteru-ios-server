package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/mochi8k/aiteru-ios-server/app/stores"
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
	return func(w http.ResponseWriter, req *http.Request) {
		accessToken := req.Header.Get("Authorization")
		session := stores.GetSession(accessToken)
		fmt.Println(session)
		if session == nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		b := bytes.NewBuffer(make([]byte, 0))
		reader := io.TeeReader(req.Body, b)

		req.Body = ioutil.NopCloser(b)
		defer req.Body.Close()

		req.ParseForm()

		var status APIStatus
		var response interface{}

		fmt.Printf("%s: %s\n", req.Method, req.URL.Path)
		fmt.Printf("Queries: %v\n", req.Form)

		switch req.Method {
		case post:
			status, response = apiResource.Post(req.URL.Path, req.Form, reader)
		case get:
			status, response = apiResource.Get(req.URL.Path, req.Form, reader)
		case put:
			status, response = apiResource.Put(req.URL.Path, req.Form, reader)
		case delete:
			status, response = apiResource.Delete(req.URL.Path, req.Form, reader)
		case patch:
			status, response = apiResource.Patch(req.URL.Path, req.Form, reader)
		case options:
			status, response = apiResource.Options(req.URL.Path, req.Form, reader)

		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status.code)

		responseBytes, _ := json.Marshal(response)
		w.Write(responseBytes)
	}
}
