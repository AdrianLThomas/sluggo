package game

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	ScreenWidth  = 1024
	ScreenHeight = 768
)

type game struct {
	slug *Slug
}

func (g game) Update() error {
	if err := g.slug.Update(); err != nil {
		return err
	}
	return nil
}

func (g game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Sluggo!")
	ebitenutil.DebugPrintAt(screen,
		fmt.Sprintf("X: %f,Y: %f", g.slug.position.X, g.slug.position.Y),
		0, ScreenHeight-20)

	g.slug.Draw(screen)
}

func (g game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func NewGame() ebiten.Game {
	return &game{
		slug: NewSlug(),
	}
}
