package lib

import (
	"math/rand"
	"testing"
)

var sinkInt int
var sinkFloat float64

var vecIntInputs []Vector2[int]
var vecFloatInputs []Vector2[float64]
var intInputs []int
var floatInputs []float64

func init() {
	const n = 1024
	rng := rand.New(rand.NewSource(99))
	vecIntInputs = make([]Vector2[int], n)
	vecFloatInputs = make([]Vector2[float64], n)
	intInputs = make([]int, n)
	floatInputs = make([]float64, n)
	for i := range n {
		vecIntInputs[i] = Vector2[int]{X: rng.Int(), Y: rng.Int()}
		vecFloatInputs[i] = Vector2[float64]{X: rng.Float64(), Y: rng.Float64()}
		intInputs[i] = rng.Int()
		floatInputs[i] = rng.Float64()
	}
}

func TestVector2IntAdd(t *testing.T) {
	v := Vector2[int]{X: 1, Y: 2}
	v = v.Add(Vector2[int]{X: 3, Y: 4})
	if v.X != 4 || v.Y != 6 {
		t.Errorf("expected (4, 6), got (%d, %d)", v.X, v.Y)
	}
}

func TestVector2Float64Add(t *testing.T) {
	v := Vector2[float64]{X: 1.5, Y: 2.5}
	v = v.Add(Vector2[float64]{X: 3.5, Y: 4.5})
	if v.X != 5.0 || v.Y != 7.0 {
		t.Errorf("expected (5, 7), got (%f, %f)", v.X, v.Y)
	}
}

func TestVector2IntMultiply(t *testing.T) {
	v := Vector2[int]{X: 2, Y: 3}
	v = v.Multiply(4)
	if v.X != 8 || v.Y != 12 {
		t.Errorf("expected (8, 12), got (%d, %d)", v.X, v.Y)
	}
}

func TestVector2Float64Multiply(t *testing.T) {
	v := Vector2[float64]{X: 2.5, Y: 3.5}
	v = v.Multiply(2.0)
	if v.X != 5.0 || v.Y != 7.0 {
		t.Errorf("expected (5, 7), got (%f, %f)", v.X, v.Y)
	}
}

func TestVector2IntAddZero(t *testing.T) {
	v := Vector2[int]{X: 5, Y: 10}
	v = v.Add(Vector2[int]{})
	if v.X != 5 || v.Y != 10 {
		t.Errorf("expected (5, 10), got (%d, %d)", v.X, v.Y)
	}
}

func TestVector2IntMultiplyByZero(t *testing.T) {
	v := Vector2[int]{X: 5, Y: 10}
	v = v.Multiply(0)
	if v.X != 0 || v.Y != 0 {
		t.Errorf("expected (0, 0), got (%d, %d)", v.X, v.Y)
	}
}

func BenchmarkVector2IntAdd(b *testing.B) {
	v := Vector2[int]{}
	for i := 0; i < b.N; i++ {
		v = v.Add(vecIntInputs[i&1023])
		sinkInt = v.X
	}
}

func BenchmarkVector2Float64Add(b *testing.B) {
	v := Vector2[float64]{}
	for i := 0; i < b.N; i++ {
		v = v.Add(vecFloatInputs[i&1023])
		sinkFloat = v.X
	}
}

func BenchmarkVector2IntMultiply(b *testing.B) {
	v := Vector2[int]{X: 1, Y: 2}
	for i := 0; i < b.N; i++ {
		v = v.Multiply(intInputs[i&1023])
		sinkInt = v.X
	}
}

func BenchmarkVector2Float64Multiply(b *testing.B) {
	v := Vector2[float64]{X: 1, Y: 2}
	for i := 0; i < b.N; i++ {
		v = v.Multiply(floatInputs[i&1023])
		sinkFloat = v.X
	}
}
