package sys

import (
	"github.com/google/uuid"
	"log"
)

type HandlerID string
type EventType string

const (
	EventClose     EventType = "eventClose"
	EventMaximize  EventType = "eventMaximize"
	EventMinimize  EventType = "eventMinimize"
	EventResize    EventType = "eventResize"
	EventMove      EventType = "eventMove"
	EventGainFocus EventType = "eventGainFocus"
	EventLostFocus EventType = "eventLostFocus"

	EventMousePress   EventType = "eventMousePressed"
	EventMouseRelease EventType = "eventMouseReleased"
	EventMouseMove    EventType = "eventMouseMoved"
	EventMouseScroll  EventType = "eventMouseWheelScrolled"
	EventKeyPress     EventType = "eventKeyPressed"
	EventKeyRelease   EventType = "eventKeyReleased"
)

type EventStore struct {
	Handlers       map[EventType]map[HandlerID]func()
	ProducedEvents []EventType
}

func NewEventStore() *EventStore {
	return &EventStore{
		Handlers: make(map[EventType]map[HandlerID]func()),
	}
}

func (es *EventStore) Register(eventType EventType, handler func()) HandlerID {
	handlerID := uuid.NewString()

	if _, exists := es.Handlers[eventType]; !exists {
		es.Handlers[eventType] = make(map[HandlerID]func())
	}

	es.Handlers[eventType][HandlerID(handlerID)] = handler

	return HandlerID(handlerID)
}

func (es *EventStore) Unregister(eventType EventType, handlerID HandlerID) {
	if handlers, exists := es.Handlers[eventType]; exists {
		delete(handlers, handlerID)
		if len(handlers) == 0 {
			delete(es.Handlers, eventType)
		}
	}
}

func (es *EventStore) Trigger(eventType EventType) {
	es.ProducedEvents = append(es.ProducedEvents, eventType)
}

func (es *EventStore) ProcessEvents() {
	if len(es.ProducedEvents) == 0 {
		return
	}

	log.Printf("Processing %d events", len(es.ProducedEvents))
	for _, eventType := range es.ProducedEvents {
		if handlers, exists := es.Handlers[eventType]; exists {
			for _, handler := range handlers {
				handler()
			}
		}
	}
	es.ProducedEvents = nil
}
