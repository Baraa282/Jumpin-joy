package logic

import (
	"github.com/sammifs/jumpin-joe/entities"
)

// CheckCollision checks for collisions between the player and any box and returns the collision type
func CheckTopBottom(player *entities.Entity, boxes []*entities.Box) string {
	playerLeft, playerTop, playerRight, playerBottom := player.Bounds()
// println("1")
	for _, box := range boxes {
		boxLeft, boxTop, boxRight, boxBottom := box.Bounds()
// println("2")
		if playerRight > boxLeft && playerLeft < boxRight && playerBottom > boxTop && playerTop < boxBottom {
			// Determine the type of collision
			collisionWidth := min(playerRight, boxRight) - max(playerLeft, boxLeft)
			collisionHeight := min(playerBottom, boxBottom) - max(playerTop, boxTop)

			if collisionHeight < collisionWidth {
				if playerBottom > boxTop && playerTop < boxTop {
					return "Top" // Player is on top of the box
				} else if playerTop < boxBottom && playerBottom > boxBottom {
					return "Bottom" // Player hits the bottom of the box
				}
			}
		}
	}
	return "None" // No collision
}

func CheckSides(player *entities.Entity, boxes []*entities.Box) string {
	playerLeft, playerTop, playerRight, playerBottom := player.Bounds()
// println("1")
	for _, box := range boxes {
		boxLeft, boxTop, boxRight, boxBottom := box.Bounds()
// println("2")
		if playerRight > boxLeft && playerLeft < boxRight && playerBottom > boxTop && playerTop < boxBottom {
			// Collision?
			collisionWidth := min(playerRight, boxRight) - max(playerLeft, boxLeft)
			collisionHeight := min(playerBottom, boxBottom) - max(playerTop, boxTop)

			if collisionHeight > collisionWidth {
				if playerRight > boxLeft && playerLeft < boxLeft {
					return "Left" // Player hits the left side of the box (maybe wrong way idk)
				} else if playerLeft < boxRight && playerRight > boxRight {
					return "Right" // Player hits the right side of the box (maybe wrong way idk)
				}
			}
		}
	}
	return "None" // No collision
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}