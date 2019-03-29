package main

import (
	"errors"
	"image"
	"image/draw"
	"image/png"
	"os"

	"github.com/go-gl/gl/v4.1-core/gl"
)

func textureFromData(img image.Image, id uint32) (uint32, error) {
	rgba := image.NewRGBA(img.Bounds())
	draw.Draw(rgba, rgba.Bounds(), img, image.Pt(0, 0), draw.Src)
	if rgba.Stride != rgba.Rect.Size().X*4 {
		return 0, errors.New("Incorrect stride")
	}

	var textureID uint32
	gl.GenTextures(1, &textureID)

	gl.ActiveTexture(id)
	gl.BindTexture(gl.TEXTURE_2D, textureID)

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)

	width := int32(rgba.Rect.Size().X)
	height := int32(rgba.Rect.Size().Y)
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, width, height, 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(rgba.Pix))
	gl.GenerateMipmap(textureID)

	return textureID, nil
}

// NewTextureFromFile ...
func NewTextureFromFile(file string, id uint32) (uint32, error) {
	imgFile, err := os.Open(file)
	if err != nil {
		return 0, err
	}
	defer imgFile.Close()

	img, err := png.Decode(imgFile)
	if err != nil {
		return 0, err
	}

	return textureFromData(img, id)
}
