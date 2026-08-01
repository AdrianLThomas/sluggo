package game

import (
	"fmt"
	"sluggo/lib"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	ScreenWidth   = 1024
	ScreenHeight  = 768
	MoveFrequency = time.Millisecond * 150
	JumpBy        = 1
	Columns       = 20
	Rows          = 15
	SlugLength    = 10
)

type Vector2 = lib.Vector2[int]

type game struct {
	arena *Arena
}

func (g *game) Update() error {
	return g.arena.Update()
}

func (g *game) Draw(screen *ebiten.Image) {
	width, height := screen.Bounds().Dx(), screen.Bounds().Dy()
	tileSize := min(width/g.arena.columns, height/g.arena.rows)
	offsetX := (width - (g.arena.columns * tileSize)) / 2
	offsetY := (height - (g.arena.rows * tileSize)) / 2

	g.arena.Draw(screen, tileSize, offsetX, offsetY)

	ebitenutil.DebugPrintAt(screen,
		fmt.Sprintf("X: %v,Y: %v", g.arena.slug.Position().X, g.arena.slug.Position().Y),
		0, ScreenHeight-20)
	ebitenutil.DebugPrint(screen, "Sluggo!")
}

func (g *game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func NewGame() ebiten.Game {
	return &game{
		arena: NewArena(Columns, Rows),
	}
}
