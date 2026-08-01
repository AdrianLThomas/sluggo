package game

import (
	"sluggo/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

type Arena struct {
	background *ebiten.Image
	columns    int
	rows       int
}

func (a Arena) Update() error {
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
}

func NewArena(columns int, rows int) *Arena {
	return &Arena{
		columns:    columns,
		rows:       rows,
		background: assets.BackgroundSprite,
	}
}
