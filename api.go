package telego

type api struct {
	token string
}

func newapi(token string) *api {
	return &api{
		token: token,
	}
}
