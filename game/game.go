package game

import (
	"fmt"
	"image/color"
	"sluggo/lib"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	MoveFrequency      = time.Millisecond * 150
	JumpBy             = 1
	StartingSlugLength = 1
)

var ScreenSize = Vector2{X: 1024, Y: 768}

type Vector2 = lib.Vector2[int]

// TODO introduce state machine
var IsGameOver = false

type game struct {
	columns int
	rows    int
	arena   *Arena
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
		0, ScreenSize.Y-20)
	ebitenutil.DebugPrint(screen, "Sluggo!")

	if IsGameOver {
		lib.NewOnscreenText("GAME OVER", lib.OnscreenTextConfig{
			Colour:     color.RGBA{A: 255, R: 255},
			Position:   lib.HorizontalCentre,
			ScreenSize: ScreenSize,
		}).Draw(screen)
		lib.NewOnscreenText("Press space to play again", lib.OnscreenTextConfig{
			Colour:     color.RGBA{A: 255, G: 255},
			Position:   lib.HorizontalCentre | lib.VerticalCentre,
			ScreenSize: ScreenSize,
		}).Draw(screen)

		if ebiten.IsKeyPressed(ebiten.KeySpace) {
			g.arena = NewArena(g.arena.columns, g.arena.rows)
			IsGameOver = false
		}
	}
}

func (g *game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func NewGame(columns, rows int) ebiten.Game {
	return &game{
		columns: columns,
		rows:    rows,
		arena:   NewArena(columns, rows),
	}
}
