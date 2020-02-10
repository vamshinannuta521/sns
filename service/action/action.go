package action

type SvcInterface interface {
}

type Svc struct {
}

func NewSvc() *Svc {
	return &Svc{}
}
