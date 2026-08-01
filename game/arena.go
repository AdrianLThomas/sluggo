package game

import (
	"fmt"
	"sluggo/assets"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Arena struct {
	background *ebiten.Image
	columns    int
	rows       int
	slug       *Slug
}

func (a Arena) Update() error {
	if err := a.slug.Update(); err != nil {
		return err
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

	a.slug.Draw(screen, tileSize, offsetX, offsetY)

	ebitenutil.DebugPrintAt(screen,
		fmt.Sprintf("X: %v,Y: %v", a.slug.Position().X, a.slug.Position().Y),
		0, ScreenHeight-20)
}

func NewArena(columns int, rows int) *Arena {
	return &Arena{
		columns:    columns,
		rows:       rows,
		background: assets.BackgroundSprite,
		slug: NewSlug(JumpBy, Vector2{
			X: Columns - 1,
			Y: Rows / 2,
		},
		SlugLength),
	}
}

func (a Arena) checkBounds() {
	slugPosition := a.slug.Position()

	// overflow mode
	switch {
	case slugPosition.X < 0:
		a.slug.Teleport(Vector2{X: a.columns - 1, Y: slugPosition.Y})
	case slugPosition.X > a.columns-1:
		a.slug.Teleport(Vector2{X: 0, Y: slugPosition.Y})
	case slugPosition.Y < 0:
		a.slug.Teleport(Vector2{X: slugPosition.X, Y: a.rows - 1})
	case slugPosition.Y > a.rows-1:
		a.slug.Teleport(Vector2{X: slugPosition.X, Y: 0})
	}
}
