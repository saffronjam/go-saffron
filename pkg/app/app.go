package app

import (
	"github.com/saffronjam/cimgui-go/imgui"
	"go-saffron/pkg/gui"
	"go-saffron/pkg/sys"
	"time"
)

type Config struct {
	WindowProps *sys.WindowProps
}

type App struct {
	Config     *Config
	EventStore *sys.EventStore
	Window     *sys.Window
}

func NewApp(config *Config) (*App, error) {
	eventStore := sys.NewEventStore()
	window, err := sys.NewWindow(config.WindowProps)
	if err != nil {
		return nil, err
	}

	app := &App{
		Config:     config,
		EventStore: eventStore,
		Window:     window,
	}

	eventStore.RegisterProducer(window)

	err = gui.Init(window, true)
	if err != nil {
		return nil, err
	}

	eventStore.RegisterHandlerByTags(func(e any) {
		gui.ProcessEvent(window, e.(sys.Event))
	}, "sfml")

	return app, nil
}

func (a *App) Run() error {
	a.EventStore.RegisterHandler(func(e any) { a.Window.Close() }, sys.EventClosed)
	dt := time.Millisecond * 16 // Target ~60 FPS first iteration
	for {
		before := time.Now()

		a.EventStore.ProcessEvents()
		if !a.Window.IsOpen() {
			println("Exiting application")
			break
		}

		gui.Update(a.Window, dt)

		gui.BeginDockSpace()
		imgui.ShowDemoWindow()

		imgui.Begin("something")
		imgui.Button("a pretty button")
		imgui.End()

		gui.EndDockSpace()

		a.Window.Clear()
		gui.Render(a.Window)
		a.Window.Display()

		after := time.Now()
		dt = after.Sub(before)
	}

	return nil
}
