package game

import (
	"sluggo/assets"
	"sluggo/lib"

	"github.com/hajimehoshi/ebiten/v2"
)

type slug struct {
	position lib.Vector2D
	length   int
	sprite   *ebiten.Image
}

func (s slug) Position() lib.Vector2D {
	return s.position
}

func (s slug) Length() int {
	return s.length
}

type Slug interface {
	Position() lib.Vector2D
	Length() int
}

func NewSlug(startPosition lib.Vector2D) Slug {
	return &slug{
		length:   1,
		position: startPosition,
		sprite:   assets.SlugSprite,
	}
}
