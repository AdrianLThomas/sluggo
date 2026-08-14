package objects

import (
	"sluggo/assets"
	"sluggo/types"

	"github.com/hajimehoshi/ebiten/v2"
)

type Food struct {
	position types.Vector2
	sprite   *ebiten.Image
}

func (f *Food) Position() types.Vector2 {
	return f.position
}

func (f *Food) Update() error {
	return nil
}

func (f *Food) Draw(screen *ebiten.Image, tileSize int, offsetX int, offsetY int) {
	bounds := f.sprite.Bounds()

	op := &ebiten.DrawImageOptions{}
	scale := float64(tileSize) / float64(bounds.Dx())
	op.GeoM.Scale(scale, scale)
	op.GeoM.Translate(float64((tileSize*f.position.X)+offsetX), float64(tileSize*f.position.Y+offsetY))

	screen.DrawImage(f.sprite, op)
}

func (f *Food) Reset(position types.Vector2) {
	f.position = position
}

func NewFood(position types.Vector2) *Food {
	return &Food{
		position: position,
		sprite:   assets.FoodSprite,
	}
}
