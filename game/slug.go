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
	positions        []Vector2
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

	newPos := s.positions[0]
	newPos.Add(moveBy)
	s.SetPosition(newPos)
}

func (s *Slug) Draw(screen *ebiten.Image, tileSize int, offsetX int, offsetY int) {
	bounds := s.sprite.Bounds()

	// head
	opHead := &ebiten.DrawImageOptions{}
	s.center(opHead, bounds)
	s.rotate(opHead)
	s.scale(tileSize, bounds, opHead)
	s.translate(offsetX, tileSize, offsetY, opHead, s.positions[0])
	screen.DrawImage(s.sprite, opHead)

	// body
	for _, body := range s.positions[1:] {
		opBody := &ebiten.DrawImageOptions{}
		s.center(opBody, bounds)
		s.scale(tileSize, bounds, opBody)
		s.translate(offsetX, tileSize, offsetY, opBody, body)
		screen.DrawImage(s.sprite, opBody)
	}
}

func (s *Slug) scale(tileSize int, bounds image.Rectangle, op *ebiten.DrawImageOptions) {
	scale := float64(tileSize) / float64(bounds.Dx())
	op.GeoM.Scale(scale, scale)
}

func (s *Slug) translate(offsetX int, tileSize int, offsetY int, op *ebiten.DrawImageOptions, position Vector2) {
	x := offsetX + position.X*tileSize
	y := offsetY + position.Y*tileSize

	op.GeoM.Translate(
		float64(x+tileSize/2),
		float64(y+tileSize/2),
	)
}

func (s *Slug) rotate(op *ebiten.DrawImageOptions) {
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

// center rotates around the center of the sprite
func (s *Slug) center(op *ebiten.DrawImageOptions, bounds image.Rectangle) {
	op.GeoM.Translate(
		-float64(bounds.Dx())/2,
		-float64(bounds.Dy())/2,
	)
}

func NewSlug(jumpBy int, startPosition Vector2) *Slug {
	sprite := assets.SlugSprite

	return &Slug{
		positions: []Vector2{
			startPosition,
			{startPosition.X + 1, startPosition.Y},
			{startPosition.X + 2, startPosition.Y},
		},
		sprite:           sprite,
		jumpBy:           jumpBy,
		currentDirection: DirectionLeft,
		nextDirection:    DirectionLeft,
	}
}

// Position returns the position of the head
func (s *Slug) Position() Vector2 {
	return s.positions[0]
}

func (s *Slug) SetPosition(position Vector2) {
	// update body to each prior element
	for i := len(s.positions) - 1; i > 0; i-- {
		s.positions[i] = s.positions[i-1]
	}

	// set head
	s.positions[0] = position
}

// 1 2 3
