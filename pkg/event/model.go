package event

type Event struct {
	TraceID   string      `json:"traceID"`
	Type      string      `json:"eventType"`
	Version   int32       `json:"version"`
	Payload   interface{} `json:"payload"`
	Timestamp int64       `json:"timestamp"`
}
