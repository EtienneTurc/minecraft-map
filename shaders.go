package main

import (
	"fmt"
	"strings"

	"github.com/go-gl/gl/v4.1-core/gl"
)

const vertexShader = `
#version 330

uniform mat4 proj;
uniform mat4 view;
uniform mat4 model;
uniform vec3 sun;
uniform ivec3 transChunk;
uniform int sizeChunk;

layout(location = 0) in vec3 vert;
layout(location = 1) in vec2 uv;
layout(location = 2) in vec3 norm;
layout(location = 3) in ivec4 trans;
out float lighting;
out vec2 uvi;
flat out int tex;

void main() {
	gl_Position = proj * view * model * vec4(vert + vec3(trans) + transChunk*sizeChunk, 1);
	lighting = dot(normalize(sun), vec3(model* vec4(normalize(norm),1)));
	uvi = uv;
	tex = int(trans.w);
}
` + "\x00"

// gl_Position = rot * vec4(vert, 1) + vec4(trans,0);

const fragmentShader = `
#version 330

uniform sampler2D grass;
uniform sampler2D stone;

flat in int tex;
in float lighting;
in vec2 uvi;
out vec4 frag_colour;

void main() {
	if (tex == 1) {
		frag_colour = vec4(tex * lighting * texture(stone, uvi).rgb, 1);
	} else {
		frag_colour = vec4(lighting * texture(grass, uvi).rgb, 1);
	}
}
` + "\x00"

// Compile the shader
func compileShader(source string, shaderType uint32) (uint32, error) {
	glSrcs, freeFn := gl.Strs(source)
	defer freeFn()

	shader := gl.CreateShader(shaderType)
	gl.ShaderSource(shader, 1, glSrcs, nil)
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

// CreateShaderProgram load the shader files and create a OpenGL program
func CreateShaderProgram(vertexSource, fragmentSource string) uint32 {
	vertexShader, err := compileShader(vertexSource, gl.VERTEX_SHADER)
	checkPanic(err)

	fragmentShader, err := compileShader(fragmentSource, gl.FRAGMENT_SHADER)
	checkPanic(err)

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

		panic(fmt.Errorf("failed to link program: %v", log))
	}

	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)

	return program
}
