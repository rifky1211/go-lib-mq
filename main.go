package main

import (
	"context"
	"encoding/json"
	"github.com/rifky1211/lib-mq-go/rabbitmq"
)

type test struct {
	ID string `json:"id"`
}

func main() {
	// Init pubsub.
	ps, err := rabbitmq.New("amqp://guest:admin123@localhost:5672")
	if err != nil {
		return
	}

	defer ps.Close()

	var payload test
	payload.ID = "ini id"

	data, _ := json.Marshal(payload)

	ctx := context.Background()
	err = ps.Publish(ctx, "test", "testqueue", data, false)
	if err != nil {
		return
	}

}
