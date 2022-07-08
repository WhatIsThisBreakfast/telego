package telego

import "context"

type ModuleKey string

type TelegoModule interface {
	InitModule(ctx context.Context) (context.Context, error)
}
