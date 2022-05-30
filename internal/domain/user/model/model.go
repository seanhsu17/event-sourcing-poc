package model

type User struct {
	UserId string `json:"userId"`
	Name   string `json:"name"`
	Age    int32  `json:"age"`
	Gender string `json:"gender"`
}

type EventPayload struct {
	Data     interface{} `json:"data"`
	Metadata Metadata    `json:"metadata"`
}

type Metadata struct {
	ModifiedFields []string `json:"modifiedFields,omitempty"`
}
