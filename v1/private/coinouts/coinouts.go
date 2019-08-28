package coinouts

import (
	"net/http"

	"github.com/go-numb/go-bitflyer/v1/time"
	"github.com/go-numb/go-bitflyer/v1/types"
	"github.com/google/go-querystring/query"
)

type Request struct {
	Pagination types.Pagination `json:",inline"`
}

type Response []Coinout

type Coinout struct {
	ID            int                `json:"id"`
	OrderID       string             `json:"order_id"`
	CurrencyCode  types.CurrencyCode `json:"currency_code"`
	Amount        float64            `json:"amount"`
	Address       string             `json:"address"`
	TxHash        string             `json:"tx_hash"`
	Fee           float64            `json:"fee"`
	AdditionalFee float64            `json:"additional_fee"`
	Status        Status             `json:"status"` // If the Bitcoin deposit is being processed, it will be listed as "PENDING". If the deposit has been completed, it will be listed as "COMPLETED"
	EventDate     time.BitflyerTime  `json:"event_date"`
}

type Status string

const (
	APIPath string = "/v1/me/getcoinouts"

	PENDING   Status = "PENDING"
	COMPLETED Status = "COMPLETED"
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
