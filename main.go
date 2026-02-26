package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	bg           *ebiten.Image
	player       *ebiten.Image
	screenWith   = 480
	screenHeight = 640
)

type Game struct {
	altitude float64
	speed    float64
	power    float64
}

func init() {

	var err error

	bg, _, err = ebitenutil.NewImageFromFile("assets/background.png")
	if err != nil {
		log.Fatal(err)
	}

	player, _, err = ebitenutil.NewImageFromFile("assets/player.png")
	if err != nil {
		log.Fatal(err)
	}
}

func (g *Game) Update() error {

	const deltatime = 1.0 / 60.0

	// Lauch rocket upwards
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		g.power += 0.1
		g.altitude += g.power
		log.Printf("g.altitude: %t", g.altitude)
	}

	// Gravity pull down
	if !ebiten.IsKeyPressed(ebiten.KeySpace) {
		log.Printf("Space key is not pressed")
		// VelocityY -= gravity * deltatime
		g.altitude -= 60 * deltatime
	}

	// Rocket is back on track
	if g.altitude < -5763 {
		g.altitude = -5763
	}

	return nil

}

func (g *Game) Draw(screen *ebiten.Image) {

	// Draw background
	bgOp := &ebiten.DrawImageOptions{}
	bgOp.GeoM.Translate(0, g.altitude)
	screen.DrawImage(bg, bgOp)

	// Draw player
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(195, 300)
	screen.DrawImage(player, op)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (w, h int) {
	return screenWith, screenHeight
}

func main() {
	ebiten.SetWindowTitle("Space X")
	ebiten.SetWindowSize(screenWith, screenHeight)

	game := &Game{
		altitude: -5763,
		speed:    0,
		power:    0,
	}

	err := ebiten.RunGame(game)

	if err != nil {
		log.Fatal(err)
	}

}
