package logic

func AABBIntersection(playerCoords [2]int, px float64, py float64, npcCoords [2]int, npcx float64, npcy float64) bool {
	// Get Entity coordinates, everithing has to be devided by 16?, called unit in main
	// px and npcx are both entities width
	// py and npcy are both entities height
	playerLeftX := playerCoords[0] / 16
	playerRightX := float64(playerCoords[0]/16) + px
	playerTopY := playerCoords[1] / 16
	playerLowY := float64(playerCoords[1]/16) + py

	npcLeftX := npcCoords[0] / 16
	npcRightX := float64(npcCoords[0]/16) + npcx
	npcTopY := npcCoords[1] / 16
	npcLowY := float64(npcCoords[1]/16) + npcy

	// If player or npc are to the right/left of eacother
	// or if player or npc are above/below of eachother
	// return false
	// else return true, cus they are intersecting
	if float64(playerLeftX) > npcRightX || float64(npcLeftX) > playerRightX {
		return false
	} else if playerLowY < float64(npcTopY) || npcLowY < float64(playerTopY) {
		return false
	} else {
		return true
	}

}
