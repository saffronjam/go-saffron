package sys

import (
	"github.com/google/uuid"
	"github.com/saffronjam/go-sfml/public/sfml"
)

type HandlerID string
type EventType string

const (
	EventClosed      EventType = "eventClosed"
	EventMaximized   EventType = "eventMaximized"
	EventMinimized   EventType = "eventMinimized"
	EventResized     EventType = "eventResized"
	EventMoved       EventType = "eventMoved"
	EventGainedFocus EventType = "eventGainFocus"
	EventLostFocus   EventType = "eventLostFocus"

	EventMouseButtonPressed  EventType = "eventMousePressed"
	EventMouseButtonReleased EventType = "eventMouseReleased"
	EventMouseMoved          EventType = "eventMouseMoved"
	EventMouseWheelScrolled  EventType = "eventMouseWheelScrolled"
	EventKeyPressed          EventType = "eventKeyPressed"
	EventKeyReleased         EventType = "eventKeyReleased"
	EventTextEntered         EventType = "eventTextEntered"
)

type EventProducer interface {
	ProduceEvents() []Event
}

type Event interface {
	EventType() EventType
	EventTags() map[string]struct{}
	SfmlHandle() sfml.Event
}

type BaseEvent struct {
	NativeHandle any
	Tags         map[string]struct{}
}

func (e *BaseEvent) EventTags() map[string]struct{} {
	if e.Tags == nil {
		e.Tags = make(map[string]struct{})
	}
	return e.Tags
}

func (e *BaseEvent) SfmlHandle() sfml.Event {
	if e.NativeHandle == nil {
		return nil
	}
	if sfmlEvent, ok := e.NativeHandle.(sfml.Event); ok {
		return sfmlEvent
	}
	panic("BaseEvent.NativeHandle is not a valid SFML event")
}

type EventStore struct {
	Producers      []EventProducer
	Handlers       map[EventType]map[HandlerID]func(any)
	TagHandlers    map[string]map[HandlerID]func(any)
	ProducedEvents []Event
}

func NewEventStore() *EventStore {
	return &EventStore{
		Handlers:    make(map[EventType]map[HandlerID]func(any)),
		TagHandlers: make(map[string]map[HandlerID]func(any)),
	}
}

func (es *EventStore) RegisterProducer(producer EventProducer) {
	es.Producers = append(es.Producers, producer)
}

func (es *EventStore) RegisterHandler(handler func(any), eventTypes ...EventType) []HandlerID {
	handlerIds := make([]HandlerID, 0, len(eventTypes))
	for _, eventType := range eventTypes {
		if es.Handlers[eventType] == nil {
			es.Handlers[eventType] = make(map[HandlerID]func(any))
		}
		id := uuid.NewString()
		es.Handlers[eventType][HandlerID(id)] = func(e any) {
			handler(e)
		}
		handlerIds = append(handlerIds, HandlerID(id))
	}
	return handlerIds
}

func (es *EventStore) RegisterHandlerByTags(handler func(any), tags ...string) []HandlerID {
	handlerIds := make([]HandlerID, 0, len(tags))
	for _, tag := range tags {
		if es.TagHandlers[tag] == nil {
			es.TagHandlers[tag] = make(map[HandlerID]func(any))
		}
		id := uuid.NewString()
		es.TagHandlers[tag][HandlerID(id)] = func(e any) {
			handler(e)
		}
		handlerIds = append(handlerIds, HandlerID(id))
	}
	return handlerIds
}

func (es *EventStore) Unregister(eventType EventType, handlerID HandlerID) {
	if handlers, exists := es.Handlers[eventType]; exists {
		delete(handlers, handlerID)
		if len(handlers) == 0 {
			delete(es.Handlers, eventType)
		}
	}
}

func (es *EventStore) ProcessEvents() {
	for _, producer := range es.Producers {
		if producer == nil {
			continue
		}
		es.ProducedEvents = append(es.ProducedEvents, producer.ProduceEvents()...)
	}

	if len(es.ProducedEvents) == 0 {
		return
	}

	for _, event := range es.ProducedEvents {
		if handlers, exists := es.Handlers[event.EventType()]; exists {
			for _, handler := range handlers {
				handler(event)
			}
		}

		for tag, _ := range event.EventTags() {
			if tagHandlers, exists := es.TagHandlers[tag]; exists {
				for _, handler := range tagHandlers {
					handler(event)
				}
			}
		}
	}
	es.ProducedEvents = nil
}
