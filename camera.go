package main

import (
	"github.com/go-gl/mathgl/mgl32"
)

type Camera struct {
	eye mgl32.Vec3
	dir mgl32.Vec3
	up  mgl32.Vec3
}

var camera Camera

func rotCamera(deltaX, deltaY float64) {
	sensi := 0.04
	rotY := mgl32.Rotate3DY(mgl32.DegToRad(float32(-deltaX * sensi)))
	camera.dir = rotY.Mul3x1(camera.dir)

	l := camera.up.Cross(camera.dir)
	rotX := mgl32.HomogRotate3D(mgl32.DegToRad(float32(deltaY*sensi)), l)
	camera.dir = (rotX.Mul4x1(camera.dir.Vec4(1))).Vec3()
}

func moveCamera() {
	sensi := float32(0.04)
	l := camera.up.Cross(camera.dir)

	if input.keyA {
		camera.eye = camera.eye.Add(l.Normalize().Mul(sensi))
	}
	if input.keyD {
		camera.eye = camera.eye.Add(l.Normalize().Mul(-sensi))
	}
	if input.keyW {
		camera.eye = camera.eye.Add(camera.dir.Normalize().Mul(sensi))
	}
	if input.keyS {
		camera.eye = camera.eye.Add(camera.dir.Normalize().Mul(-sensi))
	}
}
