package main

import (
	"image"

	"github.com/fogleman/gg"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"

	_ "embed"
)

const LINE_SPACING = 1

//go:embed res/font.ttf
var fontFile []byte

var textFont *opentype.Font

func init() {
	var err error
	textFont, err = opentype.Parse(fontFile)
	if err != nil {
		panic(err)
	}
}

func LoadFont(size float64) font.Face {
	face, err := opentype.NewFace(textFont, &opentype.FaceOptions{Size: size, DPI: 72, Hinting: font.HintingFull})
	if err != nil {
		panic(err)
	}
	return face
}

func MakeTextImg(text string) image.Image {
	dc := gg.NewContext(WIDTH, HEIGHT)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	testFontSize := float64(96)

	dc.SetFontFace(LoadFont(testFontSize))

	textWidth, textHeight := dc.MeasureMultilineString(text, LINE_SPACING)
	sizeFactor := min(WIDTH / float64(textWidth), HEIGHT / float64(textHeight))
	dc.SetFontFace(LoadFont(testFontSize * sizeFactor))
	dc.DrawStringWrapped(text, WIDTH / 2, HEIGHT / 2, 0.5, 0.5, WIDTH, LINE_SPACING, gg.AlignCenter)
	dc.Clip()
	dc.SavePNG("test.png")
	return dc.Image()
}
