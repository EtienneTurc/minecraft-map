package main

import (
	"log"
	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

var window *glfw.Window
var program uint32

func init() {
	runtime.LockOSThread()
}

func start() {
	// Init window
	checkPanic(glfw.Init())

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.Samples, 4)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	// Create a Window
	var err error
	window, err = glfw.CreateWindow(1600, 900, "OpenGL 3D Engine", nil, nil)
	checkPanic(err)

	window.SetKeyCallback(keyCallBack)
	window.SetMouseButtonCallback(mouseCallBack)
	window.SetCursorPosCallback(cursorPosCallBack)

	// Init GL context
	window.MakeContextCurrent()
	checkPanic(gl.Init())

	log.Println("OpenGL version", gl.GoStr(gl.GetString(gl.VERSION)))
}

func run() {
	program = CreateShaderProgram(vertexShader, fragmentShader)
	chunkManager := newChunkManager()
	// obj := newObject(cubeVertices, cubeUVs, cubeNormals)
	gl.Enable(gl.DEPTH_TEST)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
	gl.ClearColor(1, 1, 1, 1)

	gl.UseProgram(program)

	sun := gl.GetUniformLocation(program, gl.Str("sun\x00"))
	gl.Uniform3f(sun, 2, 2, -1)

	projection := mgl32.Perspective(mgl32.DegToRad(45), 16.0/9.0, 0.1, 1000.0)
	projUniform := gl.GetUniformLocation(program, gl.Str("proj\x00"))
	gl.UniformMatrix4fv(projUniform, 1, false, &projection[0])

	camera.eye = mgl32.Vec3{0, 10, -4}
	camera.dir = mgl32.Vec3{0, 0, 1}
	camera.up = mgl32.Vec3{0, 1, 0}

	view := mgl32.LookAtV(camera.eye, camera.dir, camera.up)
	viewUniform := gl.GetUniformLocation(program, gl.Str("view\x00"))
	gl.UniformMatrix4fv(viewUniform, 1, false, &view[0])

	model := mgl32.HomogRotate3DZ(mgl32.DegToRad(0))
	modelUniform := gl.GetUniformLocation(program, gl.Str("model\x00"))
	gl.UniformMatrix4fv(modelUniform, 1, false, &model[0])

	tex1, _ := NewTextureFromFile("grass.png", gl.TEXTURE0)
	text1Uniform := gl.GetUniformLocation(program, gl.Str("grass\x00"))
	gl.Uniform1i(text1Uniform, 0)

	tex2, _ := NewTextureFromFile("stone.png", gl.TEXTURE1)
	text2Uniform := gl.GetUniformLocation(program, gl.Str("stone\x00"))
	gl.Uniform1i(text2Uniform, 1)

	t := 0.0
	for !window.ShouldClose() {
		t++

		// Scene rendering
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		// Draw
		gl.ActiveTexture(gl.TEXTURE0)
		gl.BindTexture(gl.TEXTURE_2D, tex1)

		gl.ActiveTexture(gl.TEXTURE1)
		gl.BindTexture(gl.TEXTURE_2D, tex2)
		chunkManager.draw(camera.eye)

		moveCamera()

		view := mgl32.LookAtV(camera.eye, camera.dir.Add(camera.eye), camera.up)
		viewUniform := gl.GetUniformLocation(program, gl.Str("view\x00"))
		gl.UniformMatrix4fv(viewUniform, 1, false, &view[0])

		window.SwapBuffers()

		// Detect inputs
		glfw.PollEvents()
	}
}

func stop() {
	glfw.Terminate()
}
