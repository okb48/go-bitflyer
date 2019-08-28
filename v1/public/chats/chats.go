package chats

import (
	"net/http"

	"github.com/go-numb/go-bitflyer/v1/time"
	"github.com/google/go-querystring/query"
)

type Request struct {
	FromDate time.BitflyerTime `json:"from_date,omitempty"`
}

type Response []Chat

type Chat struct {
	Nickname string            `json:"nickname"`
	Message  string            `json:"message"`
	Date     time.BitflyerTime `json:"date"`
}

const (
	APIPath string = "/v1/getchats"
)

func (req *Request) Method() string {
	return http.MethodGet
}

func (req *Request) Query() string {
	values, _ := query.Values(req)
	return values.Encode()
}

func (req *Request) Payload() []byte {
	return nil
}
