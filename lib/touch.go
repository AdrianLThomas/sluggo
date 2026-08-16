package lib

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var (
	DirectionLeft  = Vector2[int]{X: -1, Y: 0}
	DirectionUp    = Vector2[int]{X: 0, Y: -1}
	DirectionRight = Vector2[int]{X: 1, Y: 0}
	DirectionDown  = Vector2[int]{X: 0, Y: 1}
)

type touchPoint struct {
	startX int
	startY int
	lastX  int
	lastY  int
	swiped bool
}

type TouchInput struct {
	touches          map[ebiten.TouchID]*touchPoint
	minSwipeDistance float64
}

func NewTouchInput() *TouchInput {
	return &TouchInput{
		touches:          make(map[ebiten.TouchID]*touchPoint),
		minSwipeDistance: 15.0,
	}
}

// Update processes current touch events and returns a direction vector and true if a swipe or tap was detected.
func (t *TouchInput) Update() (Vector2[int], bool) {
	// Register new touches
	justPressed := inpututil.AppendJustPressedTouchIDs(nil)
	for _, id := range justPressed {
		x, y := ebiten.TouchPosition(id)
		t.touches[id] = &touchPoint{
			startX: x,
			startY: y,
			lastX:  x,
			lastY:  y,
			swiped: false,
		}
	}

	// Update active touches
	activeIDs := ebiten.AppendTouchIDs(nil)
	activeSet := make(map[ebiten.TouchID]bool, len(activeIDs))
	var detectedDir Vector2[int]
	var foundDir bool

	for _, id := range activeIDs {
		activeSet[id] = true
		tp, exists := t.touches[id]
		if !exists {
			x, y := ebiten.TouchPosition(id)
			tp = &touchPoint{startX: x, startY: y, lastX: x, lastY: y}
			t.touches[id] = tp
		}

		x, y := ebiten.TouchPosition(id)
		tp.lastX = x
		tp.lastY = y

		dx := float64(x - tp.startX)
		dy := float64(y - tp.startY)
		dist := math.Hypot(dx, dy)

		if dist >= t.minSwipeDistance {
			detectedDir = GetDirectionFromDelta(dx, dy)
			foundDir = true
			tp.swiped = true
			// Reset start position to current position for continuous swipe dragging
			tp.startX = x
			tp.startY = y
		}
	}

	// Check released touches for tap gestures
	for id, tp := range t.touches {
		if !activeSet[id] {
			if !tp.swiped && !foundDir {
				w, h := ebiten.WindowSize()
				detectedDir = GetDirectionFromTap(tp.lastX, tp.lastY, w, h)
				foundDir = true
			}
			delete(t.touches, id)
		}
	}

	return detectedDir, foundDir
}

// GetDirectionFromDelta calculates a cardinal direction vector from delta movement.
func GetDirectionFromDelta(dx, dy float64) Vector2[int] {
	if math.Abs(dx) > math.Abs(dy) {
		if dx > 0 {
			return DirectionRight
		}
		return DirectionLeft
	}
	if dy > 0 {
		return DirectionDown
	}
	return DirectionUp
}

// GetDirectionFromTap calculates a cardinal direction vector based on tap position relative to screen center.
func GetDirectionFromTap(tapX, tapY, width, height int) Vector2[int] {
	if width <= 0 {
		width = 1024
	}
	if height <= 0 {
		height = 768
	}

	centerX := width / 2
	centerY := height / 2

	dx := float64(tapX - centerX)
	dy := float64(tapY - centerY)

	return GetDirectionFromDelta(dx, dy)
}
