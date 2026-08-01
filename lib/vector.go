package lib

type Vector2D struct {
	X, Y float64
}

func (v *Vector2D) Add(X, Y float64) {
	v.X += X
	v.Y += Y
}
