package assets

type XY struct {
	X, Y int
}

type position XY
type direction XY

type Vector struct {
	p position
	s int
	d direction
}

func (vc Vector) GetVector() Vector {
	return vc
}

func (vc Vector) GetSpeed() int {
	return vc.s
}
func (vc *Vector) SetSpeed(s int) {
	vc.s = s
}

func (vc Vector) GetPosition() position {
	return vc.p
}
func (vc *Vector) SetPosition(x int, y int) {
	vc.p.X = x
	vc.p.Y = y
}

func (vc Vector) GetDirection() direction {
	return vc.d
}
func (vc *Vector) SetDirection(x int, y int) {
	vc.d.X = x
	vc.d.Y = y
}
