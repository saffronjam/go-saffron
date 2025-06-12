package app

import (
	"github.com/saffronjam/cimgui-go/imgui"
	"github.com/saffronjam/go-saffron/pkg/core"
	"github.com/saffronjam/go-saffron/pkg/gui"
	"github.com/saffronjam/go-saffron/pkg/input"
	"github.com/saffronjam/go-saffron/pkg/log"
	"github.com/saffronjam/go-saffron/pkg/scene"
)

var MainApp *App

func SetMainApp(app *App) {
	MainApp = app
}

type Config struct {
	WindowProps *core.WindowProps
}

type App struct {
	Config      *Config
	EventStore  *core.EventStore
	Window      *core.Window
	Input       *input.Store
	ClientScene *scene.Scene
	Clock       *core.Clock
	MenuBar     *gui.MenuBar
}

type Client interface {
	Setup() error
	Update() error
}

func NewApp(config *Config) (*App, error) {
	err := log.SetupLogger()
	if err != nil {
		return nil, err
	}

	eventStore := core.NewEventStore()
	window, err := core.NewWindow(config.WindowProps)
	if err != nil {
		return nil, err
	}

	clock := core.NewClock()
	core.SetGlobalClock(clock)

	menuBar := gui.NewMenuBar()
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

	err = gui.Init(window, true)
	if err != nil {
		return nil, err
	}

	eventStore.RegisterHandlerByTags(func(e any) {
		gui.ProcessEvent(window, e.(core.Event))
	}, "sfml")

	app.Input = input.NewInput(eventStore)
	input.SetGlobalInput(app.Input)

	return app, nil
}

func (app *App) Run(client Client) error {
	app.EventStore.RegisterHandler(func(e any) { app.Window.Close() }, core.EventClosed)

	gui.SetBessDarkColors()
	err := client.Setup()
	if err != nil {
		log.Fatalln("Failed to setup client:", err)
	}

	for {
		app.Clock.Tick()
		app.EventStore.ProcessEvents()
		if !app.Window.IsOpen() {
			println("Exiting application")
			break
		}

		app.Window.Clear()
		gui.Update(app.Window)
		gui.PushFont("roboto", 18)
		err = client.Update()
		if err != nil {
			log.Fatalln("Failed to update client:", err)
		}
		app.Input.PostUpdate()

		gui.PopFont()
		gui.Render(app.Window)
		app.Window.Display()

	}

	return nil
}

func (app *App) RenderUI() {
	app.MenuBar.RenderUI()
}
