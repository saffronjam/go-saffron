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

	LoadFont("roboto", "assets/fonts/Roboto-Regular.ttf", 16)

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

func BeginDockSpace() {
	viewport := imgui.MainViewport()
	imgui.SetNextWindowPos(viewport.WorkPos())
	imgui.SetNextWindowSize(viewport.WorkSize())
	imgui.SetNextWindowViewport(viewport.ID())
	hostWindowFlags := imgui.WindowFlagsNoTitleBar | imgui.WindowFlagsNoCollapse | imgui.WindowFlagsNoResize |
		imgui.WindowFlagsNoMove | imgui.WindowFlagsNoDocking | imgui.WindowFlagsNoBringToFrontOnFocus |
		imgui.WindowFlagsNoNavFocus | imgui.WindowFlagsNoBackground | imgui.WindowFlagsMenuBar
	imgui.PushStyleVarFloat(imgui.StyleVarWindowRounding, 0.0)
	imgui.PushStyleVarFloat(imgui.StyleVarWindowBorderSize, 0.0)
	imgui.PushStyleVarVec2(imgui.StyleVarWindowPadding, imgui.Vec2{X: 0.0, Y: 0.0})
	imgui.BeginV("DockSpaceViewport", nil, hostWindowFlags)
	dockerSpaceId := imgui.IDStr("DockSpace")
	imgui.DockSpaceV(dockerSpaceId, imgui.Vec2{X: 0.0, Y: 0.0}, imgui.DockNodeFlagsNone, imgui.NewEmptyWindowClass())
}

func EndDockSpace() {
	imgui.End()
	imgui.PopStyleVarV(3) // Pop the three style vars we pushed in BeginDockSpace
}
