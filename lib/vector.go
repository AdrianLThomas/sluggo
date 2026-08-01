package lib

type Number interface {
	~int | ~float64
}

type Vector2[T Number] struct {
	X, Y T
}

func (v Vector2[T]) Add(other Vector2[T]) Vector2[T] {
	return Vector2[T]{X: v.X + other.X, Y: v.Y + other.Y}
}

func (v Vector2[T]) Multiply(amount T) Vector2[T] {
	return Vector2[T]{X: v.X * amount, Y: v.Y * amount}
}
