package event

type Message struct {
	EventID   string      `json:"eventID"`
	TraceID   string      `json:"traceID"`
	Source    string      `json:"source"`
	Type      string      `json:"eventType"`
	Version   int32       `json:"version"`
	Payload   interface{} `json:"payload"`
	Topic     string      `json:"topic"`
	Timestamp int64       `json:"timestamp"`
}
