package sys

type Window struct {
	WindowProps
}

type WindowProps struct {
	Height, Width uint
	Title         string
	Antialiasing  uint
	BitsPerPixel  uint
}

func DefaultProps() *WindowProps {
	return &WindowProps{Height: 800, Width: 600, Title: "Default Window", Antialiasing: 0, BitsPerPixel: 32}
}

func CreateWindow(existingWindowHandle uintptr, props *WindowProps) (*Window, error) {
	//makeErr := func(err error) error {
	//	return fmt.Errorf("failed to create window: %v", err)
	//}
	return nil, nil
}

func (w *Window) Close() {
}

func (w *Window) PollEvents(eventStore *EventStore) {

}

func (w *Window) Clear() {
}

func (w *Window) Display() {
}
