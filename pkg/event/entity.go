package event

// ReceivedRecord Table
type ReceivedRecord struct {
	Topic        string `bson:"topic"`
	TraceID      string `bson:"traceID"`
	EventID      string `bson:"eventID"`
	Subscriber   string `bson:"subscriber,omitempty"`
	ReceivedTime int64  `bson:"receivedTime,omitempty"`
	CreatedTime  int64  `bson:"createdTime"`
}

// PublishedRecord Table
type PublishedRecord struct {
	TraceID       string `bson:"traceID"`
	EventID       string `bson:"eventID"`
	EventType     string `bson:"eventType"`
	Topic         string `bson:"topic"`
	Publisher     string `bson:"publisher,omitempty"`
	Payload       string `bson:"payload"`
	PublishedTime int64  `bson:"publishedTime,omitempty"`
	CreatedTime   int64  `bson:"createdTime,omitempty"`
}
