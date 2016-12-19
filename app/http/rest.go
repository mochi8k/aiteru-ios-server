package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/mochi8k/aiteru-ios-server/app/models"
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
	Post(url string, queries url.Values, body io.Reader, session *models.Session) (APIStatus, interface{})
	Get(url string, queries url.Values, body io.Reader, session *models.Session) (APIStatus, interface{})
	Put(url string, queries url.Values, body io.Reader, session *models.Session) (APIStatus, interface{})
	Delete(url string, queries url.Values, body io.Reader, session *models.Session) (APIStatus, interface{})
	Patch(url string, queries url.Values, body io.Reader, session *models.Session) (APIStatus, interface{})
	Options(url string, queries url.Values, body io.Reader, session *models.Session) (APIStatus, interface{})
}

type APIResourceBase struct{}

func (APIResourceBase) Post(url string, queries url.Values, body io.Reader, session *models.Session) (APIStatus, interface{}) {
	return FailByCode(http.StatusMethodNotAllowed), nil
}

func (APIResourceBase) Get(url string, queries url.Values, body io.Reader, session *models.Session) (APIStatus, interface{}) {
	return FailByCode(http.StatusMethodNotAllowed), nil
}

func (APIResourceBase) Put(url string, queries url.Values, body io.Reader, session *models.Session) (APIStatus, interface{}) {
	return FailByCode(http.StatusMethodNotAllowed), nil
}
func (APIResourceBase) Delete(url string, queries url.Values, body io.Reader, session *models.Session) (APIStatus, interface{}) {
	return FailByCode(http.StatusMethodNotAllowed), nil
}
func (APIResourceBase) Patch(url string, queries url.Values, body io.Reader, session *models.Session) (APIStatus, interface{}) {
	return FailByCode(http.StatusMethodNotAllowed), nil
}

func (APIResourceBase) Options(url string, queries url.Values, body io.Reader, session *models.Session) (APIStatus, interface{}) {
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

// TODO: another class
func Auth(url, accessToken string) (*models.Session, bool) {

	// TODO: matching
	if url == "/auth" {
		return nil, false
	}

	session := stores.GetSession(accessToken)

	if session == nil {
		return nil, true
	}

	return session, false
}

func APIResourceHandler(apiResource APIResource) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		session, isUnauth := Auth(req.URL.Path, req.Header.Get("Authorization"))

		if isUnauth {
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
			status, response = apiResource.Post(req.URL.Path, req.Form, reader, session)
		case get:
			status, response = apiResource.Get(req.URL.Path, req.Form, reader, session)
		case put:
			status, response = apiResource.Put(req.URL.Path, req.Form, reader, session)
		case delete:
			status, response = apiResource.Delete(req.URL.Path, req.Form, reader, session)
		case patch:
			status, response = apiResource.Patch(req.URL.Path, req.Form, reader, session)
		case options:
			status, response = apiResource.Options(req.URL.Path, req.Form, reader, session)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status.code)

		responseBytes, _ := json.Marshal(response)
		w.Write(responseBytes)
	}
}
