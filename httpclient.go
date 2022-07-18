package tggo

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type decoder interface {
	decode([]byte) error
}

//lint:ignore U1000 Ignore unused function temporarily for debugging
type httpClient struct {
	endpoint string
	method   string
	params   url.Values
	client   *http.Client
}

func newHttpClient(endpoint string) *httpClient {
	return &httpClient{
		endpoint: endpoint,
		client:   http.DefaultClient,
		params:   url.Values{},
	}
}

func (r *httpClient) setMethod(method string) {
	r.method = method
}

func (r *httpClient) setParam(key string, value string) {
	r.params.Add(key, value)
}

func (r *httpClient) clear() {
	r.method = ""
	r.params = url.Values{}
}

func (r *httpClient) doPost(d decoder) error {
	defer r.clear()

	requrl := strings.Join([]string{r.endpoint, r.method}, "")
	reqparams := strings.NewReader(r.params.Encode())
	req, err := http.NewRequest("POST", requrl, reqparams)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", c_appheader)
	resp, err := r.client.Do(req)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return d.decode(data)
}
