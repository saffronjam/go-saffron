package gui

import "github.com/saffronjam/cimgui-go/imgui"

var Fonts = make(map[string]*imgui.Font)

func LoadFont(fontName string, filename string, fontSize float32) {
	font := imgui.CurrentIO().Fonts().AddFontFromFileTTF(filename, fontSize)
	if font == nil {
		panic("Failed to load font: " + filename)
	}
	Fonts[fontName] = font
}

func SetBessDarkColors() {
	style := imgui.CurrentStyle()
	colors := style.Colors()

	// Primary background
	colors[imgui.ColWindowBg] = imgui.Vec4{X: 0.07, Y: 0.07, Z: 0.09, W: 1.00}  // #131318
	colors[imgui.ColMenuBarBg] = imgui.Vec4{X: 0.12, Y: 0.12, Z: 0.15, W: 1.00} // #131318
	colors[imgui.ColPopupBg] = imgui.Vec4{X: 0.18, Y: 0.18, Z: 0.22, W: 1.00}

	// Headers
	colors[imgui.ColHeader] = imgui.Vec4{X: 0.18, Y: 0.18, Z: 0.22, W: 1.00}
	colors[imgui.ColHeaderHovered] = imgui.Vec4{X: 0.30, Y: 0.30, Z: 0.40, W: 1.00}
	colors[imgui.ColHeaderActive] = imgui.Vec4{X: 0.25, Y: 0.25, Z: 0.35, W: 1.00}

	// Buttons
	colors[imgui.ColButton] = imgui.Vec4{X: 0.20, Y: 0.22, Z: 0.27, W: 1.00}
	colors[imgui.ColButtonHovered] = imgui.Vec4{X: 0.30, Y: 0.32, Z: 0.40, W: 1.00}
	colors[imgui.ColButtonActive] = imgui.Vec4{X: 0.35, Y: 0.38, Z: 0.50, W: 1.00}

	// Frame BG
	colors[imgui.ColFrameBg] = imgui.Vec4{X: 0.15, Y: 0.15, Z: 0.18, W: 1.00}
	colors[imgui.ColFrameBgHovered] = imgui.Vec4{X: 0.22, Y: 0.22, Z: 0.27, W: 1.00}
	colors[imgui.ColFrameBgActive] = imgui.Vec4{X: 0.25, Y: 0.25, Z: 0.30, W: 1.00}

	// Tabs
	colors[imgui.ColTab] = imgui.Vec4{X: 0.18, Y: 0.18, Z: 0.22, W: 1.00}
	colors[imgui.ColTabHovered] = imgui.Vec4{X: 0.35, Y: 0.35, Z: 0.50, W: 1.00}
	colors[imgui.ColTabSelected] = imgui.Vec4{X: 0.25, Y: 0.25, Z: 0.38, W: 1.00}
	colors[imgui.ColTabDimmed] = imgui.Vec4{X: 0.13, Y: 0.13, Z: 0.17, W: 1.00}
	colors[imgui.ColTabDimmedSelected] = imgui.Vec4{X: 0.20, Y: 0.20, Z: 0.25, W: 1.00}

	// Title
	colors[imgui.ColTitleBg] = imgui.Vec4{X: 0.12, Y: 0.12, Z: 0.15, W: 1.00}
	colors[imgui.ColTitleBgActive] = imgui.Vec4{X: 0.15, Y: 0.15, Z: 0.20, W: 1.00}
	colors[imgui.ColTitleBgCollapsed] = imgui.Vec4{X: 0.10, Y: 0.10, Z: 0.12, W: 1.00}

	// Borders
	colors[imgui.ColBorder] = imgui.Vec4{X: 0.20, Y: 0.20, Z: 0.25, W: 0.50}
	colors[imgui.ColBorderShadow] = imgui.Vec4{X: 0.00, Y: 0.00, Z: 0.00, W: 0.00}

	// Text
	colors[imgui.ColText] = imgui.Vec4{X: 0.90, Y: 0.90, Z: 0.95, W: 1.00}
	colors[imgui.ColTextDisabled] = imgui.Vec4{X: 0.50, Y: 0.50, Z: 0.55, W: 1.00}

	// Highlights
	colors[imgui.ColCheckMark] = imgui.Vec4{X: 0.50, Y: 0.70, Z: 1.00, W: 1.00}
	colors[imgui.ColSliderGrab] = imgui.Vec4{X: 0.50, Y: 0.70, Z: 1.00, W: 1.00}
	colors[imgui.ColSliderGrabActive] = imgui.Vec4{X: 0.60, Y: 0.80, Z: 1.00, W: 1.00}
	colors[imgui.ColResizeGrip] = imgui.Vec4{X: 0.50, Y: 0.70, Z: 1.00, W: 0.50}
	colors[imgui.ColResizeGripHovered] = imgui.Vec4{X: 0.60, Y: 0.80, Z: 1.00, W: 0.75}
	colors[imgui.ColResizeGripActive] = imgui.Vec4{X: 0.70, Y: 0.90, Z: 1.00, W: 1.00}

	// Scrollbar
	colors[imgui.ColScrollbarBg] = imgui.Vec4{X: 0.10, Y: 0.10, Z: 0.12, W: 1.00}
	colors[imgui.ColScrollbarGrab] = imgui.Vec4{X: 0.30, Y: 0.30, Z: 0.35, W: 1.00}
	colors[imgui.ColScrollbarGrabHovered] = imgui.Vec4{X: 0.40, Y: 0.40, Z: 0.50, W: 1.00}
	colors[imgui.ColScrollbarGrabActive] = imgui.Vec4{X: 0.45, Y: 0.45, Z: 0.55, W: 1.00}

	style.SetColors(&colors)

	// Style tweaks
	style.SetWindowRounding(5.0)
	style.SetFrameRounding(5.0)
	style.SetGrabRounding(5.0)
	style.SetTabRounding(5.0)
	style.SetPopupRounding(5.0)
	style.SetScrollbarRounding(5.0)
	style.SetWindowPadding(imgui.Vec2{X: 10, Y: 10})
	style.SetFramePadding(imgui.Vec2{X: 6, Y: 4})
	style.SetItemSpacing(imgui.Vec2{X: 8, Y: 6})
	style.SetPopupBorderSize(0.0)
}
