package sprites

import (
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/seanburman/game/constants"
)

// Methods for every child of Sprite
// type SpriteInterface interface {
// 	Update()
// 	Draw(screen *ebiten.Image)
// 	SetName(n string)
// 	IsPressed(tID ebiten.TouchID) bool
// 	Move(x int, y int)
// 	GetDirection() direction
// 	GetPosition() position
// 	GetSpeed() int
// 	GetVector() Vector
// 	SetDirection(x int, y int)
// 	SetPosition(x int, y int)
// 	SetSpeed(s int)
// 	DetectBoundary(v *Vector, foreign *Collider)
// 	DetectCollision(v Vector, f SpriteInterface) *Collider
// 	GetBounds() Bounds
// 	GetBreached() Breached
// 	GetCollider() Collider
// }

var Registry = NewSpriteRegistry()

type (
	Sprite struct {
		Name string
		*ebiten.Image
		Vector
		Collider
	}
	SpriteRegistry struct {
		Player  *Player
		GameMap *GameMap
		Props   []*Sprite
		// Tiles   *Registry
		Input *Controls
	}
	NewSpriteOpts struct {
		Name string
		// Either provided source or Image
		ImgSource *string
		Shape     *ebiten.Image
		Vector    Vector
	}
)

func NewSprite(name string, opts NewSpriteOpts) *Sprite {
	var img *ebiten.Image
	if opts.ImgSource != nil {
		src, _, err := ebitenutil.NewImageFromFile("assets/img/sprites/" + *opts.ImgSource)
		if err != nil {
			log.Fatal(err)
		}
		img = src
	} else if opts.Shape != nil {
		img = opts.Shape
	} else {
		log.Fatalf("image required for sprite: %s", name)
	}

	c := NewCollider(Bounds{
		Min: image.Point{X: img.Bounds().Min.X, Y: img.Bounds().Min.Y},
		Max: image.Point{X: img.Bounds().Max.X, Y: img.Bounds().Max.Y},
	})

	return &Sprite{
		Name:     name,
		Image:    img,
		Vector:   opts.Vector,
		Collider: c,
	}
}

func (spr *Sprite) Update() {

}

func (spr Sprite) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(spr.GetPosition().X), float64(spr.GetPosition().Y))
	op.GeoM.Scale(constants.SPRITE_SCALE, constants.SPRITE_SCALE)
	screen.DrawImage(spr.Image, op)
}

func (spr *Sprite) Move(x int, y int) {
	spr.SetDirection(x, y)
	spr.SetVelocity(spr.GetSpeed())
	spr.SetPosition(
		spr.GetPosition().X+x*spr.GetSpeed(),
		spr.GetPosition().Y+y*spr.GetSpeed(),
	)
}

func (spr Sprite) IsPressed(tID ebiten.TouchID) bool {
	sc := constants.SPRITE_SCALE
	c := spr.GetCollider()
	p := spr.GetPosition()
	if (Touches[tID].CurrX/sc) > p.X+c.Bounds.Min.X && (Touches[tID].CurrX/sc) < p.X+c.Bounds.Max.X {
		if (Touches[tID].CurrY/sc) > p.Y+c.Bounds.Min.Y && (Touches[tID].CurrY/sc) < p.Y+c.Bounds.Max.Y {
			return true
		}
	}
	return false
}

func NewSpriteRegistry() *SpriteRegistry {
	gr := &SpriteRegistry{
		Player:  NewPlayer(176, 176),
		GameMap: NewGameMap(0, 0),
		Props: []*Sprite{
			NewBox(50, 50).Sprite,
			NewBox(150, 150).Sprite,
			NewBox(150, 225).Sprite,
		},
		Input: NewControls(80, 472),
	}
	return gr
}

func (gr *SpriteRegistry) Update() {
	gr.GameMap.Update()
	for _, p := range gr.Props {
		p.Update()
	}
	gr.Player.Update()
	gr.Input.Update()
}

func (gr *SpriteRegistry) Draw(screen *ebiten.Image) {
	gr.GameMap.Draw(screen)
	for _, s := range gr.Props {
		s.Draw(screen)
	}
	gr.Player.Sprite.Draw(screen)
	gr.Input.Draw(screen)
}
