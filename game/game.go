package game

import (
	"fmt"
	"image/color"
	"math"
	"sluggo/game/arena"
	"sluggo/lib"
	"sluggo/types"

	"github.com/hajimehoshi/ebiten/v2"
)

// internalTileSize smaller is more pixelated
const internalTileSize = 12

var gameState = StatePlaying

type game struct {
	columns int
	rows    int
	arena   *arena.Arena
	score   int
	buffer  *ebiten.Image
}

func (g *game) Update() error {
	if gameState == StatePlaying {
		return g.arena.Update()
	}

	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	width, height := screen.Bounds().Dx(), screen.Bounds().Dy()

	internalSize := types.Vector2{X: g.buffer.Bounds().Dx(), Y: g.buffer.Bounds().Dy()}

	scale := math.Min(float64(width)/float64(internalSize.X), float64(height)/float64(internalSize.Y))
	op := &ebiten.DrawImageOptions{}
	op.Filter = ebiten.FilterPixelated
	op.GeoM.Scale(scale, scale)
	op.GeoM.Translate(
		(float64(width)-float64(internalSize.X)*scale)/2,
		(float64(height)-float64(internalSize.Y)*scale)/2,
	)

	g.arena.Draw(g.buffer, internalTileSize, 0, 0)
	screen.DrawImage(g.buffer, op)

	screenSize := types.Vector2{X: width, Y: height}

	if gameState == StateGameOver {
		lib.NewOnscreenText("GAME OVER", lib.OnscreenTextConfig{
			Colour:     color.RGBA{A: 255, R: 255},
			Position:   lib.HorizontalCentre,
			Size:       lib.SizeLarge,
			ScreenSize: screenSize,
		}).Draw(screen)
		lib.NewOnscreenText("Press space to play again", lib.OnscreenTextConfig{
			Colour:     color.RGBA{A: 255, G: 255},
			Position:   lib.HorizontalCentre | lib.VerticalCentre,
			Size:       lib.SizeSmall,
			ScreenSize: screenSize,
		}).Draw(screen)

		if ebiten.IsKeyPressed(ebiten.KeySpace) {
			g.arena = arena.NewArena(g.columns, g.rows, onGameOver, g.incrementScore)
			g.score = 0
			gameState = StatePlaying
		}
	}

	lib.NewOnscreenText(fmt.Sprintf("Score: %d", g.score), lib.OnscreenTextConfig{
		Colour:     color.RGBA{A: 255, R: 255, G: 255},
		Position:   lib.HorizontalCentre | lib.VerticalBottom,
		Size:       lib.SizeExtraSmall,
		ScreenSize: screenSize,
	}).Draw(screen)
}

func (g *game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func onGameOver() {
	gameState = StateGameOver
}

func (g *game) incrementScore(multiplier int) {
	g.score = g.score + (1 * multiplier)
}

func NewGame(columns, rows int) ebiten.Game {
	g := &game{
		columns: columns,
		rows:    rows,
		buffer:  ebiten.NewImage(columns*internalTileSize, rows*internalTileSize),
	}
	g.arena = arena.NewArena(columns, rows, onGameOver, g.incrementScore)
	return g
}
