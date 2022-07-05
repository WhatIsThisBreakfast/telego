package telego

import (
	"fmt"
	"net/http"
)

type api struct {
	apiclient *httpClient
}

func newApi(token string, apiendpoint string) *api {
	endpoint := fmt.Sprintf(apiendpoint, token)

	header := make(http.Header)
	header.Set("Content-Type", "application/x-www-form-urlencoded")

	return &api{
		apiclient: newHttpClient(endpoint, header),
	}
}
