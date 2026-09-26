package game

import "testing"

func newTestGame(t *testing.T) *game {
	t.Helper()

	g, ok := NewGame(20, 15).(*game)
	if !ok {
		t.Fatal("expected NewGame to return a *game")
	}
	return g
}

func TestNewGameStartsOnSplash(t *testing.T) {
	g := newTestGame(t)

	if g.state != StateSplash {
		t.Errorf("expected a new game to start on %v, got %v", StateSplash, g.state)
	}
}

// The splash screen must never advance on its own, however long it is left up.
func TestSplashWaitsForInput(t *testing.T) {
	g := newTestGame(t)

	for i := range 600 {
		if err := g.Update(); err != nil {
			t.Fatal(err)
		}
		if g.state != StateSplash {
			t.Fatalf("left the splash screen without input on tick %d", i+1)
		}
	}
}

// Game over must likewise wait for input rather than restarting by itself.
func TestGameOverWaitsForInput(t *testing.T) {
	g := newTestGame(t)
	g.state = StateGameOver

	for i := range 600 {
		if err := g.Update(); err != nil {
			t.Fatal(err)
		}
		if g.state != StateGameOver {
			t.Fatalf("left game over without input on tick %d", i+1)
		}
	}
}

func TestResetResetsScoreAndArena(t *testing.T) {
	g := newTestGame(t)
	g.score = 42
	g.state = StateGameOver

	before := g.arena
	g.reset()

	if g.state != StatePlaying {
		t.Errorf("expected reset to start playing, got %v", g.state)
	}
	if g.score != 0 {
		t.Errorf("expected reset to zero the score, got %d", g.score)
	}
	if g.arena == before {
		t.Error("expected reset to build a fresh arena")
	}
}
