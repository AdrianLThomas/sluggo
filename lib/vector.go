package lib

type Number interface {
	~int | ~float64
}

type Vector2[T Number] struct {
	X, Y T
}

func (v *Vector2[T]) Add(other Vector2[T]) {
	v.X += other.X
	v.Y += other.Y
}

func (v *Vector2[T]) Multiply(amount T) {
	v.X *= amount
	v.Y *= amount
}
