package persistence

import (
	"fmt"

	"github.com/jerry-yt-chen/event-sourcing-poc/configs"
	"github.com/jerry-yt-chen/event-sourcing-poc/pkg/mongo"
)

func InitMongo() mongo.Service {
	client, err := mongo.Init(configs.C.Mongo)
	if err != nil {
		fmt.Printf("failed to connect to mongo err: %v\n", err)
		panic(err)
	}
	return client
}
