package game

import (
	"sluggo/assets"
	"sluggo/lib"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	DirectionLeft  = lib.Vector2D{X: -1}
	DirectionUp    = lib.Vector2D{Y: -1}
	DirectionRight = lib.Vector2D{X: +1}
	DirectionDown  = lib.Vector2D{Y: +1}
)

type Slug struct {
	position         *lib.Vector2D
	length           int
	sprite           *ebiten.Image
	jumpBy           int
	currentDirection *lib.Vector2D
}

func (s *Slug) Update() error {
	switch {
	case ebiten.IsKeyPressed(ebiten.KeyLeft):
		s.currentDirection = &DirectionLeft
	case ebiten.IsKeyPressed(ebiten.KeyUp):
		s.currentDirection = &DirectionUp
	case ebiten.IsKeyPressed(ebiten.KeyRight):
		s.currentDirection = &DirectionRight
	case ebiten.IsKeyPressed(ebiten.KeyDown):
		s.currentDirection = &DirectionDown
	}

	return nil
}

func (s *Slug) Move() {
	moveBy := lib.Vector2D{X: s.currentDirection.X, Y: s.currentDirection.Y}
	moveBy.Multiply(float64(s.jumpBy))
	s.position.Add(&moveBy)
}

func (s *Slug) Draw(screen *ebiten.Image) {
	bounds := s.sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(-halfW, -halfH)
	op.GeoM.Translate(
		s.position.X*float64(bounds.Dx()),
		s.position.Y*float64(bounds.Dy()),
	)

	screen.DrawImage(s.sprite, op)
}

func NewSlug(jumpBy int, startPosition *lib.Vector2D) *Slug {
	sprite := assets.SlugSprite

	return &Slug{
		length:           1,
		position:         startPosition,
		sprite:           sprite,
		jumpBy:           jumpBy,
		currentDirection: &DirectionLeft,
	}
}

func (s *Slug) Position() *lib.Vector2D {
	return s.position
}

func (s *Slug) SetPosition(position *lib.Vector2D) {
	s.position = position
}
