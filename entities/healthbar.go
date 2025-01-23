package entities

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Healthbar struct {
}

var (
	Health1 *ebiten.Image
	Health2 *ebiten.Image
	Health3 *ebiten.Image
	Health4 *ebiten.Image
	Health5 *ebiten.Image
	Health6 *ebiten.Image
)

func init() {
	img, _, err := ebitenutil.NewImageFromFile("./assets/fullhealth.png")
	if err != nil {
		panic(err)
	}
	Health1 = img

	img, _, err = ebitenutil.NewImageFromFile("./assets/80percent.png")
	if err != nil {
		panic(err)
	}
	Health2 = img

	img, _, err = ebitenutil.NewImageFromFile("./assets/60percent.png")
	if err != nil {
		panic(err)
	}
	Health3 = img

	img, _, err = ebitenutil.NewImageFromFile("./assets/40percent.png")
	if err != nil {
		panic(err)
	}
	Health4 = img

	img, _, err = ebitenutil.NewImageFromFile("./assets/20percent.png")
	if err != nil {
		panic(err)
	}
	Health5 = img

	img, _, err = ebitenutil.NewImageFromFile("./assets/0percent.png")
	if err != nil {
		panic(err)
	}
	Health6 = img
}


func (h *Healthbar) DrawHealth(screen *ebiten.Image, Health float32) {
	op := &ebiten.DrawImageOptions{}

	health := fmt.Sprintf("PLAYER 1 HEALTH: %0.2f\n", Health)
	ebitenutil.DebugPrintAt(screen, health, 25, 25)
	if Health <= 100 && Health >= 80 {
		screen.DrawImage(Health1, op)
	}
	if Health <= 80 && Health >= 60 {
		screen.DrawImage(Health2, op)
	}
	if Health <= 60 && Health >= 40 {
		screen.DrawImage(Health3, op)
	}
	if Health <= 40 && Health >= 20 {
		screen.DrawImage(Health4, op)
	}
	if Health <= 20 && Health >= 0 {
		screen.DrawImage(Health5, op)
	}
	if Health == 0 {
		screen.DrawImage(Health6, op)
	}
}
