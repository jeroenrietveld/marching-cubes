package window

import (
	"fmt"
	"log"

	"marching_cubes/glutils"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

type Container struct {
	Window *glfw.Window
	Scene  Scene
}

func NewContainer() Container {
	window, err := createWindow()
	if err != nil {
		log.Fatalln("Could not create window: ", err)
	}

	mesh := glutils.NewMesh([]glutils.Vertex{
		glutils.NewVertex(-0.5, -0.5, 0.0),
		glutils.NewVertex(0.5, -0.5, 0.0),
		glutils.NewVertex(0.0, 0.5, 0.0),
	})

	scene := NewScene()
	scene.AddMesh(mesh)

	return Container{
		Window: window,
		Scene:  scene,
	}
}

func createWindow() (*glfw.Window, error) {
	glfw.WindowHint(glfw.Resizable, glfw.True)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, gl.TRUE)

	window, err := glfw.CreateWindow(800, 600, "Golang OpenGL", nil, nil)
	if err != nil {
		return nil, err
	}
	window.MakeContextCurrent()

	if err := gl.Init(); err != nil {
		return nil, err
	}

	window.SetFramebufferSizeCallback(resizeCallback)
	window.SetKeyCallback(keyCallBack)

	version := gl.GoStr(gl.GetString(gl.VERSION))
	glsl := gl.GoStr(gl.GetString(gl.SHADING_LANGUAGE_VERSION))
	fmt.Println("OpenGL version", version, glsl)

	width, height := window.GetFramebufferSize()
	gl.Viewport(0, 0, int32(width), int32(height))

	return window, nil
}

func resizeCallback(w *glfw.Window, width int, height int) {
	gl.Viewport(0, 0, int32(width), int32(height))
}

func keyCallBack(w *glfw.Window, k glfw.Key, s int, a glfw.Action, mk glfw.ModifierKey) {
	if a == glfw.Press {
		if k == glfw.KeyEscape {
			w.SetShouldClose(true)
		}
	}
}
