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

type game struct {
	columns int
	rows    int
	arena   *arena.Arena
	score   int
	state   GameState
	buffer  *ebiten.Image
}

func (g *game) Update() error {
	if g.state != StatePlaying {
		return nil
	}

	outcome, err := g.arena.Update()
	if err != nil {
		return err
	}

	g.score += outcome.ScoreGain
	if outcome.GameOver {
		g.state = StateGameOver
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

	if g.state == StateGameOver {
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
			g.reset()
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

func (g *game) reset() {
	g.arena = arena.NewArena(g.columns, g.rows)
	g.score = 0
	g.state = StatePlaying
}

func NewGame(columns, rows int) ebiten.Game {
	g := &game{
		columns: columns,
		rows:    rows,
		state:   StatePlaying,
		buffer:  ebiten.NewImage(columns*internalTileSize, rows*internalTileSize),
	}
	g.reset()
	return g
}
