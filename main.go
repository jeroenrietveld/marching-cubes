package main

import (
	"log"

	"marching_cubes/window"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

func main() {
	initGL()

	var container window.Container
	container = window.NewContainer()

	for !container.Window.ShouldClose() {
		gl.ClearColor(0.2, 0.3, 0.3, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT)

		container.Scene.Draw()

		container.Window.SwapBuffers()
		glfw.PollEvents()
	}

	defer glfw.Terminate()
}

func initGL() {
	if err := glfw.Init(); err != nil {
		log.Fatalln("Failed to initialize glfw: ", err)
	}
}
