package arena

import (
	"math/rand"
	"sluggo/game/characters"
	"sluggo/game/objects"
	"sluggo/types"
	"testing"
	"time"
)

func TestArena_FoodNotPlacedOnRock(t *testing.T) {
	a := NewArena(2, 2)
	a.randGen = rand.New(rand.NewSource(3))
	a.rock[0] = objects.NewRock(types.Vector2{X: 0, Y: 0})
	a.food[0].Reset(a.slug.Head())

	if _, err := a.Update(); err != nil {
		t.Fatalf("Update() error = %v", err)
	}

	if a.food[0].Position() == a.rock[0].Position() {
		t.Fatalf("food placed on rock after being eaten: food=%v rock=%v", a.food[0].Position(), a.rock[0].Position())
	}
}

func TestArena_EatingFoodReportsScoreGain(t *testing.T) {
	a := NewArena(2, 2)
	// NewArena scatters the rock, pin it clear of the slug's path
	a.rock[0] = objects.NewRock(types.Vector2{X: 0, Y: 0})
	a.food[0].Reset(a.slug.Head())

	outcome, err := a.Update()
	if err != nil {
		t.Fatalf("Update() error = %v", err)
	}

	if outcome.ScoreGain != a.slug.Speed() {
		t.Errorf("Update() ScoreGain = %v, want %v", outcome.ScoreGain, a.slug.Speed())
	}

	if outcome.GameOver {
		t.Error("Update() GameOver = true, want false")
	}
}

func TestArena_UpdateWithoutFoodReportsNoScore(t *testing.T) {
	a := NewArena(4, 4)
	a.rock = nil
	a.food[0].Reset(types.Vector2{X: 0, Y: 0})

	outcome, err := a.Update()
	if err != nil {
		t.Fatalf("Update() error = %v", err)
	}

	if outcome.ScoreGain != 0 {
		t.Errorf("Update() ScoreGain = %v, want 0", outcome.ScoreGain)
	}
}

func TestArena_SelfCollisionDetectedWithoutRocks(t *testing.T) {
	const grid = 3
	const length = 4

	a := NewArena(grid, grid)
	a.rock = nil
	a.food[0].Reset(types.Vector2{X: 0, Y: 0})
	// left edge moving left, so the trailing body wraps in front of the head
	a.slug = characters.NewSlug(
		1,
		types.Vector2{X: 0, Y: 1},
		time.Millisecond*100,
		length,
		grid,
		grid,
	)

	if !a.slug.WillEatSelf() {
		t.Fatalf("test setup: WillEatSelf() = false, want true; positions=%v next=%v",
			a.slug.Positions(), a.slug.NextPosition())
	}

	outcome, err := a.Update()
	if err != nil {
		t.Fatalf("Update() error = %v", err)
	}

	if !outcome.GameOver {
		t.Error("Update() GameOver = false, want true")
	}
}

func TestArena_RockCollisionReportsGameOver(t *testing.T) {
	a := NewArena(4, 4)
	a.food[0].Reset(types.Vector2{X: 0, Y: 0})
	a.rock[0].SetPosition(a.slug.NextPosition())

	outcome, err := a.Update()
	if err != nil {
		t.Fatalf("Update() error = %v", err)
	}

	if !outcome.GameOver {
		t.Errorf("Update() GameOver = false, want true; next position %v sits on rock %v",
			a.slug.NextPosition(), a.rock[0].Position())
	}
}

func TestArena_ScoreSurvivesFatalCollisionSameTick(t *testing.T) {
	a := NewArena(4, 4)
	a.food[0].Reset(a.slug.Head())
	a.rock[0].SetPosition(a.slug.NextPosition())

	lengthBefore := len(a.slug.Positions())

	outcome, err := a.Update()
	if err != nil {
		t.Fatalf("Update() error = %v", err)
	}

	if len(a.slug.Positions()) == lengthBefore {
		t.Fatal("Update() did not eat the food, test setup is wrong")
	}

	if !outcome.GameOver {
		t.Fatal("Update() GameOver = false, want true, test setup is wrong")
	}

	if outcome.ScoreGain != a.slug.Speed() {
		t.Errorf("Update() ScoreGain = %v, want %v", outcome.ScoreGain, a.slug.Speed())
	}
}
