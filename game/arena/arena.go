package arena

import (
	"math/rand"
	"slices"
	"sluggo/assets"
	"sluggo/game/characters"
	"sluggo/game/objects"
	"sluggo/types"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Outcome struct {
	ScoreGain int
	GameOver  bool
}

type Arena struct {
	columns    int
	rows       int
	slug       *characters.Slug
	bgImage    *ebiten.Image
	bgTileSize int
	food       []*objects.Food
	rock       []*objects.Rock
	randGen    *rand.Rand
}

func (a *Arena) Update() (Outcome, error) {
	var outcome Outcome

	if err := a.slug.Update(); err != nil {
		return Outcome{}, err
	}

	for _, food := range a.food {
		if err := food.Update(); err != nil {
			return Outcome{}, err
		}

		isCollision := a.slug.Head() == food.Position()
		if isCollision {
			a.slug.Grow()

			outcome.ScoreGain += a.slug.Speed()
			food.Reset(a.nonCollidingPosition())
		}
	}

	for _, rock := range a.rock {
		if err := rock.Update(); err != nil {
			return Outcome{}, err
		}
	}

	outcome.GameOver = a.slug.WillEatSelf() || a.willHitRock()

	return outcome, nil
}

func (a *Arena) willHitRock() bool {
	return slices.ContainsFunc(a.rock, func(rock *objects.Rock) bool {
		return a.slug.NextPosition() == rock.Position()
	})
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
		a.bgImage.Deallocate()
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
		randGen: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (a *Arena) randomGridPosition() types.Vector2 {
	return types.Vector2{X: a.randGen.Intn(a.columns), Y: a.randGen.Intn(a.rows)}
}

func (a *Arena) nonCollidingPosition() types.Vector2 {
	occupied := func(candidate types.Vector2) bool {
		if slices.Contains(a.slug.Positions(), candidate) {
			return true
		}
		for _, food := range a.food {
			if candidate == food.Position() {
				return true
			}
		}
		for _, rock := range a.rock {
			if candidate == rock.Position() {
				return true
			}
		}
		return false
	}

	candidate := a.randomGridPosition()
	for occupied(candidate) {
		candidate = a.randomGridPosition()
	}

	return candidate
}

func newRandomGridPosition(columns, rows int) types.Vector2 {
	return types.Vector2{X: rand.Intn(columns), Y: rand.Intn(rows)}
}
