package game

import (
	"sluggo/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	DirectionLeft  = Vector2{X: -1}
	DirectionUp    = Vector2{Y: -1}
	DirectionRight = Vector2{X: +1}
	DirectionDown  = Vector2{Y: +1}
)

type Slug struct {
	position         Vector2
	length           int
	sprite           *ebiten.Image
	jumpBy           int
	currentDirection Vector2
}

func (s *Slug) Update() error {
	switch {
	case ebiten.IsKeyPressed(ebiten.KeyLeft):
		s.currentDirection = DirectionLeft
	case ebiten.IsKeyPressed(ebiten.KeyUp):
		s.currentDirection = DirectionUp
	case ebiten.IsKeyPressed(ebiten.KeyRight):
		s.currentDirection = DirectionRight
	case ebiten.IsKeyPressed(ebiten.KeyDown):
		s.currentDirection = DirectionDown
	}

	return nil
}

func (s *Slug) Move() {
	moveBy := Vector2{X: s.currentDirection.X, Y: s.currentDirection.Y}
	moveBy.Multiply(s.jumpBy)
	s.position.Add(moveBy)
}

func (s *Slug) Draw(screen *ebiten.Image, tileSize int, offsetX int, offsetY int) {
	bounds := s.sprite.Bounds()

	scale := float64(tileSize) / float64(bounds.Dx())

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(scale, scale)

	x := offsetX + s.position.X*tileSize
	y := offsetY + s.position.Y*tileSize

	op.GeoM.Translate(float64(x), float64(y))

	screen.DrawImage(s.sprite, op)
}

func NewSlug(jumpBy int, startPosition Vector2) *Slug {
	sprite := assets.SlugSprite

	return &Slug{
		length:           1,
		position:         startPosition,
		sprite:           sprite,
		jumpBy:           jumpBy,
		currentDirection: DirectionLeft,
	}
}

func (s *Slug) Position() Vector2 {
	return s.position
}

func (s *Slug) SetPosition(position Vector2) {
	s.position = position
}
