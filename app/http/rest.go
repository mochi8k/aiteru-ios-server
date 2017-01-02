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

	"github.com/julienschmidt/httprouter"
	"github.com/mochi8k/aiteru-ios-server/app/handlers/router"
	"github.com/mochi8k/aiteru-ios-server/app/models"
	"github.com/mochi8k/aiteru-ios-server/app/stores"
)

const (
	post    = "POST"
	get     = "GET"
	put     = "PUT"
	delete  = "DELETE"
	patch   = "PATCH"
	options = "OPTIONS"
)

type Handler func(ps httprouter.Params, queries url.Values, body io.Reader, session *models.Session) (APIStatus, interface{})

type APIStatus struct {
	isSuccess bool
	code      int
	message   string
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
		fmt.Printf("Session does not exist\n url: %s\n accessToken: %s\n", url, accessToken)
		return nil, true
	}

	return session, false
}

func Register(pattern string, requestHandlers map[string]Handler) {
	for method, requestHandler := range requestHandlers {
		switch method {
		case post:
			router.POST(pattern, apiResourceHandler(requestHandler))
		case get:
			router.GET(pattern, apiResourceHandler(requestHandler))
		case put:
			router.PUT(pattern, apiResourceHandler(requestHandler))
		case delete:
			router.DELETE(pattern, apiResourceHandler(requestHandler))
		}
	}
}

func apiResourceHandler(requestHandler Handler) httprouter.Handle {
	return func(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {

		fmt.Printf("RequestURI: %s\n", req.URL.Path)

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

		status, response = requestHandler(ps, req.Form, reader, session)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status.code)

		responseBytes, _ := json.Marshal(response)
		w.Write(responseBytes)
	}
}
