package positions

import (
	"net/http"

	"github.com/go-numb/go-bitflyer/v1/time"
	"github.com/go-numb/go-bitflyer/v1/types"
	"github.com/google/go-querystring/query"
)

type Request struct {
	ProductCode types.ProductCode `json:"product_code" url:"product_code"`
}

type Position struct {
	ProductCode         types.ProductCode `json:"product_code"`
	Side                string            `json:"side"`
	Price               float64           `json:"price"`
	Size                float64           `json:"size"`
	Commission          float64           `json:"commission"`
	SwapPointAccumulate float64           `json:"swap_point_accumulate"`
	RequireCollateral   float64           `json:"require_collateral"`
	OpenDate            time.BitflyerTime `json:"open_date"`
	Leverage            float64           `json:"leverage"`
	Pnl                 float64           `json:"pnl"`
	Sfd                 float64           `json:"sfd"`
}

type Response []Position

const (
	APIPath string = "/v1/me/getpositions"
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
