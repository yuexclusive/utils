package producer

import "testing"

func Test_Produce(t *testing.T) {
	producer, err := NewProducer([]string{"localhost:9092"}, "love2")
	if err != nil {
		t.Error(err)
	}
	if err := producer.ProduceManyString([]string{"aaa", "bbb"}); err != nil {
		t.Error(err)
	}
}
