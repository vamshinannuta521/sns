package trigger

type SvcInterface interface {
}

type Svc struct {
}

func NewSvc() *Svc {
	return &Svc{}
}
