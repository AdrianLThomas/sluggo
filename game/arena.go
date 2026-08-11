package game

import (
	"math/rand"
	"sluggo/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

type Arena struct {
	columns    int
	rows       int
	slug       *Slug
	bgImage    *ebiten.Image
	bgTileSize int
	food       []*Food
	rock       []*Rock
}

func (a *Arena) Update() error {
	if err := a.slug.Update(); err != nil {
		return err
	}

	a.slug.Wrap()

	for _, food := range a.food {
		if err := food.Update(); err != nil {
			return err
		}

		isCollision := a.slug.Position() == food.position
		if isCollision {
			a.slug.Grow()
			food.Reset(a.randomGridPosition())
		}
	}

	for _, rock := range a.rock {
		if err := rock.Update(); err != nil {
			return err
		}

		isCollision := a.slug.NextPosition() == rock.position
		if isCollision || a.slug.WillEatSelf() {
			IsGameOver = true
		}
	}

	return nil
}

func (a *Arena) Draw(screen *ebiten.Image, tileSize int, offsetX int, offsetY int) {
	if a.bgImage == nil || tileSize != a.bgTileSize {
		a.rebuildBackground(tileSize)
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(offsetX), float64(offsetY))
	screen.DrawImage(a.bgImage, op)

	for _, food := range a.food {
		food.Draw(screen, tileSize, offsetX, offsetY)
	}
	for _, rock := range a.rock {
		rock.Draw(screen, tileSize, offsetX, offsetY)
	}
	a.slug.Draw(screen, tileSize, offsetX, offsetY)
}

func (a *Arena) rebuildBackground(tileSize int) {
	if a.bgImage != nil {
		a.bgImage.Dispose()
	}

	w := a.columns * tileSize
	h := a.rows * tileSize
	a.bgImage = ebiten.NewImage(w, h)

	s := float64(tileSize) / float64(assets.BackgroundSprite.Bounds().Dx())
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(s, s)

	for y := range a.rows {
		for x := range a.columns {
			op.GeoM.SetElement(0, 2, float64(x*tileSize))
			op.GeoM.SetElement(1, 2, float64(y*tileSize))
			a.bgImage.DrawImage(assets.BackgroundSprite, op)
		}
	}

	a.bgTileSize = tileSize
}

func NewArena(columns int, rows int) *Arena {
	foodPos := RandomGridPosition(columns, rows)
	rockPos := RandomGridPosition(columns, rows)
	for foodPos.X == rockPos.X && foodPos.Y == rockPos.Y {
		rockPos = RandomGridPosition(columns, rows)
	}
	return &Arena{
		columns: columns,
		rows:    rows,
		slug: NewSlug(JumpBy, Vector2{
			X: columns - 1,
			Y: rows / 2,
		},
			StartingSlugLength,
			columns,
			rows),
		food: []*Food{
			NewFood(foodPos),
		},
		rock: []*Rock{
			NewRock(rockPos),
		},
	}
}

func (a *Arena) randomGridPosition() Vector2 {
	return RandomGridPosition(a.columns, a.rows)
}

func RandomGridPosition(columns, rows int) Vector2 {
	return Vector2{rand.Intn(columns), rand.Intn(rows)}
}
