package arena

import (
	"math/rand"
	"sluggo/game/objects"
	"sluggo/types"
	"testing"
)

func TestArena_FoodNotPlacedOnRock(t *testing.T) {
	a := NewArena(2, 2, func() {})
	a.randGen = rand.New(rand.NewSource(0))
	a.rock[0] = objects.NewRock(types.Vector2{X: 0, Y: 0})
	a.food[0].Reset(a.slug.Head())

	if err := a.Update(); err != nil {
		t.Fatalf("Update() error = %v", err)
	}

	if a.food[0].Position() == a.rock[0].Position() {
		t.Fatalf("food placed on rock after being eaten: food=%v rock=%v", a.food[0].Position(), a.rock[0].Position())
	}
}
