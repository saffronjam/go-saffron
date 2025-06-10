package core

import "github.com/saffronjam/go-sfml/public/sfml"

type KeyEvent struct {
	BaseEvent
	Type     EventType
	Code     sfml.KeyCode
	Scancode sfml.Scancode
	Alt      bool
	Control  bool
	Shift    bool
	System   bool
}

func (e *KeyEvent) EventType() EventType {
	return e.Type
}

type MouseButtonEvent struct {
	BaseEvent
	Type   EventType
	Button sfml.MouseButton
	X      int
	Y      int
}

func (e *MouseButtonEvent) EventType() EventType {
	return e.Type
}

type MouseMoveEvent struct {
	BaseEvent
	Type EventType
	X    int
	Y    int
}

func (e *MouseMoveEvent) EventType() EventType {
	return e.Type
}

type MouseWheelScrollEvent struct {
	BaseEvent
	Type  EventType
	Wheel sfml.MouseWheel
	Delta float32
	X     int
	Y     int
}

func (e *MouseWheelScrollEvent) EventType() EventType {
	return e.Type
}

type SizeEvent struct {
	BaseEvent
	Type   EventType
	Width  uint
	Height uint
}

func (e *SizeEvent) EventType() EventType {
	return e.Type
}

type TextEvent struct {
	BaseEvent
	Type    EventType
	Unicode uint
}

func (e *TextEvent) EventType() EventType {
	return e.Type
}

type ClosedEvent struct {
	BaseEvent
	Type EventType
}

func (e *ClosedEvent) EventType() EventType {
	return e.Type
}

type LostFocusEvent struct {
	BaseEvent
	Type EventType
}

func (e *LostFocusEvent) EventType() EventType {
	return e.Type
}

type GainedFocusEvent struct {
	BaseEvent
	Type EventType
}

func (e *GainedFocusEvent) EventType() EventType {
	return e.Type
}

type MouseEnteredEvent struct {
	BaseEvent
	Type EventType
}

func (e *MouseEnteredEvent) EventType() EventType {
	return e.Type
}

type MouseLeftEvent struct {
	BaseEvent
	Type EventType
}

func (e *MouseLeftEvent) EventType() EventType {
	return e.Type
}
