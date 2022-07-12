package telego

import "context"

type MwFunc func(ctx context.Context) (context.Context, error)

type IMiddleware interface {
	middleware(context.Context) (context.Context, error)
	init(context.Context) (context.Context, error)
}

//lint:ignore U1000 Ignore unused function temporarily for debugging
type mwChain struct {
	mwfunc MwFunc
	next   *mwChain
}

func newMwChain() *mwChain {
	return &mwChain{}
}

func (c *mwChain) exec(ctx context.Context) (context.Context, error) {
	newctx, err := c.mwfunc(ctx)
	if err != nil {
		return nil, err
	}

	if c.next == nil {
		return newctx, nil
	}
	return c.next.exec(newctx)
}

func (c *mwChain) addToChain(mwfunc MwFunc) {
	switch {
	case c.mwfunc == nil:
		c.mwfunc = mwfunc
	case c.next == nil:
		c.next = &mwChain{
			mwfunc: mwfunc,
		}
	default:
		c.next.addToChain(mwfunc)
	}

}
