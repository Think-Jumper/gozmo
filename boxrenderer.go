package gozmo

import (
	_ "fmt"

	"github.com/go-gl/mathgl/mgl32"
)

// BoxRenderer is an alternative renderer used for simple solid-color boxes.
type BoxRenderer struct {
	mesh *Mesh

	Width  float32
	Height float32
}

func (box *BoxRenderer) Start(gameObject *GameObject) {}

// Boxes are created at setup.
func NewBoxRenderer(width, height float32) *BoxRenderer {
	box := BoxRenderer{Width: width, Height: height}

	if shader == -1 {
		shader = int32(GLShader())
	}

	mesh := Mesh{}
	mesh.abid = GLNewArray()
	mesh.vbid = GLNewBuffer()

	mesh.vertices = []float32{-1, -1,
		-1, 1,
		1, -1,
		1, -1,
		1, 1,
		-1, 1}

	GLBufferData(0, mesh.vbid, mesh.vertices)

	mesh.mulColor = mgl32.Vec4{0, 0, 0, 0}

	box.mesh = &mesh

	return &box
}

func (box *BoxRenderer) Update(gameObject *GameObject) {

	model := mgl32.Translate3D(gameObject.Position[0], gameObject.Position[1], 0)

	model = model.Mul4(mgl32.Scale3D(gameObject.Scale[0], gameObject.Scale[1], 1))

	model = model.Mul4(mgl32.HomogRotate3DZ(gameObject.Rotation))

	view := Engine.Window.View.Mul4(model)

	ortho := Engine.Window.Projection.Mul4(view)

	GLDraw(box.mesh, uint32(shader), box.Width/2, box.Height/2, -1, 0, 0, 0, 0, ortho)
}

func (box *BoxRenderer) SetAttr(attr string, value interface{}) error {
	switch attr {
	case "red", "r", "R":
		box.mesh.addColor[0], _ = CastFloat32(value)
	case "green", "g", "G":
		box.mesh.addColor[1], _ = CastFloat32(value)
	case "blue", "b", "blu", "B":
		box.mesh.addColor[2], _ = CastFloat32(value)
	case "alpha", "a", "A":
		box.mesh.addColor[3], _ = CastFloat32(value)
	}
	return nil
}

func (box *BoxRenderer) GetAttr(attr string) (interface{}, error) {
	return nil, nil
}

func (box *BoxRenderer) GetType() string {
	return "BoxRenderer"
}

func initBoxRenderer(args []interface{}) Component {
	var width float32 = 1
	var height float32 = 1
	if len(args) > 0 {
		width, _ = CastFloat32(args[0])
	}
	if len(args) > 1 {
		height, _ = CastFloat32(args[1])
	}
	return NewBoxRenderer(width, height)
}

func init() {
	RegisterComponent("BoxRenderer", initBoxRenderer)
}
