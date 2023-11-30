package sprites

import (
	"fmt"
	"image"
)

type Breached struct {
	Min struct {
		X, Y bool
	}
	Max struct {
		X, Y bool
	}
}

func NewBreached() Breached {
	return Breached{}
}

type Bounds struct {
	Min image.Point
	Max image.Point
}

type Collider struct {
	Bounds
	Breached
}

func NewCollider(b Bounds) Collider {
	return Collider{Bounds: b, Breached: NewBreached()}
}

func (c *Collider) Update() {
	*c = NewCollider(c.Bounds)
}

func (c Collider) GetCollider() Collider {
	return c
}

func (c *Collider) GetBounds() Bounds {
	return c.Bounds
}

func (c *Collider) SetBounds(b Bounds) {
	c.Bounds = b
}

func (c Collider) GetBreached() Breached {
	return c.Breached
}

func (c *Collider) SetBreached(b Breached) {
	c.Breached = b
}

func (c *Collider) DetectCollision(v Vector, foreign *Sprite) *Collider {
	// Foreign object
	fc := foreign.Collider
	fp := foreign.Vector.GetPosition()
	fLeft := fp.X
	fRight := fp.X + fc.Bounds.Max.X
	fTop := fp.Y
	fBottom := fp.Y + fc.Bounds.Max.Y

	// Local object
	lp := v.GetPosition()
	lRight := lp.X + c.Bounds.Max.X
	lLeft := lp.X
	lBottom := lp.Y + c.Bounds.Max.Y
	lTop := lp.Y

	// Approaching from left
	// Local object is relatively parallel with Foreign object on X axis and would collide with right movement
	if ((lBottom >= fTop) && (lTop <= fBottom)) && (lRight < fLeft) && (lRight+v.GetSpeed() >= fLeft) {
		c.Breached.Max.X = true
	}
	// Approaching from right
	// Local object is relatively parallel with Foreign object on X axis and would collide with left movement
	if ((lTop <= fBottom) && (lBottom >= fTop)) && (lLeft > fRight) && (lLeft-v.GetSpeed() <= fRight) {
		c.Breached.Min.X = true
	}
	// Approaching from top
	// Local object is relatively parallel with Foreign object on Y axis and would collide with downward movement
	if ((lLeft <= fRight) && (lRight >= fLeft)) && (lBottom < fTop) && (lBottom+v.GetSpeed() >= fTop) {
		c.Breached.Max.Y = true
	}
	// Approaching from bottom
	// Local object is relatively parallel with Foreign object on Y axis and would collide with upward movement
	if ((lLeft <= fRight) && (lRight >= fLeft)) && (lTop > fBottom) && (lTop-v.GetSpeed() <= fBottom) {
		c.Breached.Min.Y = true
	}

	return c
}

func (c *Collider) DetectBoundary(v Vector, foreign *Collider) {
	b := Breached{}
	// Approaching from right
	if v.GetPosition().X-v.GetSpeed() <= foreign.Bounds.Min.X {
		c.Breached.Min.X = true
		b.Min.X = true
	}
	// Approaching from left
	if v.GetPosition().X+c.GetBounds().Max.X+v.GetSpeed() >= foreign.Bounds.Max.X {
		c.Breached.Max.X = true
		b.Max.X = true
	}
	// Approaching from bottom
	if v.GetPosition().Y-v.GetSpeed() <= foreign.Bounds.Min.Y {
		c.Breached.Min.Y = true
		b.Min.Y = true
	}
	// Approaching from top
	if v.GetPosition().Y+c.GetBounds().Max.Y+v.GetSpeed() >= foreign.Bounds.Max.Y {
		c.Breached.Max.Y = true
		b.Max.Y = true
	}

	fmt.Println(foreign.Breached)
	fmt.Println(b)
	foreign.SetBreached(b)
}

func (c *Collider) DetectMovingBoundary(v Vector, foreign *Sprite) {
	newBreached := Breached{}

	fbo := foreign.GetBounds()
	fv := foreign.GetVector()
	// Approaching from right
	if v.GetPosition().X-v.GetSpeed() <= fv.p.X {
		c.Breached.Min.X = true
		newBreached.Min.X = true
	}
	// Approaching from left
	if v.GetPosition().X+c.GetBounds().Max.X+v.GetSpeed() >= fv.p.X+fbo.Max.X {
		c.Breached.Max.X = true
		newBreached.Max.X = true
	}
	// Approaching from bottom
	if v.GetPosition().Y-v.GetSpeed() <= fv.p.Y {
		c.Breached.Min.Y = true
		newBreached.Min.Y = true
	}
	// Approaching from top
	if v.GetPosition().Y+c.GetBounds().Max.Y+v.GetSpeed() >= fv.p.Y+fbo.Max.Y {
		c.Breached.Max.Y = true
		newBreached.Max.Y = true
	}
	foreign.SetBreached(newBreached)
	fmt.Println(foreign.GetBreached())
}
