package eventbus

import (
	"fmt"
	"reflect"
	"sync"

	"github.com/google/uuid"
)

type EventBus[T any] interface {
	Subscribe(topic string, f func(topic, id string, data T))
	UnSubscribe(topic, id string) error
	Publish(topic string, data T) error
}

type channel[T any] struct {
	eventChan chan T
	topic     string
	id        string
}

type bus[T any] struct {
	subscribers map[string]map[string]channel[T]
	mutex       *sync.RWMutex
}

func newbus[T any]() EventBus[T] {
	return &bus[T]{subscribers: make(map[string]map[string]channel[T]), mutex: new(sync.RWMutex)}
}

func (b *bus[T]) Subscribe(topic string, f func(topic, id string, data T)) {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	if _, exists := b.subscribers[topic]; !exists {
		b.subscribers[topic] = make(map[string]channel[T])
	}

	id := uuid.New().String()
	b.subscribers[topic][id] = channel[T]{id: id, topic: topic, eventChan: make(chan T)}

	channel := b.subscribers[topic][id]

	go func() {
		for val := range channel.eventChan {
			f(channel.topic, channel.id, val)
		}
	}()

}

func (b *bus[T]) UnSubscribe(topic, id string) error {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	if _, exists := b.subscribers[topic]; !exists {
		return fmt.Errorf("can't find topic %s", topic)
	}

	if ch, exists := b.subscribers[topic][id]; !exists {
		return fmt.Errorf("can't find id %s", id)
	} else {
		close(ch.eventChan)
		delete(b.subscribers[topic], id)
	}
	return nil
}

func (b *bus[T]) Publish(topic string, data T) error {
	b.mutex.RLock()
	defer b.mutex.RUnlock()

	if _, exists := b.subscribers[topic]; !exists {
		return fmt.Errorf("can't find topic %s", topic)
	}

	for _, c := range b.subscribers[topic] {
		c.eventChan <- data
	}
	return nil
}

var _cache map[string]interface{}

func Init[T any](busName string) {
	if _cache == nil {
		_cache = make(map[string]interface{})
	}
	_cache[busName] = newbus[T]()
}

func Get[T any](busName string) (EventBus[T], error) {
	var res EventBus[T]
	if v, ok := _cache[busName]; !ok {
		return res, fmt.Errorf("can't find bus %s", busName)
	} else {
		if v2, ok := v.(EventBus[T]); !ok {
			var t T
			tname := ""
			if rt := reflect.TypeOf(t); rt != nil {
				tname = rt.Name()
			}
			return res, fmt.Errorf(`can't find bus "%s" as type "EventBus[%s]"`, busName, tname)
		} else {
			res = v2
		}
	}
	return res, nil
}
