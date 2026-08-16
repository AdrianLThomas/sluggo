package arena

import (
	"fmt"
	"math/rand"
	"slices"
	"sluggo/assets"
	"sluggo/game/characters"
	"sluggo/game/objects"
	"sluggo/types"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Arena struct {
	columns    int
	rows       int
	slug       *characters.Slug
	bgImage    *ebiten.Image
	bgTileSize int
	food       []*objects.Food
	rock       []*objects.Rock
	randGen    *rand.Rand
	onGameOver func()
}

func (a *Arena) Update() error {
	if err := a.slug.Update(); err != nil {
		return err
	}

	for _, food := range a.food {
		if err := food.Update(); err != nil {
			return err
		}

		isCollision := a.slug.Head() == food.Position()
		if isCollision {
			a.slug.Grow()
			food.Reset(a.nonCollidingPosition())
		}
	}

	for _, rock := range a.rock {
		if err := rock.Update(); err != nil {
			return err
		}

		isCollision := a.slug.NextPosition() == rock.Position()
		if isCollision || a.slug.WillEatSelf() {
			a.onGameOver()
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

	ebitenutil.DebugPrintAt(screen,
		fmt.Sprintf("X: %v,Y: %v", a.slug.Head().X, a.slug.Head().Y),
		0, screen.Bounds().Dy()-20)
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

func NewArena(columns int, rows int, onGameOver func()) *Arena {
	foodPos := newRandomGridPosition(columns, rows)
	rockPos := newRandomGridPosition(columns, rows)
	for foodPos.X == rockPos.X && foodPos.Y == rockPos.Y {
		rockPos = newRandomGridPosition(columns, rows)
	}
	const jumpBy = 1
	const moveFrequency = time.Millisecond * 100
	const startingSlugLength = 1
	slug := characters.NewSlug(
		jumpBy,
		types.Vector2{X: columns - 1, Y: rows / 2},
		moveFrequency,
		startingSlugLength,
		columns,
		rows)

	return &Arena{
		columns: columns,
		rows:    rows,
		slug:    slug,
		food: []*objects.Food{
			objects.NewFood(foodPos),
		},
		rock: []*objects.Rock{
			objects.NewRock(rockPos),
		},
		randGen:    rand.New(rand.NewSource(time.Now().UnixNano())),
		onGameOver: onGameOver,
	}
}

func (a *Arena) randomGridPosition() types.Vector2 {
	return types.Vector2{X: a.randGen.Intn(a.columns), Y: a.randGen.Intn(a.rows)}
}

func (a *Arena) nonCollidingPosition() types.Vector2 {
	emptyPosition := types.Vector2{}

	s, f, r := true, true, true
	for s || f || r {
		s = slices.Contains(a.slug.Positions(), emptyPosition)
		f = slices.ContainsFunc(a.food, func(food *objects.Food) bool {
			return emptyPosition == food.Position()
		})
		r = slices.ContainsFunc(a.rock, func(rock *objects.Rock) bool {
			return emptyPosition == rock.Position()
		})
		emptyPosition = a.randomGridPosition()
	}

	return emptyPosition
}

func newRandomGridPosition(columns, rows int) types.Vector2 {
	return types.Vector2{X: rand.Intn(columns), Y: rand.Intn(rows)}
}
