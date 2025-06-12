package input

import (
	"github.com/saffronjam/go-saffron/pkg/core"
	"github.com/saffronjam/go-sfml/public/sfml"
)

type Store struct {
	keyboardState     map[sfml.KeyCode]bool
	prevKeyboardState map[sfml.KeyCode]bool

	mouseButtonState      map[sfml.MouseButton]bool
	prevMouseButtonState  map[sfml.MouseButton]bool
	mousePosition         *sfml.Vector2f
	prevMousePosition     *sfml.Vector2f
	verticalScrollDelta   float32
	horizontalScrollDelta float32
}

var Input *Store

func SetGlobalInput(input *Store) {
	Input = input
}

func NewInput(eventStore *core.EventStore) *Store {
	input := &Store{
		keyboardState:        make(map[sfml.KeyCode]bool),
		prevKeyboardState:    make(map[sfml.KeyCode]bool),
		mouseButtonState:     make(map[sfml.MouseButton]bool),
		prevMouseButtonState: make(map[sfml.MouseButton]bool),
		mousePosition:        &sfml.Vector2f{X: 0, Y: 0},
		prevMousePosition:    &sfml.Vector2f{X: 0, Y: 0},
	}

	eventStore.RegisterHandler(func(e any) {
		event := e.(core.Event)
		eType := event.EventType()

		switch eType {
		case core.EventKeyPressed:
			input.onKeyPressed(event.(*core.KeyEvent))
		case core.EventKeyReleased:
			input.onKeyReleased(event.(*core.KeyEvent))
		case core.EventMouseButtonPressed:
			input.onMouseButtonPressed(event.(*core.MouseButtonEvent))
		case core.EventMouseButtonReleased:
			input.onMouseButtonReleased(event.(*core.MouseButtonEvent))
		case core.EventMouseMoved:
			input.onMouseMoved(event.(*core.MouseMoveEvent))
		case core.EventMouseWheelScrolled:
			input.onMouseWheelScrolled(event.(*core.MouseWheelScrollEvent))
		}
	}, core.EventKeyPressed, core.EventKeyReleased, core.EventMouseButtonPressed, core.EventMouseButtonReleased, core.EventMouseMoved, core.EventMouseWheelScrolled)

	return input
}

func (s *Store) PostUpdate() {
	for code := range s.keyboardState {
		s.prevKeyboardState[code] = s.keyboardState[code]
	}
	for button := range s.mouseButtonState {
		s.prevMouseButtonState[button] = s.mouseButtonState[button]
	}
	s.prevMousePosition = &sfml.Vector2f{X: s.mousePosition.X, Y: s.mousePosition.Y}

	// Reset scroll deltas
	s.verticalScrollDelta = 0
	s.horizontalScrollDelta = 0
}

func (s *Store) IsKeyDown(code sfml.KeyCode) bool {
	return s.keyboardState[code]
}

func (s *Store) IsKeyPressed(code sfml.KeyCode) bool {
	return s.keyboardState[code] && !s.prevKeyboardState[code]
}

func (s *Store) IsKeyReleased(code sfml.KeyCode) bool {
	return !s.keyboardState[code] && s.prevKeyboardState[code]
}

func (s *Store) IsMouseButtonDown(button sfml.MouseButton) bool {
	return s.mouseButtonState[button]
}

func (s *Store) IsMouseButtonPressed(button sfml.MouseButton) bool {
	return s.mouseButtonState[button] && !s.prevMouseButtonState[button]
}

func (s *Store) MousePosition() *sfml.Vector2f {
	return s.mousePosition
}

func (s *Store) MouseSwipe() *sfml.Vector2f {
	return &sfml.Vector2f{
		X: s.mousePosition.X - s.prevMousePosition.X,
		Y: s.mousePosition.Y - s.prevMousePosition.Y,
	}
}

func (s *Store) IsMouseButtonReleased(button sfml.MouseButton) bool {
	return !s.mouseButtonState[button] && s.prevMouseButtonState[button]
}

func (s *Store) VerticalScroll() float32 {
	return s.verticalScrollDelta
}

func (s *Store) HorizontalScroll() float32 {
	return s.horizontalScrollDelta
}

func (s *Store) onKeyPressed(event *core.KeyEvent) {
	s.prevKeyboardState[event.Code] = s.keyboardState[event.Code]
	s.keyboardState[event.Code] = true
}

func (s *Store) onKeyReleased(event *core.KeyEvent) {
	s.prevKeyboardState[event.Code] = s.keyboardState[event.Code]
	s.keyboardState[event.Code] = false
}

func (s *Store) onMouseButtonPressed(event *core.MouseButtonEvent) {
	s.prevMouseButtonState[event.Button] = s.mouseButtonState[event.Button]
	s.mouseButtonState[event.Button] = true
}

func (s *Store) onMouseButtonReleased(event *core.MouseButtonEvent) {
	s.prevMouseButtonState[event.Button] = s.mouseButtonState[event.Button]
	s.mouseButtonState[event.Button] = false
}

func (s *Store) onMouseMoved(event *core.MouseMoveEvent) {
	s.prevMousePosition = s.mousePosition
	s.mousePosition = &sfml.Vector2f{X: float32(event.X), Y: float32(event.Y)}
}

func (s *Store) onMouseWheelScrolled(event *core.MouseWheelScrollEvent) {
	if event.Wheel == sfml.MouseVerticalWheel {
		s.verticalScrollDelta += event.Delta
	} else if event.Wheel == sfml.MouseHorizontalWheel {
		s.horizontalScrollDelta += event.Delta
	}
}
