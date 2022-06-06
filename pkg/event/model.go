package event

import (
	"encoding/json"
)

const (
	TraceAttribute   = "Cloud-Trace-Context"
	OptionsAttribute = "Options"
)

type Metadata map[string]string

type PublishOption struct {
	// Source specify the publishing source
	Source string `json:"source"`
	// EventType specify the action of the event
	EventType string `json:"eventType"`
	// Key specify the key of the message (supported by kafka only)
	Key string `json:"key"`
	// Timestamp specify event published time
	Timestamp int64 `json:"timestamp"`
}

func (opt *PublishOption) ToString() string {
	msgBytes, err := json.Marshal(opt)
	if err != nil {
		return ""
	}
	return string(msgBytes)
}
