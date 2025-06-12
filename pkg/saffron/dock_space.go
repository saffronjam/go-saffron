package saffron

import "github.com/saffronjam/cimgui-go/imgui"

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
