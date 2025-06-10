package core

import "github.com/google/uuid"

type SubscriberList[Payload any] struct {
	subscribers map[HandlerID]func(Payload)
}

func (sl *SubscriberList[Payload]) Subscribe(fn func(Payload)) HandlerID {
	if sl.subscribers == nil {
		sl.subscribers = make(map[HandlerID]func(Payload))
	}
	handlerID := uuid.NewString()
	sl.subscribers[HandlerID(handlerID)] = fn
	return HandlerID(handlerID)
}

func (sl *SubscriberList[Payload]) Unsubscribe(handlerID HandlerID) {
	if sl.subscribers != nil {
		delete(sl.subscribers, handlerID)
	}
}

func (sl *SubscriberList[Payload]) Subscribers() map[HandlerID]func(Payload) {
	if sl.subscribers == nil {
		return nil
	}
	return sl.subscribers
}

func (sl *SubscriberList[Payload]) Clear() {
	if sl.subscribers != nil {
		sl.subscribers = make(map[HandlerID]func(Payload))
	}
}

func (sl *SubscriberList[Payload]) Has(handlerID HandlerID) bool {
	if sl.subscribers == nil {
		return false
	}
	_, exists := sl.subscribers[handlerID]
	return exists
}

func (sl *SubscriberList[Payload]) Trigger(payload Payload) {
	for _, fn := range sl.subscribers {
		if fn != nil {
			fn(payload)
		}
	}
}
