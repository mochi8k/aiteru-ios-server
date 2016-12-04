package http

import (
	"io"
	"net/url"
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
}
