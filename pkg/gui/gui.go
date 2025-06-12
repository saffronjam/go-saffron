package gui

import (
	"github.com/saffronjam/cimgui-go/backend/sfml_backend"
	"github.com/saffronjam/cimgui-go/imgui"
	"go-saffron/pkg/core"
	"unsafe"
)

var (
	backend *sfml_backend.SfmlBackend
)

func Init(window *core.Window, loadDefaultFont bool) error {
	backend = sfml_backend.NewSfmlBackend()

	if err := backend.Init(unsafe.Pointer(window.SfmlHandle().ToC()), loadDefaultFont); err != nil {
		return err
	}

	currentConfigFlags := imgui.CurrentIO().ConfigFlags()
	currentConfigFlags |= imgui.ConfigFlagsDockingEnable
	imgui.CurrentIO().SetConfigFlags(currentConfigFlags)

	fontSizes := []float32{
		8.0, 9.0, 10.0, 11.0, 12.0, 13.0, 14.0, 15.0,
		16.0, 17.0, 18.0, 19.0, 20.0, 21.0, 22.0,
		23.0, 24.0, 25.0, 26.0, 27.0, 28.0, 29.0,
	}

	for _, size := range fontSizes {
		LoadFont("roboto", "assets/fonts/Roboto-Regular.ttf", size)
		LoadFont("roboto-mono", "assets/fonts/Roboto-Mono.ttf", size)
	}

	err := backend.UpdateFontTexture()
	if err != nil {
		return err
	}

	return nil
}

func ProcessEvent(window *core.Window, event core.Event) {
	cEvent := event.SfmlHandle().BaseToC()
	backend.ProcessEvent(unsafe.Pointer(window.SfmlHandle().ToC()), unsafe.Pointer(&cEvent))
}

func Update(window *core.Window) {
	dt := core.GlobalClock.DeltaDuration()
	backend.NewFrame(unsafe.Pointer(window.SfmlHandle().ToC()), dt)
}

func Render(window *core.Window) {
	backend.Render(unsafe.Pointer(window.SfmlHandle().ToC()))
}
