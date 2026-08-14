package characters

import (
	"image"
	"math"
	"slices"
	"sluggo/assets"
	"sluggo/lib"
	"sluggo/types"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	DirectionLeft  = types.Vector2{X: -1}
	DirectionUp    = types.Vector2{Y: -1}
	DirectionRight = types.Vector2{X: +1}
	DirectionDown  = types.Vector2{Y: +1}
)

type Slug struct {
	positions        []types.Vector2
	headSprite       *ebiten.Image
	bodySprite       *ebiten.Image
	tailSprite       *ebiten.Image
	jumpBy           int
	currentDirection types.Vector2
	nextDirection    types.Vector2
	moveTimer        *lib.Timer
	gridColumns      int
	gridRows         int
}

func (s *Slug) WillEatSelf() bool {
	return slices.Contains(s.positions, s.NextPosition())
}

func (s *Slug) Update() error {
	s.checkKeyPresses()

	s.moveTimer.Update()

	if s.moveTimer.IsReady() {
		s.moveTimer.Reset()
		s.move()
	}

	// wrap if necessary
	s.positions[0] = s.wrap(s.positions[0])

	return nil
}

func (s *Slug) checkKeyPresses() {
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
}

func (s *Slug) move() {
	s.currentDirection = s.nextDirection
	delta := s.currentDirection.Multiply(s.jumpBy)
	s.setPosition(s.positions[0].Add(delta))
}

func (s *Slug) setPosition(position types.Vector2) {
	for i := len(s.positions) - 1; i > 0; i-- {
		s.positions[i] = s.positions[i-1]
	}
	s.positions[0] = position
}

func (s *Slug) wrap(position types.Vector2) types.Vector2 {
	switch {
	case position.X < 0:
		return types.Vector2{X: s.gridColumns - 1, Y: position.Y}
	case position.X >= s.gridColumns:
		return types.Vector2{X: 0, Y: position.Y}
	case position.Y < 0:
		return types.Vector2{X: position.X, Y: s.gridRows - 1}
	case position.Y >= s.gridRows:
		return types.Vector2{X: position.X, Y: 0}
	}

	return position
}

func (s *Slug) Position() types.Vector2 {
	return s.positions[0]
}

func (s *Slug) Draw(screen *ebiten.Image, tileSize int, offsetX int, offsetY int) {
	bounds := s.headSprite.Bounds()

	opHead := &ebiten.DrawImageOptions{}
	s.center(opHead, bounds)
	s.rotate(opHead)
	s.scale(tileSize, bounds, opHead)
	s.translate(offsetX, tileSize, offsetY, opHead, s.positions[0])
	screen.DrawImage(s.headSprite, opHead)

	for _, body := range s.positions[1:] {
		opBody := &ebiten.DrawImageOptions{}
		s.center(opBody, bounds)
		s.scale(tileSize, bounds, opBody)
		s.translate(offsetX, tileSize, offsetY, opBody, body)

		isTail := body == s.positions[len(s.positions)-1]
		if !isTail {
			screen.DrawImage(s.bodySprite, opBody)
		} else {
			screen.DrawImage(s.tailSprite, opBody)
		}
	}
}

func (s *Slug) scale(tileSize int, bounds image.Rectangle, op *ebiten.DrawImageOptions) {
	sc := float64(tileSize) / float64(bounds.Dx())
	op.GeoM.Scale(sc, sc)
}

func (s *Slug) translate(offsetX int, tileSize int, offsetY int, op *ebiten.DrawImageOptions, position types.Vector2) {
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

func (s *Slug) center(op *ebiten.DrawImageOptions, bounds image.Rectangle) {
	op.GeoM.Translate(
		-float64(bounds.Dx())/2,
		-float64(bounds.Dy())/2,
	)
}

func (s *Slug) Grow() {
	s.positions = append(s.positions, s.positions[len(s.positions)-1])
}

func (s *Slug) NextPosition() types.Vector2 {
	delta := s.nextDirection.Multiply(s.jumpBy)
	return s.wrap(s.Position().Add(delta))
}

func NewSlug(jumpBy int, startPosition types.Vector2, moveFrequency time.Duration, length, columns, rows int) *Slug {
	positions := make([]types.Vector2, length)
	for i := range positions {
		positions[i] = types.Vector2{X: startPosition.X + i, Y: startPosition.Y}
	}
	return &Slug{
		positions:        positions,
		headSprite:       assets.SlugHeadSprite,
		bodySprite:       assets.SlugBodySprite,
		tailSprite:       assets.SlugTailSprite,
		jumpBy:           jumpBy,
		currentDirection: DirectionLeft,
		nextDirection:    DirectionLeft,
		moveTimer:        lib.NewTimer(moveFrequency, ebiten.TPS()),
		gridColumns:      columns,
		gridRows:         rows,
	}
}
