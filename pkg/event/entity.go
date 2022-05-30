package event

// PublishedRecord Table
type PublishedRecord string

// ReceivedRecord Table
type ReceivedRecord string

type Record struct {
	TraceID       string `bson:"traceID"`
	EventID       string `bson:"eventID"`
	EventType     string `bson:"eventType"`
	Publisher     string `bson:"publisher,omitempty"`
	Subscriber    string `bson:"subscriber,omitempty"`
	Version       int32  `bson:"version"`
	Payload       string `bson:"payload"`
	PublishedTime int64  `bson:"publishedTime,omitempty"`
	ReceivedTime  int64  `bson:"receivedTime,omitempty"`
	CreatedTime   int64  `bson:"createdTime,omitempty"`
}
