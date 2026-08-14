package main

import (
	"log"
	"sluggo/game"
	"sluggo/types"

	"github.com/hajimehoshi/ebiten/v2"
)

var ScreenSize = types.Vector2{X: 1024, Y: 768}

func main() {
	ebiten.SetWindowSize(ScreenSize.X, ScreenSize.Y)
	ebiten.SetWindowTitle("Sluggo!")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeOnlyFullscreenEnabled)

	g := game.NewGame(20, 15)

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
