package main

import (
	"bytes"
	"image"

	"github.com/samuel/go-pcx/pcx"
)

func Graphics(img image.Image) ([]byte, error) {
	var buffer = new(bytes.Buffer)
	err := pcx.Encode(buffer, img)
	return buffer.Bytes(), err
}
