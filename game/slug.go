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

var startSpeed float64 = 1

func (s *Slug) Update() error {
	speed := startSpeed * 1
	s.position.Add(speed, 0)

	return nil
}

func (s *Slug) Draw(screen *ebiten.Image) {
	//bounds := s.sprite.Bounds()
	//halfW := float64(bounds.Dx()) / 2
	//halfH := float64(bounds.Dy()) / 2

	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(s.position.X, s.position.Y)

	screen.DrawImage(s.sprite, op)
}

func NewSlug(startPosition lib.Vector2D) *Slug {
	return &Slug{
		length:   1,
		position: startPosition,
		sprite:   assets.SlugSprite,
	}
}
