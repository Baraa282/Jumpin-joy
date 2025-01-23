package entities

import (
	// "fmt"

	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// x-width 28
// y- 88

type Box struct {
	X        float64
	Y        float64
	Width    float64
	Height   float64
	BoxImage *ebiten.Image
}

var (
	SliceOfBoxes []*Box
)

func init() {

	img, _, err := ebitenutil.NewImageFromFile("./assets/box.png")
	if err != nil {
		log.Fatal(err)
	}

	InitializeBoxes(img)
}

func InitializeBoxes(img *ebiten.Image) {
	boxes := []struct {
		xScale, yScale, x, y float64
	}{
		{0.3, 0.4, 100, 300},
		{0.35, 0.5, 400, 150},
		{0.25, 0.6, 250, 200},
		{4, 0.6, -100, 450},
		// {0.3, 0.4, 100, 150},
		// {0.3, 1, 100, 300},
		// {0.35, 1, 400, 150},
		// {0.5, 1, 150, 300},
	}

	for _, b := range boxes {
		box := NewBox(img, b.xScale, b.yScale, b.x, b.y)
		SliceOfBoxes = append(SliceOfBoxes, box)
	}
}

func NewBox(img *ebiten.Image, xScale, yScale, x, y float64) *Box {
	return &Box{
		X:        x,
		Y:        y,
		Width:    float64(img.Bounds().Dx()) * xScale,
		Height:   float64(img.Bounds().Dy()) * yScale,
		BoxImage: img,
	}
}

func (b *Box) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(b.Width/float64(b.BoxImage.Bounds().Dx()), b.Height/float64(b.BoxImage.Bounds().Dy()))
	op.GeoM.Translate(b.X, b.Y)
	screen.DrawImage(b.BoxImage, op)
}

func (b *Box) DrawBoxes(screen *ebiten.Image) {
	for _, box := range SliceOfBoxes {
		box.Draw(screen)
	}
}

func GetSliceOfBoxes() []*Box {
	return SliceOfBoxes
}

func (box *Box) Bounds() (left, top, right, bottom int) {
	return int(box.X), int(box.Y), int(box.X + box.Width), int(box.Y + box.Height)
}
