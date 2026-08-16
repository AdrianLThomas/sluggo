package objects

import (
	"sluggo/assets"
	"sluggo/types"
)

type Food struct {
	GameObject
}

func (f *Food) Reset(position types.Vector2) {
	f.SetPosition(position)
}

func NewFood(position types.Vector2) *Food {
	return &Food{
		GameObject: GameObject{
			position: position,
			sprite:   assets.FoodSprite,
		},
	}
}
