package square

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/seanburman/game/assets"
	"github.com/seanburman/game/config"
)

type SquareCharacter struct {
	*assets.Sprite
}

func NewSquare(x int, y int) *SquareCharacter {
	sq := &SquareCharacter{
		assets.NewSprite("character", assets.NewSpriteOpts{
			ImgSource: "square/img/square1.png",
		}),
	}
	sq.SetPosition(x, y)
	sq.SetSpeed(2)
	return sq
}

func (sc *SquareCharacter) Update() {
	sc.Collider.Update()

	for _, p := range assets.PropRegistry.Sprites {
		if p != nil {
			sc.DetectCollision(sc.Vector, p)
		}
	}
	for _, p := range assets.PropRegistry.Sprites {
		if p != nil {
			sc.DetectCollision(sc.Vector, p)
		}
	}

	sc.DetectBoundary(sc.Vector, config.GAME_BOUNDARIES)
	sc.CheckKeyBoardInput()
}

func (sc *SquareCharacter) CheckKeyBoardInput() {
	// Check keyboard input
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		if !sc.GetBreached().Min.Y {
			sc.Move(0, -1)
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		if !sc.GetBreached().Max.Y {
			sc.Move(0, 1)
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		if !sc.GetBreached().Min.X {
			sc.Move(-1, 0)
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		if !sc.GetBreached().Max.X {
			sc.Move(1, 0)
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		sc.SetSpeed(5)
	}
	if !ebiten.IsKeyPressed(ebiten.KeySpace) {
		sc.SetSpeed(2)
	}
}
