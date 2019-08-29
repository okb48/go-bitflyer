package v1

import (
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

type API struct {
	url string
}

func NewAPI(c *Client, apiPath string) *API {
	return &API{
		url: c.APIHost() + apiPath,
	}
}

func (api *API) ToURL() (*url.URL, error) {
	u, err := url.ParseRequestURI(api.url)
	if err != nil {
		return nil, errors.Wrapf(err, "can not parse url: %s", api.url)
	}
	return u, nil
}

// 基本的には5分毎リセット
type APIHeaders struct {
	Public  Limit
	Private Limit
}

func (p *APIHeaders) IsCache(h http.Header) bool {
	isCache := h.Get("Pragma")
	if isCache != "no-cache" {
		// on キャッシュ
		return true
	}
	// does not キャッシュ
	return false
}

type Limit struct {
	Period int       // Period is リセットまでの秒数
	Remain int       // Remain is 残Requests
	Reset  time.Time // Reset Remainの詳細時間(sec未満なし)
}

// FromHeader X-xxxからLimitを取得
func (p *Limit) FromHeader(h http.Header) {
	period := h.Get("X-Ratelimit-Period") // リセットまでの残秒数
	if period != "" {
		p.Period, _ = strconv.Atoi(period)
	}
	remain := h.Get("X-Ratelimit-Remaining") // 残回数
	if remain != "" {
		p.Remain, _ = strconv.Atoi(remain)
	}
	t := h.Get("X-Ratelimit-Reset") // リセットUTC時間(sec未満なし)
	if t != "" {
		reset, _ := strconv.ParseInt(t, 10, 64)
		p.toTime(reset)
	}
}

// int64 to time.Time
func (p *Limit) toTime(t int64) {
	p.Reset = time.Unix(t, 10)
}
