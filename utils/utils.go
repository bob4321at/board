package utils

import "image/color"

var Mouse_X float64
var Mouse_Y float64
var Scroll_X float64
var Scroll_Y float64

type IColor struct {
	R, G, B, A int
}

func (col *IColor) TurnToColorRGBA() color.RGBA {
	return color.RGBA{uint8(col.R), uint8(col.G), uint8(col.B), uint8(col.A)}
}

type Vec2 struct {
	X, Y float64
}

type GeoM struct {
	X, Y          float64
	Width, Height float64
}

func (geom *GeoM) Translate(x, y float64) {
	geom.X = x
	geom.Y = y
}

func (geom *GeoM) Scale(w, h float64) {
	geom.Width = w
	geom.Height = h
}
