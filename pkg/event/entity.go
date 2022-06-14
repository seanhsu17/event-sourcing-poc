package event

// ReceiveRecord Table
type ReceiveRecord struct {
	Topic       string `bson:"topic"`
	TraceID     string `bson:"traceID"`
	EventID     string `bson:"eventID"`
	PublishTime int64  `bson:"publishTime"`
	ReceiveTime int64  `bson:"receiveTime"`
	CreatedTime int64  `bson:"createdTime"`
}

// PublishRecord Table
type PublishRecord struct {
	TraceID     string `bson:"traceID"`
	EventID     string `bson:"eventID"`
	Topic       string `bson:"topic"`
	Payload     string `bson:"payload"`
	PublishTime int64  `bson:"publishTime"`
	CreatedTime int64  `bson:"createdTime"`
}
