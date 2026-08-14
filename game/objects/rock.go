package objects

import (
	"sluggo/assets"
	"sluggo/types"

	"github.com/hajimehoshi/ebiten/v2"
)

type Rock struct {
	position types.Vector2
	sprite   *ebiten.Image
}

func (r *Rock) Position() types.Vector2 {
	return r.position
}

func (f *Rock) Update() error {
	return nil
}

func (f *Rock) Draw(screen *ebiten.Image, tileSize int, offsetX int, offsetY int) {
	bounds := f.sprite.Bounds()

	op := &ebiten.DrawImageOptions{}
	scale := float64(tileSize) / float64(bounds.Dx())
	op.GeoM.Scale(scale, scale)
	op.GeoM.Translate(float64((tileSize*f.position.X)+offsetX), float64(tileSize*f.position.Y+offsetY))

	screen.DrawImage(f.sprite, op)
}

func NewRock(position types.Vector2) *Rock {
	return &Rock{
		position: position,
		sprite:   assets.RockSprite,
	}
}
