package entities

import "github.com/hajimehoshi/ebiten/v2"

type Entity struct {
	x        int
	y        int
	vx       int
	vy       int
	Height   int
	Width    int
	leftKey  ebiten.Key
	rightKey ebiten.Key
	jumpKey  ebiten.Key
}

const (
	unit    = 16
	groundY = 380
)

func (c *Entity) EntityCoords() (int, int) {
	return c.x, c.y
}

// gets the bounds of the entity like a hitbox
func (c *Entity) Bounds() (left, top, right, bottom int) {
	return c.x / 16, c.y / 16, c.x/16 + c.Width, c.y/16 + c.Height
}
