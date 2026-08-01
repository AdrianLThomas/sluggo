package game

import (
	"sluggo/assets"
	"sluggo/lib"

	"github.com/hajimehoshi/ebiten/v2"
)

type Slug struct {
	position lib.Vector2D
	length   int
	sprite   *ebiten.Image
}

func (s Slug) Update() error {
	//TODO implement me
	return nil
}

func (s Slug) Draw(screen *ebiten.Image) {
	//bounds := s.sprite.Bounds()
	//halfW := float64(bounds.Dx()) / 2
	//halfH := float64(bounds.Dy()) / 2

	op := &ebiten.DrawImageOptions{}

	screen.DrawImage(s.sprite, op)
}

func NewSlug(startPosition lib.Vector2D) *Slug {
	return &Slug{
		length:   1,
		position: startPosition,
		sprite:   assets.SlugSprite,
	}
}
