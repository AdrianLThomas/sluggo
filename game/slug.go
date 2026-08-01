package game

import "sluggo/lib"

type slug struct {
	position lib.Vector2D
	length   int
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
		position: startPosition,
	}
}
