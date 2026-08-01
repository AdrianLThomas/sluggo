package game

import (
	"sluggo/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

type Food struct {
	position Vector2
	sprite   *ebiten.Image
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

func (f *Food) Reset(position Vector2) {
	f.position = position
}

func NewFood(position Vector2) *Food {
	return &Food{
		position: position,
		sprite:   assets.FoodSprite,
	}
}
