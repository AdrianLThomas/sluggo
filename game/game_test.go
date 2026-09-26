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
	newTestGame(t)

	if gameState != StateSplash {
		t.Errorf("expected a new game to start on %v, got %v", StateSplash, gameState)
	}
}

// The splash screen must never advance on its own, however long it is left up.
func TestSplashWaitsForInput(t *testing.T) {
	g := newTestGame(t)

	for i := range 600 {
		if err := g.Update(); err != nil {
			t.Fatal(err)
		}
		if gameState != StateSplash {
			t.Fatalf("left the splash screen without input on tick %d", i+1)
		}
	}
}

// Game over must likewise wait for input rather than restarting by itself.
func TestGameOverWaitsForInput(t *testing.T) {
	g := newTestGame(t)
	gameState = StateGameOver

	for i := range 600 {
		if err := g.Update(); err != nil {
			t.Fatal(err)
		}
		if gameState != StateGameOver {
			t.Fatalf("left game over without input on tick %d", i+1)
		}
	}
}

func TestRestartResetsScoreAndArena(t *testing.T) {
	g := newTestGame(t)
	g.score = 42
	gameState = StateGameOver

	before := g.arena
	g.restart()

	if gameState != StatePlaying {
		t.Errorf("expected restart to start playing, got %v", gameState)
	}
	if g.score != 0 {
		t.Errorf("expected restart to zero the score, got %d", g.score)
	}
	if g.arena == before {
		t.Error("expected restart to build a fresh arena")
	}
}
