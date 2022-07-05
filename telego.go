package telego

type Telego struct {
	Api *api
}

func NewTelego(token string) *Telego {
	return &Telego{
		Api: newApi(token, c_apiendpoint),
	}
}
