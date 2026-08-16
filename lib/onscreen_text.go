package lib

import (
	"image/color"
	"sluggo/assets"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Position uint8

const (
	None             Position = 0
	HorizontalCentre Position = 1 << iota
	VerticalCentre
)

type OnscreenText struct {
	message string
	config  OnscreenTextConfig
}
type OnscreenTextConfig struct {
	Colour     color.Color
	Position   Position
	ScreenSize Vector2[int]
}

func (t OnscreenText) Draw(screen *ebiten.Image) {
	const margin = 50.0
	w, h := text.Measure(t.message, assets.Font, 0)
	x, y := 0.0, margin
	if t.config.Position&HorizontalCentre != 0 {
		x = (float64(t.config.ScreenSize.X) - w) / 2
	}
	if t.config.Position&VerticalCentre != 0 {
		y = (float64(t.config.ScreenSize.Y) - h) / 2
	}

	op := &text.DrawOptions{}
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

