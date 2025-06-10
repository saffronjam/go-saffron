package scene

import "C"
import (
	"github.com/saffronjam/go-sfml/public/sfml"
	"go-saffron/pkg/core"
)

type Scene struct {
	Target    *core.ControllableRenderTexture
	Reference *Camera
}

func NewScene(target *core.ControllableRenderTexture, reference *Camera) *Scene {
	return &Scene{
		Target:    target,
		Reference: reference,
	}
}

func (s *Scene) SubmitCircleShape(object *sfml.CircleShape, states *sfml.RenderStates) {
	if s.Target.Enabled {
		s.Target.RenderTexture().DrawCircleShape(object, s.GenerateRenderStates(states))
	}
}

func (s *Scene) SubmitConvexShape(object *sfml.ConvexShape, states *sfml.RenderStates) {
	if s.Target.Enabled {
		s.Target.RenderTexture().DrawConvexShape(object, s.GenerateRenderStates(states))
	}
}

func (s *Scene) SubmitPrimitives(vertices *sfml.Vertex, vertexCount uint64, primitiveType sfml.PrimitiveType, states *sfml.RenderStates) {
	if s.Target.Enabled {
		s.Target.RenderTexture().DrawPrimitives(vertices, vertexCount, primitiveType, s.GenerateRenderStates(states))
	}
}

func (s *Scene) SubmitRectangleShape(object *sfml.RectangleShape, states *sfml.RenderStates) {
	if s.Target.Enabled {
		s.Target.RenderTexture().DrawRectangleShape(object, s.GenerateRenderStates(states))
	}
}

func (s *Scene) SubmitShape(object *sfml.Shape, states *sfml.RenderStates) {
	if s.Target.Enabled {
		s.Target.RenderTexture().DrawShape(object, s.GenerateRenderStates(states))
	}
}

func (s *Scene) SubmitSprite(object *sfml.Sprite, states *sfml.RenderStates) {
	if s.Target.Enabled {
		s.Target.RenderTexture().DrawSprite(object, s.GenerateRenderStates(states))
	}
}

func (s *Scene) SubmitText(object *sfml.Text, states *sfml.RenderStates) {

	if s.Target.Enabled {
		s.Target.RenderTexture().DrawText(object, s.GenerateRenderStates(states))
	}
}

func (s *Scene) SubmitVertexArray(object *sfml.VertexArray, states *sfml.RenderStates) {
	if states == nil {
		states = sfml.RenderStatesDefault()
	}

	if s.Target.Enabled {
		s.Target.RenderTexture().DrawVertexArray(object, s.GenerateRenderStates(states))
	}
}

func (s *Scene) SubmitVertexBuffer(object *sfml.VertexBuffer, states *sfml.RenderStates) {
	if states == nil {
		states = sfml.RenderStatesDefault()
	}

	if s.Target.Enabled {
		s.Target.RenderTexture().DrawVertexBuffer(object, s.GenerateRenderStates(states))
	}
}

func (s *Scene) SubmitVertexBufferRange(object *sfml.VertexBuffer, firstVertex uint64, vertexCount uint64, states *sfml.RenderStates) {
	if s.Target.Enabled {
		s.Target.RenderTexture().DrawVertexBufferRange(object, firstVertex, vertexCount, s.GenerateRenderStates(states))
	}
}

func (s *Scene) GenerateRenderStates(in *sfml.RenderStates) *sfml.RenderStates {
	states := in
	if states == nil {
		states = sfml.RenderStatesDefault()
	}

	if s.Reference == nil {
		return states
	}

	states.Transform.Combine(s.Reference.TransformMatrix())
	return states
}
