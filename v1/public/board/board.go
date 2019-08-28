package board

import (
	"fmt"
	"net/http"

	"github.com/go-numb/go-bitflyer/v1/public/markets"
	"github.com/google/go-querystring/query"
)

type Book struct {
	Price float64 `json:"price"`
	Size  float64 `json:"size"`
}

type Request struct {
	ProductCode markets.ProductCode `json:"product_code" url:"product_code"`
}

type Response struct {
	MidPrice float64 `json:"mid_price"`
	Bids     []Book  `json:"bids"`
	Asks     []Book  `json:"asks"`
}

const (
	APIPath string = "/v1/getboard"
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

func (b Book) String() string {
	var s string
	if b.Size > 100 {
		s = "💯"
	}

	return fmt.Sprintf("%g x %g%s", b.Price, b.Size, s)
}
