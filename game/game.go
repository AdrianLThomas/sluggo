package game

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	ScreenWidth   = 800
	ScreenHeight  = 600
	MoveFrequency = time.Millisecond * 200
	JumpBy        = 1
	Columns       = 9
	Rows          = 9
)

type game struct {
	arena *Arena
}

func (g game) Update() error {
	if err := g.arena.Update(); err != nil {
		return err
	}

	return nil
}

func (g game) Draw(screen *ebiten.Image) {
	g.arena.Draw(screen)

	ebitenutil.DebugPrint(screen, "Sluggo!")
}

func (g game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func NewGame() ebiten.Game {
	return &game{
		arena: NewArena(Columns, Rows),
	}
}
