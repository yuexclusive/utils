package eventbus

import (
	"fmt"
	"testing"
	"time"
)

type Data struct {
	Name string
}

func Test_EventBus(t *testing.T) {
	bus := NewBus[Data]()

	bus.Subscribe("aa", func(topic, id string, d Data) {
		fmt.Println(topic, id, d)
	})

	bus.Subscribe("aa", func(topic, id string, d Data) {
		fmt.Println(topic, id, d)
	})

	bus.Publish("aa", Data{Name: "test"})

	time.Sleep(time.Second)
}
