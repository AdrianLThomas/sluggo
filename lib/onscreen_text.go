package lib

import (
	"image/color"
	"sluggo/assets"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
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
	const margin = 50
	strWidth := font.MeasureString(assets.Font, t.message)
	x, y := 0, margin
	if t.config.Position&HorizontalCentre != 0 {
		x = (t.config.ScreenSize.X - strWidth.Ceil()) / 2
	}
	if t.config.Position&VerticalCentre != 0 {
		y = t.config.ScreenSize.Y / 2
	}
	text.Draw(screen, t.message, assets.Font, x, y, t.config.Colour)
}

func NewOnscreenText(message string, config OnscreenTextConfig) *OnscreenText {
	return &OnscreenText{
		message,
		config,
	}
}
