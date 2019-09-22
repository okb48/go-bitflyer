package markets

import (
	"fmt"
	"net/http"

	"github.com/go-numb/go-bitflyer/v1/types"
)

type Request struct{}

type Response []Market

type Market struct {
	ProductCode types.ProductCode `json:"product_code"`
	Alias       types.ProductCode `json:"alias"`
}

type ProductCode string

const (
	APIPath string = "/v1/getmarkets"
)

func (req *Request) Method() string {
	return http.MethodGet
}

func (req *Request) Query() string {
	return ""
}

func (req *Request) Payload() []byte {
	return nil
}

func (m Market) String() string {
	if m.Alias == "" {
		return string(m.ProductCode)
	}
	return fmt.Sprintf("%s (%s)", m.ProductCode, m.Alias)
}
