package main

import (
	"cloud.google.com/go/pubsub"
	"context"
	"fmt"
	"os"
	"sync/atomic"
	"time"
)

func main() {
	ctx := context.Background()
	projectID := os.Getenv("PROJECT_ID")
	subscriptionID := os.Getenv("SUBSCRIPTION_ID")

	for {
		if err := subscribe(ctx, projectID, subscriptionID); err != nil {
			fmt.Println("error:%v", err)
		}
	}

}

// Pub/SubからSubscribeする
func subscribe(ctx context.Context, projectID, subscriptionID string) error {
	// projectID := "my-project-id"
	// subID := "my-sub"
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("pubsub.NewClient: %v", err)
	}
	defer client.Close()

	sub := client.Subscription(subscriptionID)

	// Receive messages for 10 seconds, which simplifies testing.
	// Comment this out in production, since `Receive` should
	// be used as a long running operation.
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	var received int32
	err = sub.Receive(ctx, func(_ context.Context, msg *pubsub.Message) {
		fmt.Printf("Got message: %q\n", string(msg.Data))
		atomic.AddInt32(&received, 1)
		msg.Ack()
	})
	if err != nil {
		return fmt.Errorf("sub.Receive: %v", err)
	}
	fmt.Printf("Received %d messages\n", received)

	return nil
}
