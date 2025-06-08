package gui

import (
	"github.com/saffronjam/cimgui-go/backend/sfml_backend"
	"time"
	"unsafe"
)

var (
	backend *sfml_backend.SfmlBackend
)

func Init(sfWindow uintptr, loadDefaultFont bool) error {
	backend = sfml_backend.NewSfmlBackend()

	if err := backend.Init(unsafe.Pointer(sfWindow), loadDefaultFont); err != nil {
		return err
	}

	return nil
}

func ProcessEvent(sfWindow uintptr, event uintptr) {
	backend.ProcessEvent(unsafe.Pointer(sfWindow), unsafe.Pointer(event))
}

func Update(sfWindow uintptr, dt time.Duration) {
	backend.NewFrame(unsafe.Pointer(sfWindow), dt)
}

func Render(sfWindow uintptr) {
	backend.Render(unsafe.Pointer(sfWindow))
}
