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
	}
}

func (t *Telego) Middleware(mw IMiddleware) {
	newctx, mwfunc := mw.middleware(t.ctx)
	if newctx != nil {
		t.ctx = newctx
	}

	t.mwchain.addToChain(mwfunc)
}

func (t *Telego) initMiddlware() {
	t.Middleware(newApi(t.token))
}

func (t *Telego) init() {
	t.initMiddlware()
}
