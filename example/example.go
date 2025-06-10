package main

import (
	"github.com/saffronjam/go-sfml/public/sfml"
	"go-saffron/pkg/app"
	"go-saffron/pkg/core"
	"go-saffron/pkg/gui"
	"go-saffron/pkg/scene"
	"math"
	"runtime"
)

func init() { runtime.LockOSThread() }

func main() {
	saffronApp, err := app.NewApp(&app.Config{
		WindowProps: &core.WindowProps{
			Width:  1600,
			Height: 900,
			Title:  "go-saffron Example",
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

type SaffronClient struct {
	RenderTarget *core.ControllableRenderTexture
	Scene        *scene.Scene
	Camera       *scene.Camera
	ViewportPane *gui.ViewportPane

	testCs *sfml.CircleShape
}

func NewSaffronClient() *SaffronClient {
	target := core.NewControllableRenderTexture(1600, 900, false)
	camera := scene.NewCamera()
	viewportPane := gui.NewViewportPane("Viewport", target)

	viewportPane.Resized.Subscribe(func(size *sfml.Vector2f) {
		target.Resize(int(size.X), int(size.Y))
		camera.SetViewportSize(size)
	})

	return &SaffronClient{
		RenderTarget: target,
		Scene:        scene.NewScene(target, camera),
		Camera:       camera,
		ViewportPane: viewportPane,
		testCs:       sfml.NewCircleShape(),
	}
}

func (c *SaffronClient) Setup() error {

	gui.SetBessDarkColors()
	c.testCs.SetRadius(50)
	return nil
}

func (c *SaffronClient) Update() error {
	gui.BeginDockSpace()
	c.RenderTarget.Clear(sfml.Color{R: 0, G: 0, B: 0, A: 255})

	c.Scene.SubmitCircleShape(c.testCs, nil)

	c.Camera.Update()
	c.Camera.RenderUI()
	c.ViewportPane.RenderUI()

	sinceStart := app.MainApp.Clock.SinceStart()
	sinCosVec := sfml.Vector2f{
		X: float32(math.Cos(float64(sinceStart)*2*math.Pi)) * 3,
		Y: float32(math.Sin(float64(sinceStart)*2*math.Pi)) * 3,
	}

	c.testCs.SetPosition(sinCosVec)

	c.RenderTarget.Display()
	gui.EndDockSpace()
	return nil
}
