package executions

import (
	"net/http"

	"github.com/go-numb/go-bitflyer/v1/time"
	"github.com/go-numb/go-bitflyer/v1/types"
	"github.com/google/go-querystring/query"
)

type Request struct {
	ProductCode types.ProductCode `json:"product_code" url:"product_code"`

	Pagination types.Pagination `json:",inline"`
}

type Response []Execution

type Execution struct {
	ID                         int               `json:"id"`
	Side                       string            `json:"side"`
	Price                      float64           `json:"price"`
	Size                       float64           `json:"size"`
	ExecDate                   time.BitflyerTime `json:"exec_date"`
	BuyChildOrderAcceptanceID  string            `json:"buy_child_order_acceptance_id"`
	SellChildOrderAcceptanceID string            `json:"sell_child_order_acceptance_id"`
}

const (
	APIPath string = "/v1/getexecutions"
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
