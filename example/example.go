package main

import (
	"github.com/saffronjam/cimgui-go/imgui"
	"github.com/saffronjam/go-saffron/pkg/saffron"
	"github.com/saffronjam/go-sfml/public/sfml"
	"go.uber.org/zap/zapcore"
	"runtime"
)

func init() { runtime.LockOSThread() }

func main() {
	saffronApp, err := saffron.NewApp(&saffron.Config{
		WindowProps: &saffron.WindowProps{
			Width:      1920 - 100,
			Height:     1080 - 100,
			Title:      "go-saffron Example",
			Fullscreen: false,
		},
	})
	if err != nil {
		panic(err)
	}

	saffron.SetMainApp(saffronApp)

	err = saffronApp.Run(NewSaffronClient())
	if err != nil {
		panic(err)
	}
}

type SaffronClient struct {
	App          *saffron.App
	RenderTarget *saffron.ControllableRenderTexture
	Scene        *saffron.Scene
	Camera       *saffron.Camera
	ViewportPane *saffron.ViewportPane
	Log          *saffron.LogView

	WorldSpaceDrawable  *sfml.RectangleShape
	ScreenSpaceDrawable *sfml.RectangleShape
}

func NewSaffronClient() *SaffronClient {
	target := saffron.NewControllableRenderTexture(1600, 900, false)
	camera := saffron.NewCamera()
	viewportPane := saffron.NewViewportPane("Viewport", target)

	viewportPane.Resized.Subscribe(func(size *sfml.Vector2f) {
		target.Resize(int(size.X), int(size.Y))
		camera.SetViewportSize(size)
	})

	guiLog := saffron.NewLog()
	saffron.OnLog.Subscribe(func(msg zapcore.Entry) {
		guiLog.AddEntry(msg)
	})

	rect1 := sfml.NewRectangleShape()
	rect1.SetSize(sfml.Vector2f{X: 100, Y: 100})
	rect1.SetFillColor(sfml.Color{R: 255, G: 0, B: 0, A: 255})

	rect2 := sfml.NewRectangleShape()
	rect2.SetSize(sfml.Vector2f{X: 100, Y: 100})
	rect2.SetFillColor(sfml.Color{R: 0, G: 255, B: 0, A: 255})

	return &SaffronClient{
		App:          saffron.MainApp,
		RenderTarget: target,
		Scene:        saffron.NewScene(target, camera),
		Camera:       camera,
		ViewportPane: viewportPane,
		Log:          guiLog,

		WorldSpaceDrawable:  rect1,
		ScreenSpaceDrawable: rect2,
	}
}

func (c *SaffronClient) Setup() error {
	return nil
}

func (c *SaffronClient) Update() error {
	saffron.BeginDockSpace()
	c.RenderTarget.Clear(sfml.Color{R: 0, G: 0, B: 0, A: 255})

	imgui.ShowDemoWindow()

	if saffron.Input.IsKeyPressed(sfml.KeyNum1) {
		saffron.Infoln("info log message")
	}
	if saffron.Input.IsKeyPressed(sfml.KeyNum2) {
		saffron.Debugln("debug log message")
	}
	if saffron.Input.IsKeyPressed(sfml.KeyNum3) {
		saffron.Warnln("warning log message")
	}
	if saffron.Input.IsKeyPressed(sfml.KeyNum4) {
		saffron.Errorln("error log message")
	}

	c.Camera.Update()

	c.Scene.SubmitRectangleShape(c.WorldSpaceDrawable, nil)
	c.Scene.PushOptions(saffron.ScreenSpaceRendering)
	c.Scene.SubmitRectangleShape(c.WorldSpaceDrawable, nil)
	c.Scene.PopOptions()

	c.Camera.RenderUI()
	c.Log.RenderUI()
	c.ViewportPane.RenderUI()
	c.App.RenderUI()
	saffron.EndDockSpace()

	c.RenderTarget.Display()
	return nil
}
