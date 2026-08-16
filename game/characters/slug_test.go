package characters

import (
	"sluggo/types"
	"testing"
)

func newTestSlug(startPosition types.Vector2) *Slug {
	jumpBy := 1
	length := 1
	gridSize := 20

	return NewSlug(jumpBy, startPosition, 1, length, gridSize, gridSize)
}

func TestSlug_WillEatSelf(t *testing.T) {
	tests := []struct {
		name             string
		startPosition    types.Vector2
		currentDirection types.Vector2
		nextDirection    types.Vector2
		positions        []types.Vector2
		gridSize         int
		expected         bool
	}{
		{
			name:             "true when next position is part of body",
			startPosition:    types.Vector2{X: 5, Y: 5},
			currentDirection: DirectionRight,
			nextDirection:    DirectionDown,
			positions: []types.Vector2{
				{X: 5, Y: 4},
				{X: 4, Y: 4},
				{X: 4, Y: 5},
				{X: 5, Y: 5},
			},
			expected: true,
		},
		{
			name:             "false when next position is not part of body",
			startPosition:    types.Vector2{X: 5, Y: 5},
			currentDirection: DirectionRight,
			nextDirection:    DirectionUp,
			positions: []types.Vector2{
				{X: 5, Y: 4},
				{X: 4, Y: 4},
				{X: 4, Y: 5},
				{X: 5, Y: 5},
			},
			expected: false,
		},
		{
			name:             "true when next position wraps around grid boundary into body",
			gridSize:         10,
			startPosition:    types.Vector2{X: 9, Y: 5},
			currentDirection: DirectionRight,
			nextDirection:    DirectionRight,
			positions: []types.Vector2{
				{X: 9, Y: 5},
				{X: 8, Y: 5},
				{X: 0, Y: 5},
			},
			expected: true,
		},
		{
			name:             "false when next position wraps around grid boundary into empty space",
			gridSize:         10,
			startPosition:    types.Vector2{X: 9, Y: 5},
			currentDirection: DirectionRight,
			nextDirection:    DirectionRight,
			positions: []types.Vector2{
				{X: 9, Y: 5},
				{X: 8, Y: 5},
				{X: 7, Y: 5},
			},
			expected: false,
		},
		{
			name:             "false for single segment slug",
			startPosition:    types.Vector2{X: 5, Y: 5},
			currentDirection: DirectionRight,
			nextDirection:    DirectionRight,
			positions: []types.Vector2{
				{X: 5, Y: 5},
			},
			expected: false,
		},
		{
			name:             "true when next position is tail of slug",
			startPosition:    types.Vector2{X: 5, Y: 5},
			currentDirection: DirectionUp,
			nextDirection:    DirectionRight,
			positions: []types.Vector2{
				{X: 5, Y: 5},
				{X: 5, Y: 6},
				{X: 6, Y: 6},
				{X: 6, Y: 5},
			},
			expected: true,
		},
		{
			name:             "false when moving parallel adjacent to body",
			startPosition:    types.Vector2{X: 5, Y: 5},
			currentDirection: DirectionUp,
			nextDirection:    DirectionLeft,
			positions: []types.Vector2{
				{X: 5, Y: 5},
				{X: 5, Y: 6},
				{X: 5, Y: 7},
			},
			expected: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := newTestSlug(tt.startPosition)
			if tt.gridSize > 0 {
				s.gridColumns = tt.gridSize
				s.gridRows = tt.gridSize
			}
			s.positions = tt.positions
			s.nextDirection = tt.nextDirection
			s.currentDirection = tt.currentDirection

			got := s.WillEatSelf()

			if got != tt.expected {
				t.Errorf("WillEatSelf() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestSlug_NextPosition(t *testing.T) {
	tests := []struct {
		name             string
		direction        types.Vector2
		startPosition    types.Vector2
		expectedPosition types.Vector2
	}{
		{
			name:             "when direction left, moves x by -1",
			direction:        DirectionLeft,
			startPosition:    types.Vector2{X: 10, Y: 10},
			expectedPosition: types.Vector2{X: 9, Y: 10},
		},
		{
			name:             "when direction right, moves x by +1",
			direction:        DirectionRight,
			startPosition:    types.Vector2{X: 0, Y: 0},
			expectedPosition: types.Vector2{X: 1, Y: 0},
		},
		{
			name:             "when direction down, moves y by +1",
			direction:        DirectionDown,
			startPosition:    types.Vector2{X: 0, Y: 0},
			expectedPosition: types.Vector2{X: 0, Y: 1},
		},
		{
			name:             "when direction up, moves y by -1",
			direction:        DirectionUp,
			startPosition:    types.Vector2{X: 0, Y: 1},
			expectedPosition: types.Vector2{X: 0, Y: 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := newTestSlug(tt.startPosition)
			s.currentDirection = tt.direction
			s.nextDirection = tt.direction

			actualPosition := s.NextPosition()

			if actualPosition != tt.expectedPosition {
				t.Errorf("NextPosition() = %v, want %v", actualPosition, tt.expectedPosition)
			}
		})
	}
}

func TestSlug_wrap(t *testing.T) {
	tests := []struct {
		name     string
		columns  int
		rows     int
		position types.Vector2
		expected types.Vector2
	}{
		{
			name:     "position inside grid remains unchanged",
			columns:  10,
			rows:     10,
			position: types.Vector2{X: 5, Y: 5},
			expected: types.Vector2{X: 5, Y: 5},
		},
		{
			name:     "position at top-left corner remains unchanged",
			columns:  10,
			rows:     10,
			position: types.Vector2{X: 0, Y: 0},
			expected: types.Vector2{X: 0, Y: 0},
		},
		{
			name:     "position at bottom-right corner remains unchanged",
			columns:  10,
			rows:     10,
			position: types.Vector2{X: 9, Y: 9},
			expected: types.Vector2{X: 9, Y: 9},
		},
		{
			name:     "position off left edge (X < 0) wraps to right side",
			columns:  10,
			rows:     10,
			position: types.Vector2{X: -1, Y: 5},
			expected: types.Vector2{X: 9, Y: 5},
		},
		{
			name:     "position off right edge (X >= columns) wraps to left side",
			columns:  10,
			rows:     10,
			position: types.Vector2{X: 10, Y: 5},
			expected: types.Vector2{X: 0, Y: 5},
		},
		{
			name:     "position off top edge (Y < 0) wraps to bottom side",
			columns:  10,
			rows:     10,
			position: types.Vector2{X: 5, Y: -1},
			expected: types.Vector2{X: 5, Y: 9},
		},
		{
			name:     "position off bottom edge (Y >= rows) wraps to top side",
			columns:  10,
			rows:     10,
			position: types.Vector2{X: 5, Y: 10},
			expected: types.Vector2{X: 5, Y: 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := newTestSlug(types.Vector2{X: 0, Y: 0})
			s.gridColumns = tt.columns
			s.gridRows = tt.rows

			actual := s.wrap(tt.position)
			if actual != tt.expected {
				t.Errorf("wrap() = %v, want %v", actual, tt.expected)
			}
		})
	}
}
