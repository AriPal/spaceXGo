package main

import (
	"bytes"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	bg            *ebiten.Image
	player        *ebiten.Image
	screenWith    = 480
	screenHeight  = 640
	audioContext  *audio.Context
	sfxLaunchData []byte
)

type Game struct {
	altitude     float64
	speed        float64
	power        float64
	launchPlayer *audio.Player
}

func init() {

	var err error

	// Load background
	bg, _, err = ebitenutil.NewImageFromFile("assets/background.png")
	if err != nil {
		log.Fatal(err)
	}

	// Load player sprite
	player, _, err = ebitenutil.NewImageFromFile("assets/player.png")
	if err != nil {
		log.Fatal(err)
	}

	// Add music
	sfxLaunchData, err = os.ReadFile("assets/sounds/launch.mp3")
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

		log.Printf("g.altitude: %d", g.altitude)
	}

	gravityPull := 0.1

	// Gravity pull down
	if !ebiten.IsKeyPressed(ebiten.KeySpace) {
		log.Printf("Space key is not pressed")
		gravityPull = gravityPull - 10
		// VelocityY -= gravity * deltatime
		g.altitude -= 60*deltatime - gravityPull
	}

	// Rocket is back on track
	if g.altitude < -5763 {
		g.altitude = -5763
		g.power = 0
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

func playSFX(sfxSound []byte) *audio.Player {
	// Decode the bytes
	stream, err := mp3.DecodeWithoutResampling(bytes.NewReader(sfxSound))
	if err != nil {
		log.Fatal(stream)
	}

	// Create player so that we can play, stop etc.
	musicPlayer, err := audioContext.NewPlayer(stream)
	if err != nil {
		log.Fatal(err)
	}

	musicPlayer.Play()

	return musicPlayer
}

func LoadMusic() {
	var sampleRate = 48000

	// Load music
	var err error

	// Create audo context
	audioContext = audio.NewContext(sampleRate)

	// Read file and return bytes
	data, err := os.ReadFile("assets/sounds/bossa_nova.mp3")
	if err != nil {
		log.Fatal(err)
	}

	// Decode the bytes
	stream, err := mp3.DecodeWithoutResampling(bytes.NewReader(data))
	if err != nil {
		log.Fatal(stream)
	}

	//
	musicPlayer, err := audioContext.NewPlayer(stream)
	if err != nil {
		log.Fatal(err)
	}

	musicPlayer.Play()
}

func main() {

	LoadMusic()

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
