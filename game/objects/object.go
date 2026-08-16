package objects

import (
	"sluggo/types"

	"github.com/hajimehoshi/ebiten/v2"
)

type GameObject struct {
	position types.Vector2
	sprite   *ebiten.Image
}

func (o *GameObject) Position() types.Vector2 {
	return o.position
}

func (o *GameObject) SetPosition(position types.Vector2) {
	o.position = position
}

func (o *GameObject) Update() error {
	return nil
}

func (o *GameObject) Draw(screen *ebiten.Image, tileSize int, offsetX int, offsetY int) {
	bounds := o.sprite.Bounds()

	reduceBy := 1.7
	op := &ebiten.DrawImageOptions{}
	scale := float64(tileSize) / float64(bounds.Dx()) / reduceBy
	scaled := float64(bounds.Dx()) * scale
	centered := (float64(tileSize) - scaled) / 2
	op.GeoM.Scale(scale, scale)
	op.GeoM.Translate(float64(tileSize*o.position.X)+float64(offsetX)+centered, float64(tileSize*o.position.Y+offsetY)+centered)

	screen.DrawImage(o.sprite, op)
}
