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
