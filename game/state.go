package game

type GameState int

const (
	StatePlaying GameState = iota
	StateGameOver
	StateSplash
)

var stateName = map[GameState]string{
	StatePlaying:  "Playing",
	StateGameOver: "Game Over",
	StateSplash:   "Splash",
}

func (gs GameState) String() string {
	return stateName[gs]
}
