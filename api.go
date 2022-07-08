package telego

import (
	"context"
	"fmt"
)

const (
	KeyApi = ModuleKey("telego.api")
)

func GetApiModule(ctx context.Context) *api {
	return ctx.Value(KeyApi).(*api)
}

type api struct {
	apiclient *httpClient
}

func newApi(token string, apiendpoint string) *api {
	endpoint := fmt.Sprintf(apiendpoint, token)

	return &api{
		apiclient: newHttpClient(endpoint),
	}
}

func (a *api) InitModule(ctx context.Context) (context.Context, error) {
	return context.WithValue(ctx, KeyApi, a), nil
}

func (a *api) GetMe() (*TypeGetMe, error) {
	getme := &TypeGetMe{}

	a.apiclient.setMethod("getMe")
	err := a.apiclient.doPost(getme)
	if err != nil {
		return nil, err
	}

	return getme, nil
}
