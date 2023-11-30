package sprites

type XY struct {
	X, Y int
}

type position XY
type direction XY
type velocity int

type Vector struct {
	p position
	s int
	d direction
	v velocity
}

func (vc *Vector) Update() {
	vc.v = 0
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

func (vc Vector) GetVelocity() int {
	return int(vc.v)
}

func (vc *Vector) SetVelocity(v int) {
	vc.v = velocity(v)
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
