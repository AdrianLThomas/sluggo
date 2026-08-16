package objects

import (
	"sluggo/assets"
	"sluggo/types"
)

type Rock struct {
	GameObject
}

func NewRock(position types.Vector2) *Rock {
	return &Rock{
		GameObject: GameObject{
			position: position,
			sprite:   assets.RockSprite,
		},
	}
}
