package saffron

import (
	"github.com/saffronjam/cimgui-go/backend/sfml_backend"
	"github.com/saffronjam/cimgui-go/imgui"
	"path/filepath"
	"runtime"
	"unsafe"
)

var (
	backend *sfml_backend.SfmlBackend
)

func getAssetPath(relPath string) string {
	_, filename, _, _ := runtime.Caller(1) // gets the current file's path
	return filepath.Join(filepath.Dir(filename), relPath)
}

func Init(window *Window, loadDefaultFont bool) error {
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
		LoadFont("roboto", getAssetPath("../../assets/fonts/Roboto-Regular.ttf"), size)
		LoadFont("roboto-mono", getAssetPath("../../assets/fonts/Roboto-Mono.ttf"), size)
	}

	err := backend.UpdateFontTexture()
	if err != nil {
		return err
	}

	return nil
}

func ProcessEvent(window *Window, event Event) {
	cEvent := event.SfmlHandle().BaseToC()
	backend.ProcessEvent(unsafe.Pointer(window.SfmlHandle().ToC()), unsafe.Pointer(&cEvent))
}

func Update(window *Window) {
	dt := GlobalClock.DeltaDuration()
	backend.NewFrame(unsafe.Pointer(window.SfmlHandle().ToC()), dt)
}

func Render(window *Window) {
	backend.Render(unsafe.Pointer(window.SfmlHandle().ToC()))
}
