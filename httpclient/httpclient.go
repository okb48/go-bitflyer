package httpclient

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/go-numb/go-bitflyer"
	"github.com/go-numb/go-bitflyer/auth"
	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
)

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary
)

type httpClient struct {
	authConfig *auth.AuthConfig
}

func New() *httpClient {
	return &httpClient{}
}

func (p *httpClient) Auth(conf *auth.AuthConfig) *httpClient {
	p.authConfig = conf
	return p
}

func (p *httpClient) Request(api api.API, req api.Request, result interface{}) (*http.Response, error) {
	u, err := api.ToURL()
	if err != nil {
		return nil, errors.Wrapf(err, "set base url")
	}

	// 引数をurlにセット
	// struct Reqestはインターフェース型
	u.RawQuery = req.Query()
	payload := req.Payload()

	var body io.Reader
	if len(payload) != 0 {
		body = bytes.NewReader(payload)
	}

	r, err := http.NewRequest(req.Method(), u.String(), body)
	if err != nil {
		return nil, errors.Wrapf(err, "create new post, requests from url: %s", u.String())
	}

	// Private configがあれば、sets auth to header
	if p.authConfig != nil {
		header, err := auth.SetAuthHeaders(p.authConfig, api, req)
		if err != nil {
			return nil, errors.Wrap(err, "can not generates auth, or sets header")
		}

		r.Header = *header
	}

	if len(payload) != 0 {
		r.Header.Set("Content-Type", "application/json")
	}

	c := &http.Client{
		Timeout: 10 * time.Second,
	}

	/* Header's at 2019/08/28
	&{Status:200 OK
	StatusCode:200
	Proto:HTTP/2.0
	ProtoMajor:2
	ProtoMinor:0
	Header:map[Cache-Control:[no-cache]
	Content-Security-Policy:[
	default-src
	http:
	https:
	ws:
	wss:
	data:
	'unsafe-inline'
	'unsafe-eval']
	Content-Type:[application/json; charset=utf-8]
	Date:[]
	Expires:[-1]
	Pragma:[no-cache]
	Request-Context:[appId=cid-v1:]
	Server:[Microsoft-IIS/10.0]
	Strict-Transport-Security:[max-age=31536000]
	Vary:[Accept-Encoding]
	X-Content-Type-Options:[nosniff]
	X-Frame-Options:[sameorigin]
	X-Ratelimit-Period:[153]  ********* API Limit残
	X-Ratelimit-Remaining:[494]  ********* API Limit
	X-Ratelimit-Reset:[1566997170] ********* API Limit 復活
	X-Xss-Protection:[1;
	mode=block]]
	ContentLength:-1
	TransferEncoding:[]
	Close:false
	Uncompressed:true
	Trailer:map[]
	Request:
	TLS:}
	*/

	res, err := c.Do(r)
	if err != nil {
		return nil, errors.Wrapf(err, "requests do something to url: %s", u.String())
	}
	defer res.Body.Close()

	// check status code
	if res.StatusCode != 200 {
		return nil, errors.New("response status: " + res.Status)
	}

	err = json.NewDecoder(res.Body).Decode(result)
	if err != nil {
		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, errors.Wrapf(err, "can not read(json/io) data from url: %s", u.String())
		}
		return nil, errors.Wrapf(err, "unmarshal data: %s", string(data))
	}

	return res, nil
}
