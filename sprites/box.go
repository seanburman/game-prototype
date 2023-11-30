package sprites

type Box struct {
	*Sprite
}

func NewBox(x int, y int) *Box {
	src := "props/box/box.png"
	box := &Box{
		NewSprite("box", NewSpriteOpts{
			ImgSource: &src,
		}),
	}
	box.SetPosition(x, y)
	return box
}

func (b *Box) Update() {

}
