package consumer

import (
	"fmt"
	"testing"

	"github.com/Shopify/sarama"
)

type handler struct{}

// Setup is run at the beginning of a new session, before ConsumeClaim.
func (h *handler) Setup(session sarama.ConsumerGroupSession) error {
	// session.ResetOffset("g1", 0, 0, "")
	// panic("not implemented") // TODO: Implement
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
// but before the offsets are committed for the very last time.
func (h *handler) Cleanup(_ sarama.ConsumerGroupSession) error {
	// panic("not implemented") // TODO: Implement
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
// Once the Messages() channel is closed, the Handler must finish its processing
// loop and exit.
func (h *handler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	// panic("not implemented") // TODO: Implement
	for message := range claim.Messages() {
		fmt.Println(string(message.Value))
		session.MarkMessage(message, "")
		// time.Sleep(time.Minute)
	}
	return nil
}

func TestConsume(t *testing.T) {
	c := NewConsumer([]string{"localhost:9092"}, "love", "g3", new(handler))

	c.Consume()
}
