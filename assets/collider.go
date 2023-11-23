package assets

import (
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

func (c *Collider) GetBounds() Bounds {
	return c.Bounds
}

func (c Collider) GetBreached() Breached {
	return c.Breached
}

func (c Collider) GetCollider() Collider {
	return c
}

func (c *Collider) DetectCollision(v Vector, foreign SpriteInterface) *Collider {
	// Foreign object
	fc := foreign.GetCollider()
	fp := foreign.GetVector().GetPosition()
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

func (c *Collider) DetectBoundary(v Vector, foreign Collider) {
	// Approaching from right
	if v.GetPosition().X-v.GetSpeed() <= foreign.Bounds.Min.X {
		c.Breached.Min.X = true
	}
	// Approaching from left
	if v.GetPosition().X+c.GetBounds().Max.X+v.GetSpeed() >= foreign.Bounds.Max.X {
		c.Breached.Max.X = true
	}
	// Approaching from bottom
	if v.GetPosition().Y-v.GetSpeed() <= foreign.Bounds.Min.Y {
		c.Breached.Min.Y = true
	}
	// Approaching from top
	if v.GetPosition().Y+c.GetBounds().Max.Y+v.GetSpeed() >= foreign.Bounds.Max.Y {
		c.Breached.Max.Y = true
	}
}
