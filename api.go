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

func (a *api) GetMe() (*type_GetMe, error) {
	getme := &type_GetMe{}

	a.apiclient.setMethod("getMe")
	err := a.apiclient.do(getme)
	if err != nil {
		return nil, err
	}

	return getme, nil
}
