package main

import (
	"github.com/saffronjam/go-sfml/public/sfml"
	"go-saffron/pkg/app"
	"go-saffron/pkg/core"
	"go-saffron/pkg/gui"
	"go-saffron/pkg/scene"
	"math"
	"math/rand"
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

	testCs        []*sfml.CircleShape
	basePositions []*sfml.Vector2f
	speed         []float32
}

func NewSaffronClient() *SaffronClient {
	target := core.NewControllableRenderTexture(1600, 900, false)
	camera := scene.NewCamera()
	viewportPane := gui.NewViewportPane("Viewport", target)

	viewportPane.Resized.Subscribe(func(size *sfml.Vector2f) {
		target.Resize(int(size.X), int(size.Y))
		camera.SetViewportSize(size)
	})

	n := 100
	circleShapes := make([]*sfml.CircleShape, n*n)
	basePositions := make([]*sfml.Vector2f, n*n)
	speed := make([]float32, n*n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			circleShapes[i*n+j] = sfml.NewCircleShape()
			circleShapes[i*n+j].SetRadius(10)
			circleShapes[i*n+j].SetFillColor(sfml.Color{R: 255, G: 0, B: 0, A: 255})
			basePositions[i*n+j] = &sfml.Vector2f{X: float32(i * 20), Y: float32(j * 20)}
			speed[i*n+j] = rand.Float32() * 2 // Random speed for each circle
		}
	}

	return &SaffronClient{
		RenderTarget:  target,
		Scene:         scene.NewScene(target, camera),
		Camera:        camera,
		ViewportPane:  viewportPane,
		testCs:        circleShapes,
		basePositions: basePositions,
		speed:         speed,
	}
}

func (c *SaffronClient) Setup() error {
	gui.SetBessDarkColors()
	return nil
}

func (c *SaffronClient) Update() error {
	gui.BeginDockSpace()
	c.RenderTarget.Clear(sfml.Color{R: 0, G: 0, B: 0, A: 255})

	sinceStart := core.GlobalClock.SinceStart()
	for idx, cs := range c.testCs {
		sinCosVec := sfml.Vector2f{
			X: float32(math.Cos(float64(sinceStart)*2*math.Pi*float64(c.speed[idx]))) * 3,
			Y: float32(math.Sin(float64(sinceStart)*2*math.Pi*float64(c.speed[idx]))) * 3,
		}
		cs.SetPosition(sfml.Vector2f{
			X: c.basePositions[idx].X + sinCosVec.X,
			Y: c.basePositions[idx].Y + sinCosVec.Y,
		})
	}
	c.Camera.Update()

	for _, cs := range c.testCs {
		c.Scene.SubmitCircleShape(cs, nil)
	}
	c.Camera.RenderUI()
	c.ViewportPane.RenderUI()
	c.RenderTarget.Display()
	gui.EndDockSpace()
	return nil
}
