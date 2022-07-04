package telego

type Telego struct {
	Api *api
}

func NewTelego(token string) *Telego {
	return &Telego{
		Api: newapi(token),
	}
}
