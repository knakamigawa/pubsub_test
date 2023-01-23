package main

import (
	"cloud.google.com/go/pubsub"
	"context"
	"fmt"
	"log"
	"time"
	"os"
)

func main() {
	project := os.Getenv("PROJECT_NAME")

	bulk(project)
}

func bulk(project string) {
	ctx := context.Background()
	publish(ctx, project, "order01")
	publish(ctx, project, "order01")
	publish(ctx, project, "order01")
	publish(ctx, project, "order01")
	publish(ctx, project, "order01")
	publish(ctx, project, "order01")
	publish(ctx, project, "order01")
	publish(ctx, project, "order01")
	publish(ctx, project, "order01")
	publish(ctx, project, "order01")
	publish(ctx, project, "order01")
	publish(ctx, project, "order01")
	publish(ctx, project, "order02")
}

func publish(ctx context.Context, project, orderingKey string) {
	client, err := pubsub.NewClient(ctx, project)
	if err != nil {
		log.Fatal(err)
	}

	topic := client.Topic("pubsub_test")
	defer topic.Stop()

	topic.EnableMessageOrdering = true

	message := &pubsub.Message{
		OrderingKey: orderingKey,
		Data:        []byte(orderingKey + ":" + fmt.Sprintf("%s", time.Now().Format(time.RFC3339Nano))),
	}
	r := topic.Publish(ctx, message)
	if r == nil {
		log.Fatal("error publish")
	}
	log.Printf("%+v", r)
}
