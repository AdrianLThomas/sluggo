package game

import (
	"fmt"
	"image/color"
	"sluggo/assets"
	"sluggo/lib"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

const (
	ScreenWidth   = 1024
	ScreenHeight  = 768
	MoveFrequency = time.Millisecond * 150
	JumpBy        = 1
	Columns       = 10
	Rows          = 7
	SlugLength    = 1
)

type Vector2 = lib.Vector2[int]

// TODO introduce state machine
var IsGameOver = false

type game struct {
	arena *Arena
}

func (g *game) Update() error {
	if !IsGameOver {
		return g.arena.Update()
	}

	return nil
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

	if IsGameOver {
		str := "GAME OVER"
		w := font.MeasureString(assets.Font, str)
		text.Draw(screen, str, assets.Font, (ScreenWidth-w.Ceil())/2, 50, color.RGBA{A: 255, R: 255})
	}
}

func (g *game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func NewGame() ebiten.Game {
	return &game{
		arena: NewArena(Columns, Rows),
	}
}
