package glutils

import (
	"unsafe"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type Mesh struct {
	Id       int
	Vertices []Vertex
	VAO      uint32
	VBO, EBO uint32
	shader   Shader
}

func NewMesh(v []Vertex) Mesh {
	var vao, vbo uint32
	gl.GenVertexArrays(1, &vao)
	gl.GenBuffers(1, &vbo)

	var vertexShaderPath = "shader/basic.vs"
	var fragmentShaderPath = "shader/basic.frag"

	shader := NewShader(vertexShaderPath, fragmentShaderPath)

	m := Mesh{
		Vertices: v,
		VAO:      vao,
		VBO:      vbo,
		shader:   shader,
	}
	m.setup()

	return m
}

func (m *Mesh) Draw() {
	m.shader.Use()
	gl.BindVertexArray(m.VAO)
	gl.DrawArrays(gl.TRIANGLES, 0, 3)
	gl.BindVertexArray(0)
}

func (m *Mesh) setup() {
	structSize := int(unsafe.Sizeof(m.Vertices[0]))

	gl.BindVertexArray(m.VAO)
	gl.BindBuffer(gl.ARRAY_BUFFER, m.VBO)
	gl.BufferData(gl.ARRAY_BUFFER, len(m.Vertices)*structSize, gl.Ptr(m.Vertices), gl.STATIC_DRAW)

	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 3*4, gl.PtrOffset(0))
	gl.EnableVertexAttribArray(0)

	gl.BindVertexArray(0)
}

type Vertex struct {
	Position mgl32.Vec3
}

func NewVertex(x, y, z float32) Vertex {
	vertex := Vertex{
		Position: mgl32.Vec3{x, y, z},
	}

	return vertex
}
