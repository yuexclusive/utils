package eventbus

import (
	"sync"

	"github.com/google/uuid"
)

type Channel[T any] struct {
	EventChan chan T
	Topic     string
	ID        string
}

type Bus[T any] struct {
	Subscriber map[string]map[string]Channel[T]
	Mutex      *sync.RWMutex
}

func (b *Bus[T]) Subscribe(topic string, f func(topic, id string, data T)) {
	b.Mutex.Lock()
	defer b.Mutex.Unlock()

	if _, exists := b.Subscriber[topic]; !exists {
		b.Subscriber[topic] = make(map[string]Channel[T])
	}

	id := uuid.New().String()
	b.Subscriber[topic][id] = Channel[T]{ID: id, Topic: topic, EventChan: make(chan T)}

	channel := b.Subscriber[topic][id]

	go func() {
		for val := range channel.EventChan {
			f(channel.Topic, channel.ID, val)
		}
	}()

}

func (b *Bus[T]) UnSubscribe(topic, id string) {
	b.Mutex.Lock()
	defer b.Mutex.Unlock()

	if _, exists := b.Subscriber[topic]; !exists {
		return
	}
	close(b.Subscriber[topic][id].EventChan)
	delete(b.Subscriber[topic], id)
}

func (b *Bus[T]) Publish(topic string, data T) {
	b.Mutex.RLock()
	defer b.Mutex.RUnlock()

	if _, exists := b.Subscriber[topic]; !exists {
		return
	}

	for _, c := range b.Subscriber[topic] {
		c.EventChan <- data
	}
}

func NewBus[T any]() *Bus[T] {
	return &Bus[T]{Subscriber: make(map[string]map[string]Channel[T]), Mutex: &sync.RWMutex{}}
}
