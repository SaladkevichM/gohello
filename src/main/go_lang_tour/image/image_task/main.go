package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/color"
	"image/png"
)

type Image struct {
	Width, Height int
	color         uint8
}

func (r *Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, r.Width, r.Height)
}

func (r *Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (r *Image) At(x, y int) color.Color {
	return color.RGBA{r.color + uint8(x), r.color + uint8(y), 255, 255}
}

func ShowImage(m image.Image) {
	var buf bytes.Buffer
	err := png.Encode(&buf, m)
	if err != nil {
		panic(err)
	}
	enc := base64.StdEncoding.EncodeToString(buf.Bytes())
	fmt.Println("IMAGE:" + enc)
}

func main() {
	m := Image{5, 5, 5}
	ShowImage(&m)
}
