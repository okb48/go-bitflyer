package sp

import (
	"encoding/json"
	"net/http"

	"github.com/go-numb/go-bitflyer/v1/types"
	"github.com/google/go-querystring/query"
)

type Request struct {
	OrderMethod    string  `json:"order_method"`
	MinuteToExpire int     `json:"minute_to_expire"`
	TimeInForce    string  `json:"time_in_force"`
	Parameters     []Param `json:"parameters"`
}

type Param struct {
	ProductCode   types.ProductCode `json:"product_code"`
	ConditionType string            `json:"condition_type"`
	Side          string            `json:"side"`
	Price         float64           `json:"price"`
	TriggerPrice  float64           `json:"trigger_price"`
	Size          float64           `json:"size"`
	// Offset        int     `json:"offset"`
}

type Response struct {
	ParentOrderAcceptanceId string `json:"parent_order_acceptance_id"`
}

const (
	APIPath string = "/v1/me/sendparentorder"
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
