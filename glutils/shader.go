package glutils

import (
	"fmt"
	"log"
	"strings"

	"marching_cubes/helper"

	"github.com/go-gl/gl/v4.1-core/gl"
)

type Shader struct {
	Program uint32
}

func (s Shader) Use() {
	gl.UseProgram(s.Program)
}

// func (s Shader) SetBool(name string, value bool) {
// 	gl.Uniform1i(gl.GetUniformLocation(s.Program, name), int(value))
// }

// func (s Shader) SetInt(name string, value int) {
// 	gl.Uniform1i(gl.GetUniformLocation(s.Program, name), value)
// }

// func (s Shader) SetFloat(name string, value float) {
// 	gl.Uniform1f(gl.GetUniformLocation(s.Program, name), value)
// }

func NewShader(vertexShaderPath, fragmentShaderPath string) Shader {
	var vertexShaderSource, fragmentShaderSource string
	vertexShaderSource = helper.ReadFile(vertexShaderPath)
	fragmentShaderSource = helper.ReadFile(fragmentShaderPath)

	program, err := createProgram(vertexShaderSource, fragmentShaderSource)
	if err != nil {
		log.Fatalf("Program creation failed %v: ", err)
	}

	return Shader{program}
}

func createProgram(vertexShaderSource, fragmentShaderSource string) (uint32, error) {
	vertexShader, err := compileShader(vertexShaderSource, gl.VERTEX_SHADER)
	if err != nil {
		log.Fatalf("Vertex compilation failed %v: ", err)
	}

	fragmentShader, err := compileShader(fragmentShaderSource, gl.FRAGMENT_SHADER)
	if err != nil {
		log.Fatalf("Fragment compilation failed %v: ", err)
	}

	program := gl.CreateProgram()

	gl.AttachShader(program, vertexShader)
	gl.AttachShader(program, fragmentShader)
	gl.LinkProgram(program)

	var status int32
	gl.GetProgramiv(program, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(program, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(program, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to link program: %v", log)
	}

	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)

	return program, nil
}

func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(source)
	defer free()

	gl.ShaderSource(shader, 1, csources, nil)
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)

	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return shader, nil
}
