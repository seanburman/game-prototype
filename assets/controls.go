package assets

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/seanburman/game/constants"
)

type ControllerButton int

const (
	DPadUp ControllerButton = iota
	DPadDown
	DPadLeft
	DPadRight
	ButtonA
	ButtonB
)

var TouchIDs = make([]ebiten.TouchID, 1024)
var Touches = make(map[ebiten.TouchID]*Touch, 1024)

type Touch struct {
	OriginX, OriginY int
	CurrX, CurrY     int
	Duration         int
}

type Controls struct {
	Sprites []*Sprite
	Touches map[ebiten.TouchID]*Touch
	Debug   string
	Collider
	Vector
	Name
	Sprite
}

func NewControls(x int, y int) *Controls {
	c := &Controls{}
	c.Sprites = []*Sprite{
		NewSprite(c.Button(DPadUp), NewSpriteOpts{
			ImgSource: "square/img/square1.png",
		}),
		NewSprite(c.Button(DPadDown), NewSpriteOpts{
			ImgSource: "square/img/square1.png",
		}),
		NewSprite(c.Button(DPadLeft), NewSpriteOpts{
			ImgSource: "square/img/square1.png",
		}),
		NewSprite(c.Button(DPadRight), NewSpriteOpts{
			ImgSource: "square/img/square1.png",
		}),
		NewSprite(c.Button(ButtonA), NewSpriteOpts{
			ImgSource: "square/img/square1.png",
		}),
		NewSprite(c.Button(ButtonB), NewSpriteOpts{
			ImgSource: "square/img/square1.png",
		}),
	}
	c.Sprites[DPadUp].SetPosition(x-16, y-60)
	c.Sprites[DPadDown].SetPosition(x-16, y+28)
	c.Sprites[DPadLeft].SetPosition(x-60, y-16)
	c.Sprites[DPadRight].SetPosition(x+28, y-16)
	c.Sprites[ButtonA].SetPosition(x+230, y-42)
	c.Sprites[ButtonB].SetPosition(x+168, y+12)
	return c
}

func (c Controls) Draw(screen *ebiten.Image) {
	controllerBG := ebiten.NewImage(constants.GetDimensions(), 497)
	cbgOpts := &ebiten.DrawImageOptions{}
	controllerBG.Fill(color.Black)
	cbgOpts.GeoM.Translate(0, float64(constants.GetDimensions())+7)
	screen.DrawImage(controllerBG, cbgOpts)

	for _, s := range c.Sprites {
		s.Draw(screen)
	}
	ebitenutil.DebugPrint(screen, c.Debug)
}

func (c *Controls) Update() {
	for id := range Touches {
		if inpututil.IsTouchJustReleased(id) {
			delete(Touches, id)
		}
	}

	TouchIDs = inpututil.AppendJustPressedTouchIDs(TouchIDs)
	for _, id := range TouchIDs {
		x, y := ebiten.TouchPosition(id)
		Touches[id] = &Touch{
			OriginX: x, OriginY: y,
			CurrX: x, CurrY: y,
		}

		char := CharacterRegistry.Sprites[0]
		for _, s := range c.Sprites {
			n := string(s.GetName())
			if s.IsPressed(id) && n == c.Button(ButtonB) {
				char.SetSpeed(5)
			}
			switch n {
			case c.Button(DPadUp):
				if s.IsPressed(id) && !char.GetBreached().Min.Y {
					char.Move(0, -1)
				}
			case c.Button(DPadDown):
				if s.IsPressed(id) && !char.GetBreached().Max.Y {
					char.Move(0, 1)
				}
			case c.Button(DPadRight):
				if s.IsPressed(id) && !char.GetBreached().Max.X {
					char.Move(1, 0)
				}
			case c.Button(DPadLeft):
				if s.IsPressed(id) && !char.GetBreached().Min.X {
					char.Move(-1, 0)
				}
			default:
			}
		}
	}
}

func (c Controls) Button(cb ControllerButton) string {
	switch cb {
	case DPadUp:
		return "DPadUp"
	case DPadDown:
		return "DpadDown"
	case DPadLeft:
		return "DpadLeft"
	case DPadRight:
		return "DpadRight"
	case ButtonA:
		return "ButtonA"
	case ButtonB:
		return "ButtonB"
	default:
		return ""
	}
}

func (b *Controls) Move(x int, y int)    {}
func (b Controls) GetBreached() Breached { return Breached{} }
