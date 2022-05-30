package event

type Publisher interface {
	Send(traceID string, event Event) error
}
