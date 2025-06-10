package scene

import (
	"fmt"
	"github.com/saffronjam/cimgui-go/imgui"
	"github.com/saffronjam/go-sfml/public/sfml"
	"go-saffron/pkg/core"
	"go-saffron/pkg/input"
)

type Camera struct {
	Reset core.SubscriberList[any]

	Enabled bool

	transform         *sfml.Transform
	positionTransform *sfml.Transform
	rotationTransform *sfml.Transform
	zoomTransform     *sfml.Transform

	position     *sfml.Vector2f
	rotation     float32
	rps          float32 // Rotations per second
	zoom         *sfml.Vector2f
	viewportSize *sfml.Vector2f

	follow *sfml.Vector2f // Optional pointer to a vector to follow
}

func NewCamera() *Camera {
	camera := &Camera{
		Enabled: true,

		transform:         identityMatrix(),
		positionTransform: identityMatrix(),
		rotationTransform: identityMatrix(),
		zoomTransform:     identityMatrix(),

		position:     &sfml.Vector2f{X: 0.0, Y: 0.0},
		rotation:     0.0,
		rps:          0.2,
		zoom:         &sfml.Vector2f{X: 1.0, Y: 1.0},
		viewportSize: &sfml.Vector2f{X: 100.0, Y: 100.0},
	}

	camera.UpdateTransform()

	return camera
}

func (c *Camera) Update() {
	if !c.Enabled {
		return
	}

	dt := core.GlobalClock.Delta()

	if c.follow != nil {
		c.SetCenter(c.follow)
	} else {
		if input.Input.IsMouseButtonDown(sfml.MouseLeft) && input.Input.IsMouseButtonDown(sfml.MouseRight) {
			delta := input.Input.MouseSwipe()
			if delta.LengthSquared() > 0.0 {
				delta = c.rotationTransform.Inverse().TransformPoint(*delta)
				delta = c.zoomTransform.Inverse().TransformPoint(*delta)
				delta = delta.MultiplyScalar(-1.0)
				c.ApplyMovement(delta)
			}
		}
	}

	c.ApplyZoom((input.Input.VerticalScroll() / 100.0) + 1.0)

	var angle float32

	if input.Input.IsKeyDown(sfml.KeyQ) {
		angle += c.rps * dt
	}

	if input.Input.IsKeyDown(sfml.KeyE) {
		angle -= c.rps * 360.0 * dt
	}
	c.ApplyRotation(angle)

	if input.Input.IsKeyPressed(sfml.KeyR) {
		c.ResetTransformation()
	}
}

func (c *Camera) RenderUI() {
	imgui.Begin("Camera")
	imgui.Text(fmt.Sprintf("Position: (%.2f, %.2f)", c.position.X, c.position.Y))
	imgui.Text(fmt.Sprintf("Zoom: (%.2f, %.2f)", c.zoom.X, c.zoom.Y))
	imgui.Text(fmt.Sprintf("Rotation: %.2f", c.rotation))
	imgui.Text(fmt.Sprintf("Rotation Speed: %.2f", c.rps))
	if c.follow != nil {
		imgui.Text(fmt.Sprintf("Following: (%.2f, %.2f)", c.follow.X, c.follow.Y))
	} else {
		imgui.Text("Not following any point")
	}
	imgui.End()
}

func (c *Camera) ApplyMovement(offset *sfml.Vector2f) {
	c.SetCenter(c.position.Add(offset))
}

func (c *Camera) ApplyZoom(factor float32) {
	c.zoom = c.zoom.MultiplyScalar(factor)
	c.zoomTransform.Scale(factor, factor)
	c.UpdateTransform()
}

func (c *Camera) ApplyRotation(angle float32) {
	c.SetRotation(c.rotation + angle)
}

func (c *Camera) SetCenter(center *sfml.Vector2f) {
	c.position = center
	c.positionTransform = identityMatrix()
	c.positionTransform.Translate(center.X, center.Y)
	c.UpdateTransform()
}

func (c *Camera) SetZoom(zoom float32) {
	if zoom != 0.0 {
		c.zoom = &sfml.Vector2f{X: zoom, Y: zoom}
		c.zoomTransform = identityMatrix()
		c.zoomTransform.Scale(zoom, zoom)
		c.UpdateTransform()
	}
}

func (c *Camera) SetRotation(angle float32) {
	c.rotation = angle
	c.rotationTransform = sfml.TransformFromMatrix(1.0, 0.0, 0.0, 1.0, 0.0, 0.0, angle, 0.0, 1.0)
	c.UpdateTransform()
}

func (c *Camera) Follow(follow *sfml.Vector2f) {
	c.follow = follow
}

func (c *Camera) Unfollow() {
	c.follow = nil
}

func (c *Camera) ScreenToWorldPoint(point *sfml.Vector2f) *sfml.Vector2f {
	return c.transform.Inverse().TransformPoint(*point)
}

func (c *Camera) ScreenToWorldRect(rect *sfml.FloatRect) *sfml.FloatRect {
	return c.transform.Inverse().TransformRect(*rect)
}

func (c *Camera) WorldToScreenPoint(point *sfml.Vector2f) *sfml.Vector2f {
	return c.transform.TransformPoint(*point)
}

func (c *Camera) WorldToScreenRect(rect *sfml.FloatRect) *sfml.FloatRect {
	return c.transform.TransformRect(*rect)
}

func (c *Camera) TransformMatrix() *sfml.Transform {
	return c.transform
}

func (c *Camera) SetRotationSpeed(rps float32) {
	c.rps = rps
}

func (c *Camera) SetTransform(transform *sfml.Transform) {
	c.transform = transform
}

func (c *Camera) Viewport() (*sfml.Vector2f, *sfml.Vector2f) {
	vs := c.ViewportSize()
	screenRect := sfml.FloatRect{Left: 0.0, Top: 0.0, Width: vs.X, Height: vs.Y}
	tl := &sfml.Vector2f{X: screenRect.Left, Y: screenRect.Top}
	br := &sfml.Vector2f{X: screenRect.Left + screenRect.Width, Y: screenRect.Top + screenRect.Height}

	return c.transform.Inverse().TransformPoint(*tl), c.transform.Inverse().TransformPoint(*br)
}

func (c *Camera) Offset() *sfml.Vector2f {
	return c.ViewportSize().DivideScalar(2.0)
}

func (c *Camera) UpdateTransform() {
	c.transform = identityMatrix()
	c.transform.Translate(c.Offset().X, c.Offset().Y)
	c.transform.Scale(c.zoom.X, c.zoom.Y)
	c.transform.Rotate(c.rotation)
	c.transform.Translate(-c.position.X, -c.position.Y)
}

func (c *Camera) ResetTransformation() {
	c.SetCenter(&sfml.Vector2f{X: 0.0, Y: 0.0})
	c.SetRotation(0.0)
	c.SetZoom(1.0)

	c.positionTransform = identityMatrix()
	c.rotationTransform = identityMatrix()
	c.zoomTransform = identityMatrix()

	c.Reset.Trigger(struct{}{})
}

func (c *Camera) ViewportSize() *sfml.Vector2f {
	return c.viewportSize
}

func (c *Camera) SetViewportSize(size *sfml.Vector2f) {
	c.viewportSize = size
}

func identityMatrix() *sfml.Transform {
	return sfml.TransformFromMatrix(1.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 1.0)
}
