package model

type User struct {
	UserId string `json:"userId"`
	Name   string `json:"name"`
	Age    int32  `json:"age"`
	Gender string `json:"gender"`
}
