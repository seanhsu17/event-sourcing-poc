package main

import (
	"context"
	"crypto/md5"
	"fmt"
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/go-co-op/gocron"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/api/iterator"

	"github.com/jerry-yt-chen/event-sourcing-poc/configs"
	"github.com/jerry-yt-chen/event-sourcing-poc/pkg/event"
	"github.com/jerry-yt-chen/event-sourcing-poc/pkg/mongo"
)

func main() {
	configs.InitConfigs()
	logrus.Info("start scheduler")

	// TODO should keep offset to shared storage
	var pubOffset int64
	var subOffset int64

	// Mongo
	mongoSvc, _ := mongo.Init(configs.C.Mongo)
	s := gocron.NewScheduler(time.UTC)
	s.Every(5).Second().Do(func() {
		topicSubscriberCountMap := make(map[string]int)
		revCountMap := make(map[string]int)
		total := 0
		// Get published records
		pubColl := mongoSvc.Collection(new(event.PublishedRecord))
		var pubRecords []event.PublishedRecord
		pubColl.Find(context.Background(), bson.M{"publishedTime": bson.M{"$gt": pubOffset}}).Sort("publishedTime").All(&pubRecords)

		// check topic and subscription count
		for _, record := range pubRecords {
			if _, ok := topicSubscriberCountMap[record.Topic]; ok {
				logrus.Warn(record.Topic, " is already exist")
				continue
			}
			subscriptions, err := listSubscriptions(configs.C.Pub.ProjectID, record.Topic)
			if err != nil {
				fmt.Println("err", err)
			}
			key := fmt.Sprintf("%s:%s:%s", record.Topic, record.TraceID, record.EventID)
			hash := md5.Sum([]byte(key))
			count := len(subscriptions)
			topicSubscriberCountMap[string(hash[:])] = count
			total += count
		}

		if len(pubRecords) > 0 {
			pubOffset = pubRecords[len(pubRecords)-1].PublishedTime
		} else {
			return
		}

		// Get receive records after offset
		recColl := mongoSvc.Collection(new(event.ReceivedRecord))
		var revRecords []event.ReceivedRecord
		filter := bson.M{
			"receivedTime": bson.M{
				"$gt": subOffset,
			},
		}
		// Check receive records
		recColl.Find(context.Background(), filter).Sort("receivedTime").All(&revRecords)
		missed := 0
		for _, record := range revRecords {
			fmt.Printf("rev record: %s, %s, %s\n", record.Topic, record.TraceID, record.EventID)
			key := fmt.Sprintf("%s:%s:%s", record.Topic, record.TraceID, record.EventID)
			hash := md5.Sum([]byte(key))
			revCountMap[string(hash[:])] += 1
		}

		for key, expCount := range topicSubscriberCountMap {
			if actualCount, ok := revCountMap[key]; !ok {
				missed += 1
			} else if (expCount - actualCount) > 0 {
				missed += 1
			}
		}

		if total != 0 {
			fmt.Printf("total: %v \n", total)
			fmt.Printf("missed: %v \n", missed)
			completionRate := (float32(total-missed) / float32(total)) * 100
			fmt.Printf("completion rate: %f \n", completionRate)
			subOffset = revRecords[len(revRecords)-1].ReceivedTime
		}
	})
	s.StartBlocking()
}

func listSubscriptions(projectID, topicID string) ([]*pubsub.Subscription, error) {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return nil, fmt.Errorf("pubsub.NewClient: %v", err)
	}
	defer client.Close()

	var subs []*pubsub.Subscription

	it := client.Topic(topicID).Subscriptions(ctx)
	for {
		sub, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("next: %v", err)
		}
		subs = append(subs, sub)
	}
	return subs, nil
}
