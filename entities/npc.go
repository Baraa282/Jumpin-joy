package entities

func CreateNpc(pulse chan bool, coords chan [2]int) *Entity {
	ent := &Entity{x: 500 * unit, y: groundY * unit, vx: -30, vy: 0}

	// go routine for npc
	go npcThread(ent, pulse, coords)

	return ent
}

func CreateBird(pulse chan bool, coords chan [2]int) *Entity {
	ent := &Entity{x: 200 * unit, y: 20 * unit, vx: -30, vy: 0}

	go birdThread(ent, pulse, coords)

	return ent
}

func birdThread(c *Entity, pulse chan bool, coords chan [2]int) {
	for signal := range pulse {
		switch signal {
		case true:
			coords <- c.BirdUpdate()
		case false:
			return
		}
	}
}

func (c *Entity) BirdUpdate() [2]int {
	c.x += c.vx
	if c.x < 0 || c.x > 640*14 {
		c.vx = c.vx * -1
	}
	return [2]int{c.x, c.y}
}

func npcThread(c *Entity, pulse chan bool, coords chan [2]int) {
	for signal := range pulse {
		switch signal {
		case true:
			coords <- c.NpcUpdate()
		case false:
			return
		}
	}
}
func (c *Entity) NpcUpdate() [2]int {
	c.x += c.vx
	if c.x < 0 || c.x > 640*14 {
		c.vx = c.vx * -1
	}
	return [2]int{c.x, c.y}
}

func CreateEgg(pulse chan bool, coords chan [2]int, coordX int) *Entity {
	ent := &Entity{x: coordX, y: 20 * unit, vx: -30, vy: 0}

	go eggThread(ent, pulse, coords)

	return ent
}

func eggThread(c *Entity, pulse chan bool, coords chan [2]int) {
	for signal := range pulse {
		switch signal {
		case true:
			coords <- c.eggUpdate()
		case false:
			return
		}
	}
}

func (c *Entity) eggUpdate() [2]int {
	c.x += c.vx
	c.y += c.vy

	c.vy = 5 * (unit * 1.5)

	if c.x < 0 || c.x > 640*14 {
		c.vx = c.vx * -1
		// c.vy = 10 * (unit * 1.5)
		// c.vy = 5 * (unit * 1.5)
		//c.vy += 8
	}

	// if c.y == groundY*unit {
	// 	//c.y = 0
	// }

	// if c.y == (groundY*unit) {
	// 	c.y = (groundY * unit) - 600
	// 	//c.vy = -30
	// 	c.vy = 0
	// }

	// if c.y == (groundY*unit) {
	// 	c.y = 0

	// }

	return [2]int{c.x, c.y}
}
