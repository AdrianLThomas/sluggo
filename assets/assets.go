package assets

import (
	"embed"
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed *.png
var assets embed.FS

var SlugHeadSprite = mustLoadImage("slug_head.png")
var SlugBodySprite = mustLoadImage("slug_body.png")
var SlugTailSprite = mustLoadImage("slug_tail.png")
var BackgroundSprite = mustLoadImage("background_tile.png")

func mustLoadImage(name string) *ebiten.Image {
	f, err := assets.Open(name)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}

	return ebiten.NewImageFromImage(img)
}
