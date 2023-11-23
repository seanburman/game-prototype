package config

import (
	"image"

	"github.com/seanburman/game/assets"
)

var GAME_BOUNDARIES = assets.NewCollider(assets.Bounds{
	Min: image.Point{X: 0, Y: 0},
	Max: image.Point{X: 384, Y: 384},
})
