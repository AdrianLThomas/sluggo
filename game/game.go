package game

import (
	"fmt"
	"image/color"
	"math"
	"sluggo/game/arena"
	"sluggo/lib"
	"sluggo/types"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// internalTileSize smaller is more pixelated
const internalTileSize = 12

// splashDimAlpha is how opaque the veil over the arena is during splash
const splashDimAlpha = 0.6

type game struct {
	columns int
	rows    int
	arena   *arena.Arena
	score   int
	state   GameState
	buffer  *ebiten.Image
	veil    *ebiten.Image
}

func (g *game) Update() error {
	switch g.state {
	case StateSplash:
		if g.anyInputPressed() {
			g.state = StatePlaying
		}
	case StateGameOver:
		if g.anyInputPressed() {
			g.reset()
		}
	case StatePlaying:
		outcome, err := g.arena.Update()
		if err != nil {
			return err
		}

		g.score += outcome.ScoreGain
		if outcome.GameOver {
			g.state = StateGameOver
		}
	}

	return nil
}

// anyInputPressed reports whether the player asked to move on, by pressing a
// key, clicking or tapping. The splash and game over screens both wait on it;
// neither advances on its own.
func (g *game) anyInputPressed() bool {
	if len(inpututil.AppendJustPressedKeys(nil)) > 0 {
		return true
	}

	return inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) ||
		len(inpututil.AppendJustPressedTouchIDs(nil)) > 0
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

	if g.state == StateSplash {
		g.drawSplash(screen, screenSize)
		return
	}

	if g.state == StateGameOver {
		lib.NewOnscreenText("GAME OVER", lib.OnscreenTextConfig{
			Colour:     color.RGBA{A: 255, R: 255},
			Position:   lib.HorizontalCentre,
			Size:       lib.SizeLarge,
			ScreenSize: screenSize,
		}).Draw(screen)
		lib.NewOnscreenText("Press any key to play again", lib.OnscreenTextConfig{
			Colour:     color.RGBA{A: 255, G: 255},
			Position:   lib.HorizontalCentre | lib.VerticalCentre,
			Size:       lib.SizeSmall,
			ScreenSize: screenSize,
		}).Draw(screen)
	}

	lib.NewOnscreenText(fmt.Sprintf("Score: %d", g.score), lib.OnscreenTextConfig{
		Colour:     color.RGBA{A: 255, R: 255, G: 255},
		Position:   lib.HorizontalCentre | lib.VerticalBottom,
		Size:       lib.SizeExtraSmall,
		ScreenSize: screenSize,
	}).Draw(screen)
}

// drawSplash renders the splash screen. The arena is still drawn underneath
// it, dimmed by a veil, but the slug is never updated so it sits still.
func (g *game) drawSplash(screen *ebiten.Image, screenSize types.Vector2) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(float64(screenSize.X), float64(screenSize.Y))
	op.ColorScale.Scale(0, 0, 0, splashDimAlpha)
	screen.DrawImage(g.veil, op)

	lib.NewOnscreenText("Sluggo!", lib.OnscreenTextConfig{
		Colour:     color.RGBA{A: 255, R: 255, G: 255},
		Position:   lib.HorizontalCentre | lib.VerticalCentre,
		ScreenSize: screenSize,
		Size:       lib.SizeExtraLarge,
		Offset:     lib.Vector2[int]{Y: -60},
	}).Draw(screen)

	lib.NewOnscreenText("Arrow keys to steer", lib.OnscreenTextConfig{
		Colour:     color.RGBA{A: 255, R: 255, G: 255},
		Position:   lib.HorizontalCentre | lib.VerticalCentre,
		ScreenSize: screenSize,
		Size:       lib.SizeSmall,
		Offset:     lib.Vector2[int]{Y: -10},
	}).Draw(screen)

	lib.NewOnscreenText("Press any key or tap to start", lib.OnscreenTextConfig{
		Colour:     color.RGBA{A: 160, R: 255, G: 255, B: 255},
		Position:   lib.HorizontalCentre | lib.VerticalCentre,
		ScreenSize: screenSize,
		Size:       lib.SizeExtraSmall,
		Offset:     lib.Vector2[int]{Y: 24},
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
		state:   StateSplash,
		buffer:  ebiten.NewImage(columns*internalTileSize, rows*internalTileSize),
		veil:    newVeil(),
	}
	g.arena = arena.NewArena(columns, rows)
	return g
}

func newVeil() *ebiten.Image {
	veil := ebiten.NewImage(1, 1)
	veil.Fill(color.White)
	return veil
}
