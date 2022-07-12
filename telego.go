package telego

import (
	"context"
)

type ContextKey string

//lint:ignore U1000 Ignore unused function temporarily for debugging
type Telego struct {
	token   string
	ctx     context.Context
	mwchain *mwChain
}

func NewTelego(token string) *Telego {
	return &Telego{
		token:   token,
		mwchain: newMwChain(),
		ctx:     context.Background(),
	}
}

func (t *Telego) Middleware(mw IMiddleware) {
	newctx, _ := mw.init(t.ctx)

	//TODO: add err output after adding log
	//if err != nil {
	//}

	if newctx != nil {
		t.ctx = newctx
	}

	t.mwchain.addToChain(mw.middleware)
}

func (t *Telego) initMiddlware() {
	t.Middleware(newApi(t.token))
}

func (t *Telego) init() {
	t.initMiddlware()
}
