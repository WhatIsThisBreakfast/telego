package tggo

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	tmwkey1 = ContextKey("tggo.test.tmw1")
	tmwkey2 = ContextKey("tggo.test.tmw2")
	tmwkey3 = ContextKey("tggo.test.tmw3")
)

//lint:ignore U1000 Ignore unused function temporarily for debugging
type testMwStruct struct {
	key     ContextKey
	counter int
}

func (tmw *testMwStruct) init(ctx context.Context) (context.Context, error) {
	tmwctx := context.WithValue(ctx, tmw.key, tmw)

	return tmwctx, nil
}

func (tmw *testMwStruct) middleware(ctx context.Context) (context.Context, error) {
	tmw.counter += 1

	return ctx, nil
}

func TestTggoMiddleware(t *testing.T) {
	tmw1 := &testMwStruct{
		key:     tmwkey1,
		counter: 4,
	}
	tmw2 := &testMwStruct{
		key:     tmwkey2,
		counter: 5,
	}
	tmw3 := &testMwStruct{
		key:     tmwkey3,
		counter: 6,
	}

	tlg.Middleware(tmw1)
	tlg.Middleware(tmw2)
	tlg.Middleware(tmw3)

	tlg.mwchain.exec(tlg.ctx)
	tlg.mwchain.exec(tlg.ctx)

	tmw := tlg.ctx.Value(tmwkey1).(*testMwStruct)
	assert.Equal(t, tmw.counter, 6)

	tmw = tlg.ctx.Value(tmwkey2).(*testMwStruct)
	assert.Equal(t, tmw.counter, 7)

	tmw = tlg.ctx.Value(tmwkey3).(*testMwStruct)
	assert.Equal(t, tmw.counter, 8)
}
