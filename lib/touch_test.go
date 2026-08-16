package lib

import (
	"testing"
)

func TestGetDirectionFromDelta(t *testing.T) {
	tests := []struct {
		name     string
		dx, dy   float64
		expected Vector2[int]
	}{
		{"Swipe Right", 50, 5, DirectionRight},
		{"Swipe Left", -50, 5, DirectionLeft},
		{"Swipe Up", 5, -50, DirectionUp},
		{"Swipe Down", 5, 50, DirectionDown},
		{"Diagonal Right-Down (more right)", 50, 40, DirectionRight},
		{"Diagonal Down-Right (more down)", 40, 50, DirectionDown},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetDirectionFromDelta(tt.dx, tt.dy)
			if got != tt.expected {
				t.Errorf("GetDirectionFromDelta(%v, %v) = %v, want %v", tt.dx, tt.dy, got, tt.expected)
			}
		})
	}
}

func TestGetDirectionFromTap(t *testing.T) {
	width := 1000
	height := 800
	// Center is (500, 400)

	tests := []struct {
		name       string
		tapX, tapY int
		expected   Vector2[int]
	}{
		{"Tap Right Sector", 800, 400, DirectionRight},
		{"Tap Left Sector", 100, 400, DirectionLeft},
		{"Tap Top Sector", 500, 100, DirectionUp},
		{"Tap Bottom Sector", 500, 700, DirectionDown},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetDirectionFromTap(tt.tapX, tt.tapY, width, height)
			if got != tt.expected {
				t.Errorf("GetDirectionFromTap(%v, %v) = %v, want %v", tt.tapX, tt.tapY, got, tt.expected)
			}
		})
	}
}
