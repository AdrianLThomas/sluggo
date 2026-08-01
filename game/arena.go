package game

import (
	"sluggo/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

type Arena struct {
	columns int
	rows    int
	slug    *Slug
}

func (a *Arena) Update() error {
	if err := a.slug.Update(); err != nil {
		return err
	}

	a.slug.Wrap(a.columns, a.rows)

	return nil
}

func (a *Arena) Draw(screen *ebiten.Image, tileSize int, offsetX int, offsetY int) {
	for y := range a.rows {
		for x := range a.columns {
			op := &ebiten.DrawImageOptions{}
			s := float64(tileSize) / float64(assets.BackgroundSprite.Bounds().Dx())
			op.GeoM.Scale(s, s)
			op.GeoM.Translate(
				float64(offsetX+x*tileSize),
				float64(offsetY+y*tileSize),
			)
			screen.DrawImage(assets.BackgroundSprite, op)
		}
	}

	a.slug.Draw(screen, tileSize, offsetX, offsetY)
}

func NewArena(columns int, rows int) *Arena {
	return &Arena{
		columns: columns,
		rows:    rows,
		slug: NewSlug(JumpBy, Vector2{
			X: columns - 1,
			Y: rows / 2,
		},
			SlugLength),
	}
}
