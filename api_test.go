package tggo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApiGetMe(t *testing.T) {
	api := newApi(token)

	getme, err := api.GetMe()

	assert.NoError(t, err)
	assert.NotNil(t, getme)
}
