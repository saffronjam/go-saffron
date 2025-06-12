package core

import (
	"github.com/saffronjam/go-sfml/public/sfml"
)

type Window struct {
	WindowProps
	sfmlWindow            *sfml.RenderWindow
	preFullscreenPosition *sfml.Vector2i // Used to restore position after exiting fullscreen
}

type WindowProps struct {
	Width, Height uint
	Fullscreen    bool
	Borderless    bool
	Title         string
	Antialiasing  uint
	BitsPerPixel  uint
}

func FullscreenProps(title string) *WindowProps {
	return &WindowProps{
		Fullscreen:   true,
		Title:        title,
		Antialiasing: 0,
		BitsPerPixel: 32,
	}
}

func NewWindow(props *WindowProps) (*Window, error) {
	windowFlags := sfml.DefaultStyle
	if props.Fullscreen {
		windowFlags |= sfml.Fullscreen
	}

	sfmlWindow := sfml.NewRenderWindow(sfml.VideoMode{
		Width:        uint32(props.Width),
		Height:       uint32(props.Height),
		BitsPerPixel: uint32(props.BitsPerPixel),
	}, props.Title, uint32(windowFlags), &sfml.ContextSettings{
		AntialiasingLevel: uint32(props.Antialiasing),
	})

	return &Window{
		WindowProps: *props,
		sfmlWindow:  sfmlWindow,
	}, nil
}

func (w *Window) Close() {
	w.sfmlWindow.Close()
}

func (w *Window) ProduceEvents() []Event {
	var events []Event
	for {
		sfEvent, hasMore := w.sfmlWindow.PollEvent()
		if !hasMore {
			break
		}
		event := sfmlEventToSaffronEvent(sfEvent)
		if event != nil {
			events = append(events, event)
		}
	}
	return events
}

func (w *Window) Clear() {
	w.sfmlWindow.Clear(sfml.Color{R: 0, G: 0, B: 0, A: 255}) // Clear with black color
}

func (w *Window) Display() {
	w.sfmlWindow.Display()
}

func (w *Window) IsOpen() bool {
	return w.sfmlWindow.IsOpen()
}

func (w *Window) Position() (x, y int) {
	pos := w.sfmlWindow.Position()
	return int(pos.X), int(pos.Y)
}

func (w *Window) Size() (width, height int) {
	size := w.sfmlWindow.Size()
	return int(size.X), int(size.Y)
}

func (w *Window) SetFullscreen(fullscreen bool) {
	if fullscreen && !w.Fullscreen {
		w.preFullscreenPosition = w.sfmlWindow.Position()
		w.sfmlWindow.Free()
		w.sfmlWindow = sfml.NewRenderWindow(sfml.VideoMode{
			Width:        uint32(w.Width),
			Height:       uint32(w.Height),
			BitsPerPixel: uint32(w.BitsPerPixel),
		}, w.Title, uint32(sfml.Fullscreen), &sfml.ContextSettings{
			AntialiasingLevel: uint32(w.Antialiasing),
		})
	} else if !fullscreen && w.Fullscreen {
		if w.preFullscreenPosition == nil {
			w.preFullscreenPosition = &sfml.Vector2i{X: 0, Y: 0} // Default position if not set
		}
		w.sfmlWindow.Free()
		w.sfmlWindow = sfml.NewRenderWindow(sfml.VideoMode{
			Width:        uint32(w.Width),
			Height:       uint32(w.Height),
			BitsPerPixel: uint32(w.BitsPerPixel),
		}, w.Title, uint32(sfml.DefaultStyle), &sfml.ContextSettings{
			AntialiasingLevel: uint32(w.Antialiasing),
		})
		w.sfmlWindow.SetPosition(*w.preFullscreenPosition)
	}
	w.Fullscreen = fullscreen
}

func (w *Window) SfmlHandle() *sfml.RenderWindow {
	return w.sfmlWindow
}

func sfmlEventTypeToSaffronEventType(sfmlType sfml.EventType) EventType {
	switch sfmlType {
	case sfml.EvtKeyPressed:
		return EventKeyPressed
	case sfml.EvtKeyReleased:
		return EventKeyReleased
	case sfml.EvtMouseButtonPressed:
		return EventMouseButtonPressed
	case sfml.EvtMouseButtonReleased:
		return EventMouseButtonReleased
	case sfml.EvtMouseMoved:
		return EventMouseMoved
	case sfml.EvtMouseWheelScrolled:
		return EventMouseWheelScrolled
	case sfml.EvtResized:
		return EventResized
	case sfml.EvtTextEntered:
		return EventTextEntered
	case sfml.EvtClosed:
		return EventClosed
	case sfml.EvtLostFocus:
		return EventLostFocus
	default:
		panic("Unknown SFML event type: " + string(sfmlType))
	}
}

func sfmlEventToSaffronEvent(sfEvent sfml.Event) Event {
	base := BaseEvent{
		NativeHandle: sfEvent,
		Tags:         map[string]struct{}{"sfml": {}},
	}

	switch sfEvent.EventType() {
	case sfml.EvtKeyPressed, sfml.EvtKeyReleased:
		e := sfEvent.(*sfml.KeyEvent)
		return &KeyEvent{
			BaseEvent: base,
			Type:      sfmlEventTypeToSaffronEventType(e.EventType()),
			Code:      e.Code,
			Scancode:  e.Scancode,
			Alt:       e.Alt,
			Control:   e.Control,
			Shift:     e.Shift,
			System:    e.System,
		}
	case sfml.EvtMouseButtonPressed, sfml.EvtMouseButtonReleased:
		e := sfEvent.(*sfml.MouseButtonEvent)
		return &MouseButtonEvent{
			BaseEvent: base,
			Type:      sfmlEventTypeToSaffronEventType(e.EventType()),
			Button:    e.Button,
			X:         int(e.X),
			Y:         int(e.Y),
		}
	case sfml.EvtMouseMoved:
		e := sfEvent.(*sfml.MouseMoveEvent)
		return &MouseMoveEvent{
			BaseEvent: base,
			Type:      EventMouseMoved,
			X:         int(e.X),
			Y:         int(e.Y),
		}
	case sfml.EvtMouseWheelScrolled:
		e := sfEvent.(*sfml.MouseWheelScrollEvent)
		return &MouseWheelScrollEvent{
			BaseEvent: base,
			Type:      sfmlEventTypeToSaffronEventType(e.EventType()),
			Wheel:     e.Wheel,
			Delta:     e.Delta,
			X:         int(e.X),
			Y:         int(e.Y),
		}
	case sfml.EvtResized:
		e := sfEvent.(*sfml.SizeEvent)
		return &SizeEvent{
			BaseEvent: base,
			Type:      sfmlEventTypeToSaffronEventType(e.EventType()),
			Width:     uint(e.Width),
			Height:    uint(e.Height),
		}
	case sfml.EvtTextEntered:
		e := sfEvent.(*sfml.TextEvent)
		return &TextEvent{
			BaseEvent: base,
			Type:      sfmlEventTypeToSaffronEventType(e.EventType()),
			Unicode:   uint(e.Unicode),
		}
	case sfml.EvtClosed:
		e := sfEvent.(*sfml.ClosedEvent)
		return &ClosedEvent{
			BaseEvent: base,
			Type:      sfmlEventTypeToSaffronEventType(e.EventType()),
		}
	case sfml.EvtLostFocus:
		e := sfEvent.(*sfml.LostFocusEvent)
		return &LostFocusEvent{
			BaseEvent: base,
			Type:      sfmlEventTypeToSaffronEventType(e.EventType()),
		}
	default:
		return nil
	}
}
