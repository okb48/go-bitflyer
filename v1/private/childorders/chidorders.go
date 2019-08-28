package childorders

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/go-numb/go-bitflyer/v1/time"
	"github.com/go-numb/go-bitflyer/v1/types"
)

type Request struct {
	ProductCode            types.ProductCode `json:"product_code" url:"product_code"`
	Pagination             types.Pagination  `json:",inline"`
	ChildOrderState        string            `json:"child_order_state,omitempty"`
	ChildOrderID           string            `json:"child_order_id,omitempty"`
	ChildOrderAcceptanceID string            `json:"child_order_acceptance_id,omitempty"`
	ParentOrderID          string            `json:"parent_order_id,omitempty"`
}

type Response []ChildOrder

type ChildOrder struct {
	ID                     int               `json:"id"`
	ChildOrderID           string            `json:"child_order_id"`
	ProductCode            types.ProductCode `json:"product_code"`
	Side                   string            `json:"side"`
	ChildOrderType         string            `json:"child_order_type"`
	Price                  float64           `json:"price"`
	AveragePrice           float64           `json:"average_price"`
	Size                   float64           `json:"size"`
	ChildOrderState        string            `json:"child_order_state"`
	ExpireDate             time.BitflyerTime `json:"expire_date"`
	ChildOrderDate         time.BitflyerTime `json:"child_order_date"`
	ChildOrderAcceptanceID string            `json:"child_order_acceptance_id"`
	OutstandingSize        float64           `json:"outstanding_size"`
	CancelSize             float64           `json:"cancel_size"`
	ExecutedSize           float64           `json:"executed_size"`
	TotalCommission        float64           `json:"total_commission"`
}

const (
	APIPath string = "/v1/me/getchildorders"
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

	if req.ChildOrderState != "" {
		q += fmt.Sprintf("&child_order_state=%s", req.ChildOrderState)
	}

	if req.ChildOrderAcceptanceID != "" {
		q += fmt.Sprintf("&child_order_acceptance_id=%s", req.ChildOrderAcceptanceID)
	}

	return q
}

func (req *Request) Payload() []byte {
	return nil
}
