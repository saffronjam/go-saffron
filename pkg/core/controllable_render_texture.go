package core

import "github.com/saffronjam/go-sfml/public/sfml"

type ControllableRenderTexture struct {
	sfmlRenderTexture *sfml.RenderTexture
	Enabled           bool
}

func NewControllableRenderTexture(width, height int, depthBuffer bool) *ControllableRenderTexture {
	renderTexture := sfml.NewRenderTexture(int32(width), int32(height), depthBuffer)

	return &ControllableRenderTexture{
		sfmlRenderTexture: renderTexture,
		Enabled:           true,
	}
}

func (crt *ControllableRenderTexture) RenderTexture() *sfml.RenderTexture {
	return crt.sfmlRenderTexture
}

func (crt *ControllableRenderTexture) Resize(width, height int) {
	crt.sfmlRenderTexture.Free()
	crt.sfmlRenderTexture = sfml.NewRenderTexture(int32(width), int32(height), false)
}

func (crt *ControllableRenderTexture) Display() {
	if crt.Enabled {
		crt.sfmlRenderTexture.Display()
	}
}

func (crt *ControllableRenderTexture) Clear(color sfml.Color) {
	if crt.Enabled {
		crt.sfmlRenderTexture.Clear(color)
	}
}
