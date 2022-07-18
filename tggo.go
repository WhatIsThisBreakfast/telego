package tggo

import (
	"context"
)

type ContextKey string

//lint:ignore U1000 Ignore unused function temporarily for debugging
type Tggo struct {
	token   string
	ctx     context.Context
	mwchain *mwChain
}

func NewTggo(token string) *Tggo {
	return &Tggo{
		token:   token,
		mwchain: newMwChain(),
		ctx:     context.Background(),
	}
}

func (t *Tggo) Middleware(mw IMiddleware) {
	newctx, _ := mw.init(t.ctx)

	//TODO: add err output after adding log
	//if err != nil {
	//}

	if newctx != nil {
		t.ctx = newctx
	}

	t.mwchain.addToChain(mw.middleware)
}

func (t *Tggo) initMiddlware() {
	t.Middleware(newApi(t.token))
}

func (t *Tggo) init() {
	t.initMiddlware()
}
