package pulsar

import (
	"context"

	"github.com/apache/pulsar-client-go/pulsar"
)

func Consume(ctx context.Context, url string, topic string, subscriptionName string, subscriptionType pulsar.SubscriptionType) {
	client, err := pulsar.NewClient(pulsar.ClientOptions{URL: url})

	if err != nil {
		panic(err)
	}

	consumer, err := client.Subscribe(pulsar.ConsumerOptions{
		Topic:            topic,
		SubscriptionName: subscriptionName,
		Type:             subscriptionType,
	})

	if err != nil {
		panic(err)
	}

	msg, err := consumer.Receive(ctx)

	if err != nil {
		panic(err)
	}
}
