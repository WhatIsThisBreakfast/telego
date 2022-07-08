package telego

import (
	"context"
	"fmt"
)

//lint:ignore U1000 Ignore unused function temporarily for debugging
type Telego struct {
	token   string
	ctx     context.Context
	modules []TelegoModule
}

func NewTelego(token string) *Telego {
	return &Telego{
		token: token,
	}
}

func (t *Telego) AddModule(module TelegoModule) {
	t.modules = append(t.modules, module)
}

func (t *Telego) addStandartModules() {
	t.AddModule(newApi(t.token, c_apiendpoint))
}

func (t *Telego) init() error {
	t.addStandartModules()

	if err := t.initContext(); err != nil {
		return err
	}

	return nil
}

func (t *Telego) initContext() error {
	ctx := context.Background()
	var err error

	//load modules
	for _, module := range t.modules {
		ctx, err = module.InitModule(ctx)
		if err != nil {
			return err
		}

		if ctx == nil {
			return fmt.Errorf("INIT MODULES ERROR{ description: ctx = nil }")
		}
	}

	return nil
}
