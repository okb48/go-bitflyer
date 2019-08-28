package health

import (
	"net/http"

	"github.com/go-numb/go-bitflyer/v1/public/markets"
	"github.com/google/go-querystring/query"
)

type Request struct {
	ProductCode markets.ProductCode `json:"product_code" url:"product_code"`
}

type Response struct {
	Status Status `json:"status"`
}

type Status string

const (
	APIPath string = "/v1/gethealth"

	Normal    Status = "NORMAL"
	Busy      Status = "BUSY"
	VeryBusy  Status = "VERY BUSY"
	SuperBusy Status = "SUPER BUSY"
	Stop      Status = "STOP"
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
