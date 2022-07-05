package telego

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

//lint:ignore U1000 Ignore unused function temporarily for debugging
type httpClient struct {
	header   http.Header
	endpoint string
	method   string
	params   url.Values
	client   *http.Client
}

func newHttpClient(endpoint string, header http.Header) *httpClient {
	return &httpClient{
		header:   header,
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

func (r *httpClient) do() ([]byte, error) {
	defer r.clear()

	requrl := strings.Join([]string{r.endpoint, r.method}, "")
	reqparams := strings.NewReader(r.params.Encode())
	req, err := http.NewRequest("POST", requrl, reqparams)
	if err != nil {
		return nil, err
	}

	req.Header = r.header
	resp, err := r.client.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}
