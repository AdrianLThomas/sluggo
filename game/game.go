package game

import (
	"image/color"
	"sluggo/game/arena"
	"sluggo/lib"
	"sluggo/types"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var gameState = StatePlaying

type game struct {
	columns int
	rows    int
	arena   *arena.Arena
}

func (g *game) Update() error {
	if gameState == StatePlaying {
		return g.arena.Update()
	}

	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	width, height := screen.Bounds().Dx(), screen.Bounds().Dy()
	tileSize := min(width/g.columns, height/g.rows)
	offsetX := (width - (g.columns * tileSize)) / 2
	offsetY := (height - (g.rows * tileSize)) / 2

	g.arena.Draw(screen, tileSize, offsetX, offsetY)

	ebitenutil.DebugPrint(screen, "Sluggo!")

	if gameState == StateGameOver {
		bounds := screen.Bounds()
		screenSize := types.Vector2{X: bounds.Dx(), Y: bounds.Dy()}
		lib.NewOnscreenText("GAME OVER", lib.OnscreenTextConfig{
			Colour:     color.RGBA{A: 255, R: 255},
			Position:   lib.HorizontalCentre,
			ScreenSize: screenSize,
		}).Draw(screen)
		lib.NewOnscreenText("Press space or tap to play again", lib.OnscreenTextConfig{
			Colour:     color.RGBA{A: 255, G: 255},
			Position:   lib.HorizontalCentre | lib.VerticalCentre,
			ScreenSize: screenSize,
		}).Draw(screen)

		if ebiten.IsKeyPressed(ebiten.KeySpace) || len(inpututil.AppendJustPressedTouchIDs(nil)) > 0 {
			g.arena = arena.NewArena(g.columns, g.rows, onGameOver)
			gameState = StatePlaying
		}
	}
}

func (g *game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func onGameOver() {
	gameState = StateGameOver
}

func NewGame(columns, rows int) ebiten.Game {
	return &game{
		columns: columns,
		rows:    rows,
		arena:   arena.NewArena(columns, rows, onGameOver),
	}
}
