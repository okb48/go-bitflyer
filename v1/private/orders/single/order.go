package single

import (
	"encoding/json"
	"net/http"

	"github.com/go-numb/go-bitflyer/v1/types"
	"github.com/google/go-querystring/query"
)

type Request struct {
	ProductCode    types.ProductCode `json:"product_code"`
	ChildOrderType string            `json:"child_order_type"`
	Side           string            `json:"side"`
	Price          float64           `json:"price"`
	Size           float64           `json:"size"`
	MinuteToExpire int               `json:"minute_to_expire"`
	TimeInForce    string            `json:"time_in_force"`
}

type Response struct {
	ChildOrderAcceptanceId string `json:"child_order_acceptance_id"`
}

const (
	APIPath string = "/v1/me/sendchildorder"
)

func (req *Request) Method() string {
	return http.MethodPost
}

func (req *Request) Query() string {
	values, _ := query.Values(req)
	return values.Encode()
}

func (req *Request) Payload() []byte {
	b, err := json.Marshal(*req)
	if err != nil {
		return nil
	}
	return b
}
