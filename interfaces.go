package api

import "net/url"

type Client interface {
	APIHost() string
}

type API interface {
	ToURL() (*url.URL, error)
}

type Request interface {
	Method() string
	Query() string
	Payload() []byte
}
