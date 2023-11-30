package sprites

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	*Sprite
}

func NewPlayer(x int, y int) *Player {
	src := "characters/square1.png"
	sq := &Player{
		Sprite: NewSprite("character", NewSpriteOpts{
			ImgSource: &src,
		}),
	}
	sq.SetPosition(x, y)
	sq.SetSpeed(2)
	return sq
}

func (p *Player) Update() {
	p.Collider.Update()
	p.Vector.Update()

	for _, p := range Registry.Props {
		if p != nil {
			p.Collider.DetectCollision(p.Vector, Registry.GameMap.Sprite)
		}
	}

	// Make grounds aware of player vector bounds breaching
	// p.DetectBoundary(p.Vector, &Registry.GameMap.Collider)
	p.DetectMovingBoundary(p.Vector, Registry.GameMap.Sprite)

	p.CheckKeyBoardInput()
}

// Implementation of Player move is dependant on the map.
// If the map is static, then the Player Sprite moves.
// If the map is moving to keep the Player in center screen,
// then the Player's relative Vector is updated.
func (p *Player) Move(x int, y int) {
	p.SetDirection(x, y)
	p.SetVelocity(p.GetSpeed())
	// If map is static, character sprite should move
}

func (p *Player) CheckKeyBoardInput() {
	// Check keyboard input
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		if !p.GetBreached().Min.Y {
			p.Move(0, -1)
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		if !p.GetBreached().Max.Y {
			p.Move(0, 1)
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		if !p.GetBreached().Min.X {
			p.Move(-1, 0)
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		if !p.GetBreached().Max.X {
			p.Move(1, 0)
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		p.SetSpeed(4)
	}
	if !ebiten.IsKeyPressed(ebiten.KeySpace) {
		p.SetSpeed(2)
	}
}
