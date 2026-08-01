package lib

type Vector2D struct {
	X, Y float64
}

func (v *Vector2D) Add(other *Vector2D) {
	v.X += other.X
	v.Y += other.Y
}

func (v *Vector2D) Multiply(amount float64) {
	v.X *= amount
	v.Y *= amount
}
