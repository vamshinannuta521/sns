package event

type EventSvc interface {
	Get() string
}

type Event struct {
}

func NewEventSvc() *Event {
	return &Event{}
}

func (e *Event) Get() string {
	return "event"
}
