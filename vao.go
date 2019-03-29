package main

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

// Object ...
type Object struct {
	VertexCount   int32
	InstanceCount int32

	Ptr uint32

	VertexBuffer   uint32
	UVBuffer       uint32
	NormalsBuffer  uint32
	InstanceBuffer uint32
}

const sizeOfFloat32 = 4

// Returns a vertex array from the vertices provided
func newObject(vertices, uvs, normals []float32) Object {
	// Create VAO buffer
	var obj Object

	obj.VertexCount = int32(len(vertices) / 3)

	gl.GenVertexArrays(1, &obj.Ptr)
	gl.BindVertexArray(obj.Ptr)

	// Create Vertices buffer
	if vertices != nil {
		gl.GenBuffers(1, &obj.VertexBuffer)
		gl.BindBuffer(gl.ARRAY_BUFFER, obj.VertexBuffer)
		gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*sizeOfFloat32, gl.Ptr(vertices), gl.STATIC_DRAW)

		gl.EnableVertexAttribArray(0)
		gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)
	}

	// Create UVs buffer
	if uvs != nil {
		gl.GenBuffers(1, &obj.UVBuffer)
		gl.BindBuffer(gl.ARRAY_BUFFER, obj.UVBuffer)
		gl.BufferData(gl.ARRAY_BUFFER, len(uvs)*sizeOfFloat32, gl.Ptr(uvs), gl.STATIC_DRAW)

		gl.EnableVertexAttribArray(1)
		gl.VertexAttribPointer(1, 2, gl.FLOAT, false, 0, nil)
	}

	// Create Normals buffer
	if normals != nil {
		gl.GenBuffers(1, &obj.NormalsBuffer)
		gl.BindBuffer(gl.ARRAY_BUFFER, obj.NormalsBuffer)
		gl.BufferData(gl.ARRAY_BUFFER, len(normals)*sizeOfFloat32, gl.Ptr(normals), gl.STATIC_DRAW)

		gl.EnableVertexAttribArray(2)
		gl.VertexAttribPointer(2, 3, gl.FLOAT, false, 0, nil)
	}

	// Create Instance buffer
	gl.GenBuffers(1, &obj.InstanceBuffer)
	gl.BindBuffer(gl.ARRAY_BUFFER, obj.InstanceBuffer)
	gl.BufferData(gl.ARRAY_BUFFER, 0, nil, gl.STATIC_DRAW)

	gl.EnableVertexAttribArray(3)
	gl.VertexAttribPointer(3, 4, gl.FLOAT, false, 0, nil)
	gl.VertexAttribDivisor(3, 1)

	// Disable o.Ptr
	gl.BindVertexArray(0)

	return obj
}

// UpdateInstances ...
func (o *Object) UpdateInstances(newInstances []int32) {
	gl.BindBuffer(gl.ARRAY_BUFFER, o.InstanceBuffer)
	gl.BufferData(gl.ARRAY_BUFFER, len(newInstances)*4, gl.Ptr(newInstances), gl.STATIC_DRAW)
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	o.InstanceCount = int32(len(newInstances) / 4)
}

// Draw ...
func (o *Object) Draw() {
	gl.BindVertexArray(o.Ptr)
	gl.DrawArraysInstanced(gl.TRIANGLES, 0, o.VertexCount, o.InstanceCount)
}
