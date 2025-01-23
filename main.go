package main

import (
	"bytes"
	"fmt"
	_ "image/png"
	"io/ioutil"

	// "golang.org/x/image/font"
	// "golang.org/x/image/font/basicfont"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/sammifs/jumpin-joe/entities"
	"github.com/sammifs/jumpin-joe/logic"
)

const sampleRate = 10000 //10000   //44100 //This is a common sample rate for audio files
var jumpSound *audio.Player
var audioContext *audio.Context
var bgm *audio.Player
var hurtSound *audio.Player

const (
	// Settings
	screenWidth  = 640
	screenHeight = 480
)

var (
	playerSprite     *ebiten.Image
	playerSpriteL    *ebiten.Image
	playerSpriteIdle *ebiten.Image
	npcSprite        *ebiten.Image
	birdSprite       *ebiten.Image
	birdSpriteL      *ebiten.Image
	eggSprite        *ebiten.Image
	backgroundImage  *ebiten.Image
	gameoverImage    *ebiten.Image
	damage           bool
)

func init() {
	audioContext := audio.NewContext(sampleRate)
	// Load the jump sound
	jumpData, err := ioutil.ReadFile("./assets/jump .wav")
	if err != nil {
		panic(err)
	}

	jumpD, err := wav.Decode(audioContext, bytes.NewReader(jumpData))
	if err != nil {
		panic(err)
	}

	jumpSound, err = audio.NewPlayer(audioContext, jumpD)
	if err != nil {
		panic(err)
	}

	// Initialize and preload all assets
	img, _, err := ebitenutil.NewImageFromFile("./assets/background.png")
	if err != nil {
		panic(err)
	}
	backgroundImage = img

	img, _, err = ebitenutil.NewImageFromFile("./assets/player1.png")
	if err != nil {
		panic(err)
	}
	playerSprite = img

	img, _, err = ebitenutil.NewImageFromFile("./assets/player1_flipped.png")
	if err != nil {
		panic(err)
	}
	playerSpriteL = img

	img, _, err = ebitenutil.NewImageFromFile("./assets/player1_front.png")
	if err != nil {
		panic(err)
	}
	playerSpriteIdle = img

	img, _, err = ebitenutil.NewImageFromFile("./assets/enemy.png")
	if err != nil {
		panic(err)
	}
	npcSprite = img

	img, _, err = ebitenutil.NewImageFromFile("./assets/bird.png")
	if err != nil {
		panic(err)
	}
	birdSprite = img

	img, _, err = ebitenutil.NewImageFromFile("./assets/birdL.png")
	if err != nil {
		panic(err)
	}
	birdSpriteL = img

	img, _, err = ebitenutil.NewImageFromFile("./assets/egg.png")
	if err != nil {
		panic(err)
	}
	eggSprite = img

	img, _, err = ebitenutil.NewImageFromFile("./assets/game_over_cut.png")
	if err != nil {
		panic(err)
	}
	gameoverImage = img
	// Load the background music
	bgmData, err := ioutil.ReadFile("./assets/back.wav")
	if err != nil {
		panic(err)
	}

	bgmD, err := wav.Decode(audioContext, bytes.NewReader(bgmData))
	if err != nil {
		panic(err)
	}

	bgm, err = audio.NewPlayer(audioContext, bgmD)
	if err != nil {
		panic(err)
	}
	bgm.SetVolume(0.5) // Adjust volume if needed
	// Load the hurt sound
	hurtData, err := ioutil.ReadFile("./assets/hurt.wav")
	if err != nil {
		panic(err)
	}

	hurtD, err := wav.Decode(audioContext, bytes.NewReader(hurtData))
	if err != nil {
		panic(err)
	}

	hurtSound, err = audio.NewPlayer(audioContext, hurtD)
	if err != nil {
		panic(err)
	}
}

type Game struct {
	npcExists    bool
	npcPulse     chan bool
	npcCoordChan chan [2]int
	npcCoords    [2]int

	birdExists    bool
	birdPulse     chan bool
	birdCoordChan chan [2]int
	birdCoords    [2]int
	birdImages    *ebiten.Image

	eggExists    bool
	eggPulse     chan bool
	eggCoordChan chan [2]int
	eggCoords    [2]int

	health    float32
	healthbar *entities.Healthbar
	timer     float32
	gameOver  bool
	boxes     *entities.Box

	player          *entities.Entity
	playerExists    bool
	playerPulse     chan bool
	playerCoordChan chan [2]int
	playerCoords    [2]int
}

func (g *Game) timerUpdate() bool {
	if g.timer > 0 {
		g.timer -= 0.015
		return false
	} else {
		return true
	}
}

func (g *Game) PlayerInitialization() {
	if !g.playerExists {
		// Initialize playerPulse and player coordinates channel
		g.playerPulse = make(chan bool)
		g.playerCoordChan = make(chan [2]int, 1)

		// CreatePlayer initializes goroutine for player
		g.player = entities.CreatePlayer(g.playerPulse, g.playerCoordChan)
		g.playerExists = true
	}
}

func (g *Game) NpcInitialization() {
	if !g.npcExists {
		g.npcPulse = make(chan bool)
		g.npcCoordChan = make(chan [2]int, 1)

		entities.CreateNpc(g.npcPulse, g.npcCoordChan)
		g.npcExists = true
	}
}

func (g *Game) Update() error {
	//Handle Player input
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		jumpSound.Rewind()
		jumpSound.Play()
	}

	// Create player and npc if they dont exist
	g.PlayerInitialization()
	g.NpcInitialization()

	// Checks for collision
	collisionTopBottom := logic.CheckTopBottom(g.player, entities.SliceOfBoxes)
	collisionSides := logic.CheckSides(g.player, entities.SliceOfBoxes)

	// Change players movement
	entities.PlatformTopBottom(g.player, collisionTopBottom)
	entities.PlatformSides(g.player, collisionSides)

	if !g.birdExists {
		g.birdPulse = make(chan bool)
		g.birdCoordChan = make(chan [2]int, 1)

		entities.CreateBird(g.birdPulse, g.birdCoordChan)
		g.birdExists = true
	}

	if !g.eggExists {
		g.eggPulse = make(chan bool)
		g.eggCoordChan = make(chan [2]int, 1)

		entities.CreateEgg(g.eggPulse, g.eggCoordChan, 100*unit)
		g.eggExists = true
	}

	// Pulse player
	g.playerPulse <- true
	// Get return from player
	g.playerCoords = <-g.playerCoordChan

	// Pulse npc
	g.npcPulse <- true
	// Get return from npc
	g.npcCoords = <-g.npcCoordChan
	// Pulse bird
	g.birdPulse <- true
	// Get return from bird
	g.birdCoords = <-g.birdCoordChan
	// Pulse egg
	g.eggPulse <- true

	// Get return from egg
	if int(g.timer)%3 == 0 {

		g.eggCoords = <-g.eggCoordChan
		g.eggCoords = g.birdCoords
		g.eggPulse <- false
		entities.CreateEgg(g.eggPulse, g.eggCoordChan, g.birdCoords[0])

	} else {
		g.eggCoords = <-g.eggCoordChan
	}

	if g.timerUpdate() {
		if g.timer < 0 {
			g.gameOver = true
		}
	}

	// Game ends on player touching enemy or timer running out
	//if logic.AABBIntersection(g.playerCoords, 28, 88, g.npcCoords, 64, 22) || g.timerUpdate() { // hårdkodad
	if logic.AABBIntersection(g.playerCoords, float64(playerSprite.Bounds().Dx())*0.199, float64(playerSprite.Bounds().Dy())*0.199,
		g.npcCoords, float64(npcSprite.Bounds().Dx())*0.057, float64(npcSprite.Bounds().Dy())*0.057) ||
		logic.AABBIntersection(g.playerCoords, float64(playerSprite.Bounds().Dx())*0.199, float64(playerSprite.Bounds().Dy())*0.199,
			g.eggCoords, float64(eggSprite.Bounds().Dx())*0.125, float64(eggSprite.Bounds().Dy())*0.125) || g.timerUpdate() { // bildstorlek, forfarande skev scaling dock, så vi måste div med nära nog konstanter (int)

		if g.health == 0 || g.timer < 0 {
			g.gameOver = true
		}
		// Checks if hp > 0, subtracts 20 from player and then sets damage flag to false
		if (g.health > 0) && damage {
			g.health -= 20
			damage = false
			// Play the hurt sound when the player takes damage
			hurtSound.Rewind()
			hurtSound.Play()

		}
	} else {
		damage = true
	}

	return nil
}

const (
	unit = 16
)

func (g *Game) Drawbird(screen *ebiten.Image, birdImage *ebiten.Image) {
	bo := &ebiten.DrawImageOptions{}
	bo.GeoM.Scale(0.125, 0.125)
	// bo.GeoM.Translate(float64(g.birdCoords[0])/unit, float64(g.birdCoords[1])/unit)
	bo.GeoM.Translate(float64(g.birdCoords[0])/unit, 0)
	screen.DrawImage(birdImage, bo)

}

func (g *Game) Draw(screen *ebiten.Image) {
	// Draws Background Image.
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(1, 1)
	screen.DrawImage(backgroundImage, op)

	// Draws the player.
	po := &ebiten.DrawImageOptions{}
	po.GeoM.Scale(0.199, 0.199)
	po.GeoM.Translate(float64(g.playerCoords[0])/unit, float64(g.playerCoords[1])/unit)
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		screen.DrawImage(playerSpriteL, po)
	} else if ebiten.IsKeyPressed(ebiten.KeyD) {
		screen.DrawImage(playerSprite, po)
	} else {
		screen.DrawImage(playerSpriteIdle, po)
	}

	// Draws the npc
	no := &ebiten.DrawImageOptions{}
	no.GeoM.Scale(0.057, 0.057)
	no.GeoM.Translate(float64(g.npcCoords[0])/unit, float64(g.npcCoords[1])/unit)
	screen.DrawImage(npcSprite, no)

	//Bird
	birdx := g.birdCoords[0] / 16
	switch birdx {
	// if g.timer < 25 && g.timer > 20 || g.timer < 15 && g.timer > 10 {
	//if int(g.timer) % 6 == 0 {
	case 0:

		g.birdImages = birdSpriteL

	case 561:
		g.birdImages = birdSprite
	}
	g.Drawbird(screen, g.birdImages)

	//Draw the egg
	ko := &ebiten.DrawImageOptions{}
	ko.GeoM.Scale(0.125, 0.125)
	ko.GeoM.Translate(float64(g.eggCoords[0])/unit, (float64(g.eggCoords[1]) / unit))
	//ko.GeoM.Translate(float64(g.eggCoords[0])/unit, 0)
	screen.DrawImage(eggSprite, ko)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("%f", float64(g.playerCoords[0])/unit))
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%d", g.playerCoords[0]), 0, 10)

	g.healthbar.DrawHealth(screen, g.health)

	g.boxes.DrawBoxes(screen)

	if g.gameOver {
		op.GeoM.Translate(190, 180)
		screen.DrawImage(gameoverImage, op)
		ebitenutil.DebugPrintAt(screen, "Press SPACE to restart", 265, 160)
		// ebitenutil.DebugPrintAt(screen, "|", g.playerCoords[0]/unit, g.playerCoords[1]/unit)
		// ebitenutil.DebugPrintAt(screen, "|", g.playerCoords[0]/unit+playerSprite.Bounds().Dx()/7, g.playerCoords[1]/unit)
		// ebitenutil.DebugPrintAt(screen, "|", g.playerCoords[0]/unit, g.playerCoords[1]/unit+playerSprite.Bounds().Dy()/6)
		// ebitenutil.DebugPrintAt(screen, "|", g.playerCoords[0]/unit+28, g.playerCoords[1]/unit+87)
		// ebitenutil.DebugPrintAt(screen, "|", g.npcCoords[0]/unit, g.npcCoords[1]/unit)
		// ebitenutil.DebugPrintAt(screen, "|", g.npcCoords[0]/unit+npcSprite.Bounds().Dx()/20, g.npcCoords[1]/unit)
		// ebitenutil.DebugPrintAt(screen, "|", g.npcCoords[0]/unit, g.npcCoords[1]/unit+playerSprite.Bounds().Dy()/21)
		if ebiten.IsKeyPressed(ebiten.KeySpace) {
			restartGame(g)
		}
	}
	if !g.gameOver {
		msg := fmt.Sprintf("TIME LEFT: %0.2f\n", g.timer)
		ebitenutil.DebugPrintAt(screen, msg, 265, 160)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func restartGame(g *Game) {
	g.timer = 30.0
	g.gameOver = false
	g.health = 100
	g.npcPulse <- false
	entities.CreateNpc(g.npcPulse, g.npcCoordChan)
	g.playerPulse <- false
	g.player = entities.CreatePlayer(g.playerPulse, g.playerCoordChan)
	g.birdPulse <- false
	entities.CreateBird(g.birdPulse, g.birdCoordChan)
	g.eggPulse <- false
	entities.CreateEgg(g.eggPulse, g.eggCoordChan, 100*unit)
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Jumpin' Joe")
	// Play background music
	bgm.Rewind()
	bgm.Play()

	// Create game struct with set end-time
	game := &Game{timer: 30.0, gameOver: false, health: 100, birdImages: birdSprite}
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
