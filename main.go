package main

import (
	"log"
	"sluggo/game"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(game.ScreenSize.X, game.ScreenSize.Y)
	ebiten.SetWindowTitle("Sluggo!")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeOnlyFullscreenEnabled)

	g := game.NewGame(10, 7)

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
