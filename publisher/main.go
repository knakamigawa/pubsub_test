package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	project := os.Getenv("PROJECT_NAME")

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/bulk", bulk(project))

	e.Logger.Fatal(e.Start(":8080"))
}

func bulk(project string) func(c echo.Context) error {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		publish(ctx, project, "order01")
		publish(ctx, project, "order01")
		publish(ctx, project, "order01")
		publish(ctx, project, "order01")
		publish(ctx, project, "order01")
		publish(ctx, project, "order01")
		publish(ctx, project, "order02")

		return c.String(http.StatusOK, "bulk publish")
	}
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
