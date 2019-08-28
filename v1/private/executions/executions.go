package executions

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/go-numb/go-bitflyer/v1/time"
	"github.com/go-numb/go-bitflyer/v1/types"
)

type Request struct {
	ProductCode types.ProductCode `json:"product_code" url:"product_code"`
	Pagination  types.Pagination  `json:",inline"`
}

type Response []Execution

type Execution struct {
	ID                     int               `json:"id"`
	ChildOrderID           string            `json:"child_order_id"`
	Side                   string            `json:"side"`
	Price                  float64           `json:"price"`
	Size                   float64           `json:"size"`
	Commission             float64           `json:"commission"`
	ExecDate               time.BitflyerTime `json:"exec_date"`
	ChildOrderAcceptanceID string            `json:"child_order_acceptance_id"`
}

const (
	APIPath string = "/v1/me/getexecutions"
)

func (req *Request) Method() string {
	return http.MethodGet
}

func (req *Request) Query() string {
	// values, _ := query.Values(req)
	q := "product_code=" + string(req.ProductCode)
	if !reflect.DeepEqual(req.Pagination, types.Pagination{}) {
		if req.Pagination.Count != 0 {
			q += fmt.Sprintf("&count=%d", req.Pagination.Count)
		}
		if req.Pagination.Before != 0 {
			q += fmt.Sprintf("&before=%d", req.Pagination.Before)
		}
		if req.Pagination.After != 0 {
			q += fmt.Sprintf("&after=%d", req.Pagination.After)
		}
	}

	return q
}

func (req *Request) Payload() []byte {
	return nil
}
