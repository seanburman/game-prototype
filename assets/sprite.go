package assets

import (
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/seanburman/game/constants"
)

type SpriteInterface interface {
	Update()
	Draw(screen *ebiten.Image)
	GetName() Name
	SetName(n string)
	IsPressed(tID ebiten.TouchID) bool
	Move(x int, y int)
	GetDirection() direction
	GetPosition() position
	GetSpeed() int
	GetVector() Vector
	SetDirection(x int, y int)
	SetPosition(x int, y int)
	SetSpeed(s int)
	DetectBoundary(v Vector, foreign Collider)
	DetectCollision(v Vector, f SpriteInterface) *Collider
	GetBounds() Bounds
	GetBreached() Breached
	GetCollider() Collider
}

type Name string

func (a Name) GetName() Name {
	return a
}

func (a *Name) SetName(n string) {
	*a = Name(n)
}

type Sprite struct {
	Name
	*ebiten.Image
	Vector
	Collider
}

type NewSpriteOpts struct {
	Name      string
	ImgSource string
	Vector    Vector
}

func NewSprite(name string, opts NewSpriteOpts) *Sprite {
	img, _, err := ebitenutil.NewImageFromFile("assets/sprites/" + opts.ImgSource)
	if err != nil {
		log.Fatal(err)
	}

	c := NewCollider(Bounds{
		Min: image.Point{X: img.Bounds().Min.X, Y: img.Bounds().Min.Y},
		Max: image.Point{X: img.Bounds().Max.X, Y: img.Bounds().Max.Y},
	})

	return &Sprite{
		Name(name),
		img,
		opts.Vector,
		c,
	}
}

func (spr Sprite) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(spr.GetPosition().X), float64(spr.GetPosition().Y))
	op.GeoM.Scale(constants.SPRITE_SCALE, constants.SPRITE_SCALE)
	screen.DrawImage(spr.Image, op)
}

func (spr *Sprite) Move(x int, y int) {
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
