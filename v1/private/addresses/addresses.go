package addresses

import (
	"net/http"

	"github.com/go-numb/go-bitflyer/v1/types"
	"github.com/google/go-querystring/query"
)

type Request struct{}

type Response []Address

type Address struct {
	Type         DepositType        `json:"type"`
	CurrencyCode types.CurrencyCode `json:"currency_code"`
	Address      string             `json:"address"`
}

type DepositType string

const (
	APIPath string = "/v1/me/getaddresses"
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
