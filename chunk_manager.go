package main

import (
	"strconv"

	"github.com/aquilax/go-perlin"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

const chunkSize = 8
const alpha = 5.
const beta = 2.
const n = 3
const seed = 100

type Chunk struct {
	object   Object
	position [3]int32
}

func newChunk(pos [3]int32, p *perlin.Perlin) Chunk {
	var chunk Chunk
	chunk.object = newObject(cubeVertices, cubeUVs, cubeNormals)

	var newInstancies = make([]int32, chunkSize*chunkSize*5*4)
	for i := 0; i < chunkSize; i++ {
		for j := 0; j < chunkSize; j++ {
			height := int32(10 * p.Noise2D(
				float64(int32(i)+pos[0]*chunkSize)/10, float64(int32(j)+pos[2]*chunkSize)/10))
			newInstancies[i*chunkSize*4*5+j*4*5] = int32(i)
			newInstancies[i*chunkSize*4*5+j*4*5+1] = int32(10 * p.Noise2D(float64(int32(i)+pos[0]*chunkSize)/10, float64(int32(j)+pos[2]*chunkSize)/10))
			newInstancies[i*chunkSize*4*5+j*4*5+2] = int32(j)
			newInstancies[i*chunkSize*4*5+j*4*5+3] = 0 //grass
			for s := 1; s < 5; s++ {
				newInstancies[i*chunkSize*4*5+j*4*5+4*s] = int32(i)
				newInstancies[i*chunkSize*4*5+j*4*5+4*s+1] = height - int32(s)
				newInstancies[i*chunkSize*4*5+j*4*5+4*s+2] = int32(j)
				newInstancies[i*chunkSize*4*5+j*4*5+4*s+3] = 1 //stone
			}
		}
	}
	chunk.object.UpdateInstances(newInstancies)
	chunk.position = pos

	return chunk
}

type ChunkManager struct {
	chunks map[string]Chunk
	p      *perlin.Perlin
}

func newChunkManager() ChunkManager {
	// var chunks = make([]Chunk, 25)

	var chunkManager ChunkManager
	chunkManager.chunks = make(map[string]Chunk)
	chunkManager.p = perlin.NewPerlin(alpha, beta, n, seed)

	return chunkManager
}

func (c *ChunkManager) addChunk(key string, i, j int32) {
	c.chunks[key] = newChunk([3]int32{i, 0, j}, c.p)
}

func (c *ChunkManager) draw(pos mgl32.Vec3) {
	for i := -4; i <= 4; i++ {
		for j := -4; j <= 4; j++ {
			I := int(pos.X()/chunkSize) + i
			J := int(pos.Z()/chunkSize) + j
			key := strconv.Itoa(I) + ";" + strconv.Itoa(J)

			if val, ok := c.chunks[key]; ok {
				transChunk := gl.GetUniformLocation(program, gl.Str("transChunk\x00"))
				gl.Uniform3i(transChunk, val.position[0], val.position[1], val.position[2])

				sizeChunk := gl.GetUniformLocation(program, gl.Str("sizeChunk\x00"))
				gl.Uniform1i(sizeChunk, chunkSize)

				gl.BindVertexArray(val.object.Ptr)
				gl.DrawArraysInstanced(gl.TRIANGLES, 0, val.object.VertexCount, val.object.InstanceCount)
			} else {
				c.addChunk(key, int32(I), int32(J))
			}
		}

	}
}
