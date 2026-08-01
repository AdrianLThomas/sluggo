package game

import (
	"sluggo/lib"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type game struct {
	player *Slug
}

func (g game) Update() error {
	if err := g.player.Update(); err != nil {
		return err
	}
	return nil
}

func (g game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Sluggo!")

	g.player.Draw(screen)
}

func (g game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func NewGame() ebiten.Game {
	return &game{
		player: NewSlug(lib.Vector2D{X: 1.0, Y: 1.0}), // TODO - center of screen
	}
}
