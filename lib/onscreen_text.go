package lib

import (
	"image/color"
	"sluggo/assets"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Position uint8

const (
	None           Position = 0
	HorizontalLeft Position = 1 << iota
	HorizontalCentre
	HorizontalRight
	VerticalTop
	VerticalCentre
	VerticalBottom
)

type Size float64

const (
	SizeExtraSmall Size = 0.5
	SizeSmall      Size = 1
	SizeMedium     Size = 1.2
	SizeLarge      Size = 1.4
	SizeExtraLarge Size = 1.6
)

type OnscreenText struct {
	message string
	config  OnscreenTextConfig
}
type OnscreenTextConfig struct {
	Colour     color.Color
	Position   Position
	ScreenSize Vector2[int]
	Size       Size
}

func (t OnscreenText) Draw(screen *ebiten.Image) {
	const margin = 50.0
	scale := float64(t.config.Size)
	w, h := text.Measure(t.message, assets.Font, 0)
	w, h = w*scale, h*scale
	x, y := 0.0, margin
	switch {
	case t.config.Position&HorizontalLeft != 0:
		x = margin
	case t.config.Position&HorizontalCentre != 0:
		x = (float64(t.config.ScreenSize.X) - w) / 2
	case t.config.Position&HorizontalRight != 0:
		x = float64(t.config.ScreenSize.X) - w - margin
	}

	switch {
	case t.config.Position&VerticalTop != 0:
		y = margin
	case t.config.Position&VerticalCentre != 0:
		y = (float64(t.config.ScreenSize.Y) - h) / 2
	case t.config.Position&VerticalBottom != 0:
		y = float64(t.config.ScreenSize.Y) - h - margin
	}

	op := &text.DrawOptions{}
	op.GeoM.Scale(scale, scale)
	op.GeoM.Translate(x, y)
	if t.config.Colour != nil {
		op.ColorScale.ScaleWithColor(t.config.Colour)
	}
	text.Draw(screen, t.message, assets.Font, op)
}

func NewOnscreenText(message string, config OnscreenTextConfig) *OnscreenText {
	return &OnscreenText{
		message,
		config,
	}
}
