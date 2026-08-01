package game

import (
	"image"
	"math"
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
	nextDirection    Vector2
}

func (s *Slug) Update() error {
	switch {
	case ebiten.IsKeyPressed(ebiten.KeyLeft) && s.currentDirection != DirectionRight:
		s.nextDirection = DirectionLeft
	case ebiten.IsKeyPressed(ebiten.KeyUp) && s.currentDirection != DirectionDown:
		s.nextDirection = DirectionUp
	case ebiten.IsKeyPressed(ebiten.KeyRight) && s.currentDirection != DirectionLeft:
		s.nextDirection = DirectionRight
	case ebiten.IsKeyPressed(ebiten.KeyDown) && s.currentDirection != DirectionUp:
		s.nextDirection = DirectionDown
	}

	return nil
}

func (s *Slug) Move() {
	s.currentDirection = s.nextDirection
	moveBy := Vector2{X: s.currentDirection.X, Y: s.currentDirection.Y}
	moveBy.Multiply(s.jumpBy)
	s.position.Add(moveBy)
}

func (s *Slug) Draw(screen *ebiten.Image, tileSize int, offsetX int, offsetY int) {
	bounds := s.sprite.Bounds()

	op := &ebiten.DrawImageOptions{}

	s.rotate(op, bounds)
	s.scale(tileSize, bounds, op)
	s.translate(offsetX, tileSize, offsetY, op)

	screen.DrawImage(s.sprite, op)
}

func (s *Slug) scale(tileSize int, bounds image.Rectangle, op *ebiten.DrawImageOptions) {
	scale := float64(tileSize) / float64(bounds.Dx())
	op.GeoM.Scale(scale, scale)
}

func (s *Slug) translate(offsetX int, tileSize int, offsetY int, op *ebiten.DrawImageOptions) {
	x := offsetX + s.position.X*tileSize
	y := offsetY + s.position.Y*tileSize

	op.GeoM.Translate(
		float64(x+tileSize/2),
		float64(y+tileSize/2),
	)
}

func (s *Slug) rotate(op *ebiten.DrawImageOptions, bounds image.Rectangle) {
	// rotate around the center of the sprite
	op.GeoM.Translate(
		-float64(bounds.Dx())/2,
		-float64(bounds.Dy())/2,
	)

	switch s.currentDirection {
	case DirectionLeft:
		op.GeoM.Rotate(0)
	case DirectionUp:
		op.GeoM.Rotate(90 * math.Pi / 180)
	case DirectionRight:
		op.GeoM.Rotate(180 * math.Pi / 180)
	case DirectionDown:
		op.GeoM.Rotate(270 * math.Pi / 180)
	}
}

func NewSlug(jumpBy int, startPosition Vector2) *Slug {
	sprite := assets.SlugSprite

	return &Slug{
		length:           1,
		position:         startPosition,
		sprite:           sprite,
		jumpBy:           jumpBy,
		currentDirection: DirectionLeft,
		nextDirection:    DirectionLeft,
	}
}

func (s *Slug) Position() Vector2 {
	return s.position
}

func (s *Slug) SetPosition(position Vector2) {
	s.position = position
}
