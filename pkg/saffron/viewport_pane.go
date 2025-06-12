package saffron

import (
	"github.com/saffronjam/cimgui-go/imgui"
	"github.com/saffronjam/go-sfml/public/sfml"
)

type ViewportPane struct {
	WindowTitle     string
	Target          *ControllableRenderTexture
	TopLeft         *sfml.Vector2f
	BottomRight     *sfml.Vector2f
	Hovered         bool
	Focused         bool
	FallbackTexture *sfml.Texture

	// Subscribers for rendering and resizing events
	Rendered SubscriberList[any]
	Resized  SubscriberList[*sfml.Vector2f]

	DockID uint
}

func NewViewportPane(windowTitle string, target *ControllableRenderTexture) *ViewportPane {
	return &ViewportPane{
		WindowTitle: windowTitle,
		Target:      target,
		TopLeft:     &sfml.Vector2f{X: 0.0, Y: 0.0},
		BottomRight: &sfml.Vector2f{X: 100.0, Y: 100.0},
		Hovered:     false,
		Focused:     false,
		DockID:      0,
	}
}

func (vp *ViewportPane) RenderUI() {
	tl := vp.TopLeft
	br := vp.BottomRight

	imgui.PushStyleVarVec2(imgui.StyleVarWindowPadding, imgui.Vec2{X: 0, Y: 0})

	windowName := vp.WindowTitle
	imgui.Begin(windowName)

	vp.DockID = uint(imgui.WindowDockID())

	vp.Hovered = imgui.IsWindowHovered()
	vp.Focused = imgui.IsWindowFocused()

	viewportOffset := imgui.CursorPos()
	minBound := imgui.WindowPos()
	minBound.X += viewportOffset.X
	minBound.Y += viewportOffset.Y

	windowSize := imgui.WindowSize()
	maxBound := imgui.Vec2{
		X: minBound.X + windowSize.X - viewportOffset.X,
		Y: minBound.Y + windowSize.Y - viewportOffset.Y,
	}
	vp.TopLeft = &sfml.Vector2f{X: minBound.X, Y: minBound.Y}
	vp.BottomRight = &sfml.Vector2f{X: maxBound.X, Y: maxBound.Y}

	viewportSize := vp.ViewportSize()

	if vp.Target.Enabled {
		imgui.ImageV(imgui.TextureID(vp.Target.RenderTexture().Texture().NativeHandle()),
			imgui.Vec2{
				X: viewportSize.X,
				Y: viewportSize.Y,
			},
			imgui.Vec2{
				X: 0.0,
				Y: 1.0,
			}, imgui.Vec2{
				X: 1.0,
				Y: 0.0,
			})
	} else {
		if vp.FallbackTexture == nil {
			panic("ViewportPane: FallbackTexture is nil, please set a fallback texture before rendering.")
		}

		imgui.ImageV(imgui.TextureID(vp.FallbackTexture.NativeHandle()),
			imgui.Vec2{
				X: viewportSize.X,
				Y: viewportSize.Y,
			},
			imgui.Vec2{
				X: 0.0,
				Y: 1.0,
			}, imgui.Vec2{
				X: 1.0,
				Y: 0.0,
			})
	}

	color := sfml.Color{R: 255, G: 140, B: 0, A: 180}
	colorUint32 := uint32(color.A)<<24 | uint32(color.B)<<16 | uint32(color.G)<<8 | uint32(color.R)

	imgui.WindowDrawList().AddRectV(imgui.Vec2{X: vp.TopLeft.X, Y: tl.Y}, imgui.Vec2{X: br.X, Y: br.Y},
		colorUint32, 0.0,
		imgui.DrawFlagsRoundCornersAll, 4)

	vp.Rendered.Trigger(struct{}{})

	imgui.End()
	imgui.PopStyleVar()

	uVecViewportSize := sfml.Vector2u{
		X: uint32(viewportSize.X),
		Y: uint32(viewportSize.Y),
	}
	renderTargetSize := vp.Target.RenderTexture().Size()

	if !uVecViewportSize.Equals(renderTargetSize) {
		vp.Resized.Trigger(vp.ViewportSize())
	}

}

func (vp *ViewportPane) InViewport(positionNDC *sfml.Vector2f) bool {
	positionNDC.X -= vp.TopLeft.X
	positionNDC.Y -= vp.TopLeft.Y
	return positionNDC.X < vp.BottomRight.X && positionNDC.Y < vp.BottomRight.Y
}

func (vp *ViewportPane) MousePosition(normalized bool) *sfml.Vector2f {
	position := Input.MousePosition()
	position.X -= vp.TopLeft.X
	position.Y -= vp.TopLeft.Y

	if normalized {
		viewportWidth := vp.BottomRight.X - vp.TopLeft.X
		viewportHeight := vp.BottomRight.Y - vp.TopLeft.Y
		return &sfml.Vector2f{
			X: position.X/viewportWidth*2.0 - 1.0,
			Y: (position.Y/viewportHeight*2.0 - 1.0) * -1.0,
		}
	}
	return &sfml.Vector2f{X: position.X, Y: position.Y}
}

func (vp *ViewportPane) ViewportSize() *sfml.Vector2f {
	return &sfml.Vector2f{
		X: vp.BottomRight.X - vp.TopLeft.X,
		Y: vp.BottomRight.Y - vp.TopLeft.Y,
	}
}
