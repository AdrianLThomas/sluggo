package game

type GameState int

const (
	StatePlaying GameState = iota
	StateGameOver
)

var stateName = map[GameState]string{
	StatePlaying:  "Playing",
	StateGameOver: "Game Over",
}

func (gs GameState) String() string {
	return stateName[gs]
}
