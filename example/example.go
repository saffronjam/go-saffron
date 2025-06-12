package main

import (
	"github.com/saffronjam/cimgui-go/imgui"
	"github.com/saffronjam/go-sfml/public/sfml"
	"go-saffron/pkg/app"
	"go-saffron/pkg/core"
	"go-saffron/pkg/gui"
	"go-saffron/pkg/input"
	"go-saffron/pkg/log"
	"go-saffron/pkg/scene"
	"go.uber.org/zap/zapcore"
	"runtime"
)

func init() { runtime.LockOSThread() }

func main() {
	saffronApp, err := app.NewApp(&app.Config{
		WindowProps: &core.WindowProps{
			Width:      1920 - 100,
			Height:     1080 - 100,
			Title:      "go-saffron Example",
			Fullscreen: false,
		},
	})
	if err != nil {
		panic(err)
	}

	app.SetMainApp(saffronApp)

	err = saffronApp.Run(NewSaffronClient())
	if err != nil {
		panic(err)
	}
}

type Algorithm struct {
	Values []float64
}

type SaffronClient struct {
	App          *app.App
	RenderTarget *core.ControllableRenderTexture
	Scene        *scene.Scene
	Camera       *scene.Camera
	ViewportPane *gui.ViewportPane
	Log          *gui.Log
	Algorithm    Algorithm
}

func NewSaffronClient() *SaffronClient {
	target := core.NewControllableRenderTexture(1600, 900, false)
	camera := scene.NewCamera()
	viewportPane := gui.NewViewportPane("Viewport", target)

	viewportPane.Resized.Subscribe(func(size *sfml.Vector2f) {
		target.Resize(int(size.X), int(size.Y))
		camera.SetViewportSize(size)
	})

	guiLog := gui.NewLog()
	log.OnLog.Subscribe(func(msg zapcore.Entry) {
		guiLog.AddEntry(msg)
	})

	return &SaffronClient{
		App:          app.MainApp,
		RenderTarget: target,
		Scene:        scene.NewScene(target, camera),
		Camera:       camera,
		ViewportPane: viewportPane,
		Log:          guiLog,
	}
}

func (c *SaffronClient) Setup() error {
	return nil
}

func (c *SaffronClient) Update() error {
	gui.BeginDockSpace()
	c.RenderTarget.Clear(sfml.Color{R: 0, G: 0, B: 0, A: 255})

	imgui.ShowDemoWindow()

	if input.Input.IsKeyPressed(sfml.KeyNum1) {
		log.Infoln("info log message")
	}
	if input.Input.IsKeyPressed(sfml.KeyNum2) {
		log.Debugln("debug log message")
	}
	if input.Input.IsKeyPressed(sfml.KeyNum3) {
		log.Warnln("warning log message")
	}
	if input.Input.IsKeyPressed(sfml.KeyNum4) {
		log.Errorln("error log message")
	}

	c.Camera.Update()
	c.Camera.RenderUI()
	c.Log.RenderUI()
	c.ViewportPane.RenderUI()
	c.App.RenderUI()
	gui.EndDockSpace()

	c.RenderTarget.Display()
	return nil
}
