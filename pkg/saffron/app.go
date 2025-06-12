package saffron

import (
	"github.com/saffronjam/cimgui-go/imgui"
)

var MainApp *App

func SetMainApp(app *App) {
	MainApp = app
}

type Config struct {
	WindowProps *WindowProps
}

type App struct {
	Config      *Config
	EventStore  *EventStore
	Window      *Window
	Input       *Store
	ClientScene *Scene
	Clock       *Clock
	MenuBar     *MenuBar
}

type Client interface {
	Setup() error
	Update() error
}

func NewApp(config *Config) (*App, error) {
	err := SetupLogger()
	if err != nil {
		return nil, err
	}

	eventStore := NewEventStore()
	window, err := NewWindow(config.WindowProps)
	if err != nil {
		return nil, err
	}

	clock := NewClock()
	SetGlobalClock(clock)

	menuBar := NewMenuBar()
	menuBar.AddMenu("File", func() {
		if imgui.MenuItemBoolV("Fullscreen", "Alt+Enter", window.Fullscreen, true) {
			window.SetFullscreen(!window.Fullscreen)
		}
	})

	app := &App{
		Config:     config,
		EventStore: eventStore,
		Window:     window,
		Clock:      clock,
		MenuBar:    menuBar,
	}

	eventStore.RegisterProducer(window)

	err = Init(window, true)
	if err != nil {
		return nil, err
	}

	eventStore.RegisterHandlerByTags(func(e any) {
		ProcessEvent(window, e.(Event))
	}, "sfml")

	app.Input = NewInput(eventStore)
	SetGlobalInput(app.Input)

	return app, nil
}

func (app *App) Run(client Client) error {
	app.EventStore.RegisterHandler(func(e any) { app.Window.Close() }, EventClosed)

	SetBessDarkColors()
	err := client.Setup()
	if err != nil {
		Fatalln("Failed to setup client:", err)
	}

	for {
		app.Clock.Tick()
		app.EventStore.ProcessEvents()
		if !app.Window.IsOpen() {
			println("Exiting application")
			break
		}

		app.Window.Clear()
		Update(app.Window)
		PushFont("roboto", 18)
		err = client.Update()
		if err != nil {
			Fatalln("Failed to update client:", err)
		}
		app.Input.PostUpdate()

		PopFont()
		Render(app.Window)
		app.Window.Display()

	}

	return nil
}

func (app *App) RenderUI() {
	app.MenuBar.RenderUI()
}
