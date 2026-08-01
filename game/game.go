package game

import (
	"fmt"
	"sluggo/lib"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
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
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("X: %f,Y: %f", g.slug.position.X, g.slug.position.Y), 150, 50)

	g.slug.Draw(screen)
}

func (g game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func NewGame() ebiten.Game {
	return &game{
		slug: NewSlug(lib.Vector2D{X: 1.0, Y: 1.0}), // TODO - center of screen
	}
}
