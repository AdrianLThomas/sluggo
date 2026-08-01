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
	MoveFrequency = time.Millisecond * 200
	JumpBy        = 30.0
)

type game struct {
	slug      *Slug
	moveTimer *lib.Timer
}

func (g game) Update() error {
	if err := g.slug.Update(); err != nil {
		return err
	}

	g.moveTimer.Update()

	if g.moveTimer.IsReady() {
		g.moveTimer.Reset()

		g.slug.Move()
	}

	g.checkBounds()

	return nil
}

func (g game) checkBounds() {
	slugPosition := g.slug.Position()
	switch {
	case slugPosition.X < 0:
		g.slug.SetPosition(&lib.Vector2D{X: ScreenWidth, Y: slugPosition.Y})
	case slugPosition.X > ScreenWidth:
		g.slug.SetPosition(&lib.Vector2D{X: 0, Y: slugPosition.Y})
	case slugPosition.Y < 0:
		g.slug.SetPosition(&lib.Vector2D{X: slugPosition.X, Y: ScreenHeight})
	case slugPosition.Y > ScreenHeight:
		g.slug.SetPosition(&lib.Vector2D{X: slugPosition.X, Y: 0})
	}
}

func (g game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Sluggo!")
	ebitenutil.DebugPrintAt(screen,
		fmt.Sprintf("X: %f,Y: %f", g.slug.Position().X, g.slug.Position().Y),
		0, ScreenHeight-20)

	g.slug.Draw(screen)
}

func (g game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func NewGame() ebiten.Game {
	return &game{
		slug:      NewSlug(JumpBy),
		moveTimer: lib.NewTimer(MoveFrequency),
	}
}
