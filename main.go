package main

import (
	"cloud.google.com/go/pubsub"
	"context"
	"fmt"
	"os"
	"sync/atomic"
)

func main() {
	ctx := context.Background()
	projectID := os.Getenv("PROJECT_ID")
	subscriptionID := os.Getenv("SUBSCRIPTION_ID")

	if err := subscribe(ctx, projectID, subscriptionID); err != nil {
		fmt.Println("error:%v", err)
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
