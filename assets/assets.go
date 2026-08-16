package assets

import (
	"bytes"
	"embed"
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

//go:embed *.png *.ttf
var assets embed.FS

var BackgroundSprite = mustLoadImage("background_tile.png")
var Font = mustLoadFont("font.ttf")
var SlugHeadSprite = mustLoadImage("slug_head.png")
var SlugBodySprite = mustLoadImage("slug_body.png")
var SlugTailSprite = mustLoadImage("slug_tail.png")
var FoodSprite = mustLoadImage("food.png")
var RockSprite = mustLoadImage("rock.png")

func mustLoadImage(name string) *ebiten.Image {
	f, err := assets.Open(name)
	if err != nil {
		panic(err)
	}
	defer func() { _ = f.Close() }()

	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}

	return ebiten.NewImageFromImage(img)
}

func mustLoadFont(name string) text.Face {
	f, err := assets.ReadFile(name)
	if err != nil {
		panic(err)
	}

	src, err := text.NewGoTextFaceSource(bytes.NewReader(f))
	if err != nil {
		panic(err)
	}

	return &text.GoTextFace{
		Source: src,
		Size:   48,
	}
}

