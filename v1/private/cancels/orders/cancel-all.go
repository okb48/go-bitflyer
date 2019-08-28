package orders

import (
	"encoding/json"
	"net/http"

	"github.com/go-numb/go-bitflyer/v1/types"
	"github.com/google/go-querystring/query"
)

type Request struct {
	ProductCode string `json:"product_code" url:"product_code"`
}

type Response struct {
	ProductCode             types.ProductCode `json:"product_code"`
	ChildOrderId            string            `json:"child_order_id"`
	ChildOrderAcceptanceId  string            `json:"child_order_acceptance_id"`
	ParentOrderId           string            `json:"parent_order_id"`
	ParentOrderAcceptanceId string            `json:"parent_order_acceptance_id"`
}

const (
	APIPath string = "/v1/me/cancelallchildorders"
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
