package lib

import "testing"

func TestVector2IntAdd(t *testing.T) {
	v := Vector2[int]{X: 1, Y: 2}
	v.Add(Vector2[int]{X: 3, Y: 4})
	if v.X != 4 || v.Y != 6 {
		t.Errorf("expected (4, 6), got (%d, %d)", v.X, v.Y)
	}
}

func TestVector2Float64Add(t *testing.T) {
	v := Vector2[float64]{X: 1.5, Y: 2.5}
	v.Add(Vector2[float64]{X: 3.5, Y: 4.5})
	if v.X != 5.0 || v.Y != 7.0 {
		t.Errorf("expected (5, 7), got (%f, %f)", v.X, v.Y)
	}
}

func TestVector2IntMultiply(t *testing.T) {
	v := Vector2[int]{X: 2, Y: 3}
	v.Multiply(4)
	if v.X != 8 || v.Y != 12 {
		t.Errorf("expected (8, 12), got (%d, %d)", v.X, v.Y)
	}
}

func TestVector2Float64Multiply(t *testing.T) {
	v := Vector2[float64]{X: 2.5, Y: 3.5}
	v.Multiply(2.0)
	if v.X != 5.0 || v.Y != 7.0 {
		t.Errorf("expected (5, 7), got (%f, %f)", v.X, v.Y)
	}
}

func TestVector2IntAddZero(t *testing.T) {
	v := Vector2[int]{X: 5, Y: 10}
	v.Add(Vector2[int]{})
	if v.X != 5 || v.Y != 10 {
		t.Errorf("expected (5, 10), got (%d, %d)", v.X, v.Y)
	}
}

func TestVector2IntMultiplyByZero(t *testing.T) {
	v := Vector2[int]{X: 5, Y: 10}
	v.Multiply(0)
	if v.X != 0 || v.Y != 0 {
		t.Errorf("expected (0, 0), got (%d, %d)", v.X, v.Y)
	}
}
