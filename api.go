package telego

import (
	"fmt"
)

type api struct {
	apiclient *httpClient
}

func newApi(token string, apiendpoint string) *api {
	endpoint := fmt.Sprintf(apiendpoint, token)

	return &api{
		apiclient: newHttpClient(endpoint),
	}
}

func (a *api) GetMe() (*type_GetMe, error) {
	getme := &type_GetMe{}

	a.apiclient.setMethod("getMe")
	err := a.apiclient.doPost(getme)
	if err != nil {
		return nil, err
	}

	return getme, nil
}
