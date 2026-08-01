package game

import (
	"fmt"
	"sluggo/assets"
	"sluggo/lib"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Arena struct {
	background *ebiten.Image
	columns    int
	rows       int
	slug       *Slug
	moveTimer  *lib.Timer
}

func (a Arena) Update() error {
	if err := a.slug.Update(); err != nil {
		return err
	}

	a.moveTimer.Update()

	if a.moveTimer.IsReady() {
		a.moveTimer.Reset()

		a.slug.Move()
	}

	a.checkBounds()

	return nil
}

func (a Arena) Draw(screen *ebiten.Image) {
	width, height := screen.Bounds().Dx(), screen.Bounds().Dy()
	tileSize := min(width/a.columns, height/a.rows)
	offsetX := (width - (a.columns * tileSize)) / 2
	offsetY := (height - (a.rows * tileSize)) / 2
	for y := range a.rows {
		for x := range a.columns {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Scale(
				float64(tileSize)/float64(a.background.Bounds().Dx()),
				float64(tileSize)/float64(a.background.Bounds().Dx()),
			)
			op.GeoM.Translate(
				float64(offsetX+x*tileSize),
				float64(offsetY+y*tileSize),
			)
			screen.DrawImage(a.background, op)
		}
	}

	a.slug.Draw(screen)

	ebitenutil.DebugPrintAt(screen,
		fmt.Sprintf("X: %f,Y: %f", a.slug.Position().X, a.slug.Position().Y),
		0, ScreenHeight-20)
}

func NewArena(columns int, rows int) *Arena {
	return &Arena{
		columns:    columns,
		rows:       rows,
		background: assets.BackgroundSprite,
		slug: NewSlug(JumpBy, &lib.Vector2D{
			X: Columns,
			Y: Rows / 2,
		}),
		moveTimer: lib.NewTimer(MoveFrequency),
	}
}

func (a Arena) checkBounds() {
	//slugPosition := a.slug.Position()
	//switch {
	//case slugPosition.X < 0:
	//	a.slug.SetPosition(&lib.Vector2D{X: ScreenWidth, Y: slugPosition.Y})
	//case slugPosition.X > ScreenWidth:
	//	a.slug.SetPosition(&lib.Vector2D{X: 0, Y: slugPosition.Y})
	//case slugPosition.Y < 0:
	//	a.slug.SetPosition(&lib.Vector2D{X: slugPosition.X, Y: ScreenHeight})
	//case slugPosition.Y > ScreenHeight:
	//	a.slug.SetPosition(&lib.Vector2D{X: slugPosition.X, Y: 0})
	//}
}
