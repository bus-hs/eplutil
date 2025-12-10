package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
	"strings"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/math/fixed"

	_ "embed"
)

//go:embed res/font.ttf
var fontFile []byte

var textFont *opentype.Font

const PADDING = 20
const MARGIN = 20

func init() {
	var err error
	textFont, err = opentype.Parse(fontFile)
	if err != nil {
		panic(err)
	}
}

func LoadFont(size float64) font.Face {
	face, err := opentype.NewFace(textFont, &opentype.FaceOptions{Size: size, DPI: 72})
	if err != nil {
		panic(err)
	}
	return face
}

func MakeTextImg(text string) image.Image {
	img := image.NewGray(image.Rect(0, 0, WIDTH, HEIGHT))
	draw.Draw(img, image.Rect(0, 0, WIDTH, HEIGHT), image.NewUniform(color.White), image.Pt(0, 0), draw.Src)

	testFontSize := float64(96)

	face := LoadFont(testFontSize)

	lines := strings.Split(text, "\n")
	lineHeight := face.Metrics().CapHeight.Ceil()
	height := len(lines) * lineHeight
	widths := make([]int, len(lines))
	width := 0
	for idx, line := range lines {
		lineWidth := font.MeasureString(face, line).Ceil()
		widths[idx] = lineWidth
		if idx == 0 || lineWidth > width {
			width = lineWidth
		}
	}

	widthSizeFactor := (WIDTH - 2 * MARGIN) / float64(width)
	heightSizeFactor := (HEIGHT - float64((len(lines) - 1) * PADDING + 2 * MARGIN)) / float64(height)
	adjustWidth := true
	sizeFactor := widthSizeFactor
	if heightSizeFactor < widthSizeFactor {
		adjustWidth = false
		sizeFactor = heightSizeFactor
	}
	face = LoadFont(testFontSize * sizeFactor)

	realLineHeight := face.Metrics().CapHeight.Ceil()
	realHeight := realLineHeight * len(lines) + PADDING * (len(lines) - 1)

	for idx, line := range lines {
		startX := MARGIN
		startY := MARGIN
		if adjustWidth {
			startY = int(HEIGHT - realHeight) / 2
		} else {
			startX = int(WIDTH - float64(widths[idx]) * sizeFactor) / 2
		}
		d := &font.Drawer{
			Dst: img,
			Src: image.NewUniform(color.Black),
			Face: face,
			Dot: fixed.P(startX, startY + realLineHeight * (idx + 1) + PADDING * idx),
		}
		d.DrawString(line)
	}

	f, err := os.Create("test.png")
	if err != nil {
		panic(err)
	}
	err = png.Encode(f, img)
	if err != nil {
		panic(err)
	}

	return img
}
