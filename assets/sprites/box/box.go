package box

import "github.com/seanburman/game/assets"

type Box struct {
	*assets.Sprite
}

func NewBox(x int, y int) *Box {
	box := &Box{
		assets.NewSprite("box", assets.NewSpriteOpts{
			ImgSource: "box/img/box.png",
		}),
	}
	box.SetPosition(x, y)
	return box
}

func (b *Box) Update() {

}
