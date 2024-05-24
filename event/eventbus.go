package event

import "sync"

type (
	Event struct {
		Payload []byte
	}

	EventChan chan Event

	EventBus struct {
		lock        sync.RWMutex
		subscribers map[string][]EventChan
	}
)

func NewEventBus() *EventBus {
	return &EventBus{
		subscribers: make(map[string][]EventChan),
	}
}

func (e *EventBus) Subscribe(topic string) EventChan {
	e.lock.Lock()
	defer e.lock.Unlock()
	ch := make(EventChan)
	e.subscribers[topic] = append(e.subscribers[topic], ch)
	return ch
}

func (e *EventBus) UnSubscribe(topic string, ch EventChan) {
	e.lock.Lock()
	defer e.lock.Unlock()
	if v, ok := e.subscribers[topic]; ok {
		for i, vv := range v {
			if vv == ch {
				e.subscribers[topic] = append(v[:i], v[i+1:]...)
				close(ch)
				for range ch {
				}
				return
			}
		}
	}

}

func (e *EventBus) Publish(topic string, event Event) {
	e.lock.RLock()
	defer e.lock.RUnlock()
	subscribers := append([]EventChan{}, e.subscribers[topic]...)
	for _, subscriber := range subscribers {
		subscriber <- event
	}

	return
}
