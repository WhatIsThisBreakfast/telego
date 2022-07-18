package tggo

import (
	"context"
	"fmt"
)

const (
	KeyApi = ContextKey("tggo.api")
)

func GetApiFromCtx(ctx context.Context) *api {
	return ctx.Value(KeyApi).(*api)
}

type api struct {
	apiclient *httpClient
	ctx       context.Context
}

func newApi(token string) *api {
	endpoint := fmt.Sprintf(c_apiendpoint, token)

	return &api{
		apiclient: newHttpClient(endpoint),
	}
}

// middleware init
func (a *api) init(ctx context.Context) (context.Context, error) {
	apictx := context.WithValue(ctx, KeyApi, a)

	return apictx, nil
}

// middleware func
func (a *api) middleware(ctx context.Context) (context.Context, error) {
	a.ctx = ctx

	return ctx, nil
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
