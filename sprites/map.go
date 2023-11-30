package sprites

import (
	"fmt"
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type (
	GameMap struct {
		*Sprite
	}
)

func NewGameMap(x int, y int) *GameMap {
	// src := "tilemaps/ground/green.png"
	s := ebiten.NewImage(768, 768)
	g := &GameMap{
		NewSprite("grass", NewSpriteOpts{
			// ImgSource: &src,
			Shape: s,
		}),
	}
	b := g.Sprite.Collider.Bounds
	g.Sprite.SetBounds(Bounds{
		Min: image.Point{X: x, Y: y},
		Max: image.Point{X: b.Max.X, Y: b.Max.Y},
	})

	s.Fill(color.RGBA{R: 0, G: 135, B: 81, A: 1})
	g.Sprite.SetPosition(x, y)
	return g
}

func (mg *GameMap) Update() {
	p := Registry.Player

	mg.SetSpeed(p.GetSpeed())
	mg.SetVelocity(p.GetVelocity())

	coords := XY{}
	if p.GetVelocity() > 0 {
		b := mg.Breached
		fmt.Println(b)
		// Move Up
		if !b.Min.Y && p.d.Y == -1 {
			coords = XY{0, 1}
		}
		// Move Down
		if !b.Max.Y && p.d.Y == 1 {
			coords = XY{0, -1}
		}
		// Move Left
		if !b.Min.X && p.d.X == -1 {
			coords = XY{1, 0}
		}
		// Move Right
		if !b.Max.X && p.d.X == 1 {
			coords = XY{-1, 0}
		}
	}
	mg.Move(coords.X, coords.Y)

	for _, s := range Registry.Props {
		s.SetSpeed(p.GetSpeed())
		s.SetVelocity(p.GetVelocity())
		s.Move(coords.X, coords.Y)
	}
}

func (mg *GameMap) Draw(screen *ebiten.Image) {
	mg.Sprite.Draw(screen)
}

func (mg *GameMap) Move(x int, y int) {
	mg.Sprite.Move(x, y)
}
