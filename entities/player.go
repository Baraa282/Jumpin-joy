package entities

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func CreatePlayer(pulse chan bool, coords chan [2]int) *Entity {
	ent := &Entity{x: 50 * unit, y: 160 * unit, Height: 102, Width: 28, vx: 0, vy: 0, leftKey: ebiten.KeyA,
		rightKey: ebiten.KeyD, jumpKey: ebiten.KeyW}

	// run seprate goroutine for player
	go playerThread(ent, pulse, coords)

	return ent
}

var leftside bool
var rightside bool


func PlatformTopBottom(player *Entity, collisionTopBottom string) {
	switch collisionTopBottom {
	case "Top":
		player.vy = 0 // Stop movement in y-led
		player.y -= 1 // Moves player to top of platform, works without but then the players feet are in the platform
	case "Bottom":
		player.vy = 0  // Stop movement in y-led
		player.y += 60 // Dont know why, but it works.
	}
}

func PlatformSides(player *Entity, collisionSides string) {
	leftside = false
	rightside = false
	switch collisionSides {
	case "Left": // mixed up in my head need to change sides but it works
		player.vx = 0 // Stop movement in x-led
		leftside = true
	case "Right": // mixed up in my head need to change sides but it works
		player.vx = 0 // Stop movement in x-led
		rightside = true
	}
}

func playerThread(c *Entity, pulse chan bool, coords chan [2]int) {
	// for range pulse will block itself until pulse recieves message.
	// pulse is sent from inside Update() - this will therefore sync player with Update().
	for signal := range pulse {
		switch signal {
		case true:

			c.PlayerKeyPress()
			coords <- c.PlayerUpdate()

		case false:
			return
		}
	}
}

func (c *Entity) PlayerKeyPress() {
	// Controls
	if ebiten.IsKeyPressed(c.leftKey) && !rightside {
		if c.x > 0 {
			c.vx = -4 * unit
		}
	} else if ebiten.IsKeyPressed(c.rightKey) && !leftside {
		if leftside {
			c.vx = 0

		} else if c.x < 640*unit {
			c.vx = 4 * unit
		}
	}
	if inpututil.IsKeyJustPressed(c.jumpKey) {
		c.tryJump()
	}
}

func (c *Entity) tryJump() {
	if c.vy == 0 {
		c.vy = -10 * (unit * 1.5)
	}
}

func (c *Entity) PlayerUpdate() [2]int {
	// Change position based on current velocity
	c.x += c.vx
	c.y += c.vy

	// Reduce x-velocity every tick
	if c.vx > 0 {
		c.vx -= 4
	} else if c.vx < 0 {
		c.vx += 4
	}
	// Gravity
	if c.y < groundY*unit && c.vy < unit*unit {
		c.vy += 8
	}

	return [2]int{c.x, c.y}
}
