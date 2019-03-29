package main

import (
	"github.com/go-gl/glfw/v3.2/glfw"
)

type Input struct {
	mouseX, mouseY         float64
	deltaX, deltaY         float64
	keyA, keyW, keyS, keyD bool
	mouseClick             bool
}

var input Input

func keyCallBack(w *glfw.Window, k glfw.Key, st int, a glfw.Action, mk glfw.ModifierKey) {
	if a == glfw.Press && k == glfw.KeyA {
		input.keyA = true
	}
	if a == glfw.Release && k == glfw.KeyA {
		input.keyA = false
	}
	if a == glfw.Press && k == glfw.KeyW {
		input.keyW = true
	}
	if a == glfw.Release && k == glfw.KeyW {
		input.keyW = false
	}
	if a == glfw.Press && k == glfw.KeyS {
		input.keyS = true
	}
	if a == glfw.Release && k == glfw.KeyS {
		input.keyS = false
	}
	if a == glfw.Press && k == glfw.KeyD {
		input.keyD = true
	}
	if a == glfw.Release && k == glfw.KeyD {
		input.keyD = false
	}
}

func cursorPosCallBack(w *glfw.Window, xPos, yPos float64) {
	input.deltaX = input.mouseX - xPos
	input.deltaY = input.mouseY - yPos
	input.mouseX = xPos
	input.mouseY = yPos
	if input.mouseClick {
		rotCamera(input.deltaX, input.deltaY)
	}
}

func mouseCallBack(w *glfw.Window, button glfw.MouseButton, action glfw.Action, mod glfw.ModifierKey) {
	// if button == glfw.MouseButton && action == glfw.Press {
	if action == glfw.Press {
		input.mouseClick = true
	}
	if action == glfw.Release {
		input.mouseClick = false
	}
}
