package ticker

import (
	"net/http"

	"github.com/go-numb/go-bitflyer/v1/time"
	"github.com/go-numb/go-bitflyer/v1/types"
	"github.com/google/go-querystring/query"
)

type Request struct {
	ProductCode types.ProductCode `json:"product_code" url:"product_code"`
}

type Response struct {
	ProductCode types.ProductCode `json:"product_code"`

	Timestamp       time.BitflyerTime `json:"timestamp"`
	TickID          int               `json:"tick_id"`
	BestBid         float64           `json:"best_bid"`
	BestAsk         float64           `json:"best_ask"`
	BestBidSize     float64           `json:"best_bid_size"`
	BestAskSize     float64           `json:"best_ask_size"`
	TotalBidDepth   float64           `json:"total_bid_depth"`
	TotalAskDepth   float64           `json:"total_ask_depth"`
	LTP             float64           `json:"ltp"`
	Volume          float64           `json:"volume"`
	VolumeByProduct float64           `json:"volume_by_product"`
}

const (
	APIPath string = "/v1/getticker"
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
