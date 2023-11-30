package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/seanburman/game/config"
	"github.com/seanburman/game/constants"
	"github.com/seanburman/game/sprites"
)

// var MessageClient *client.Client

// var Connected = -1

func init() {
	//TODO: UPDATE MAPGROUNDS BOUNDS, BECAUSE THEY ARE ALWAYS CHANGING
	// MessageClient = client.NewClient()
	// // go MessageClient.Dial("/messages")
	// go func() {
	// 	c := MessageClient.Handshake()
	// 	if c {
	// 		Connected = 1
	// 	} else {
	// 		Connected = 0
	// 	}
	// }()
}

type Game struct {
	bounds sprites.Collider
}

func (g *Game) Update() error {
	sprites.Registry.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 0, 0})
	// ebitenutil.DebugPrint(screen, fmt.Sprint("Connection status:  "+fmt.Sprint(Connected)))
	sprites.Registry.Draw(screen)
	if config.Message != "" {
		ebitenutil.DebugPrintAt(screen, string(config.Message), 300, 150)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return constants.GetDimensions(), constants.GetDimensions() + 500
}

func main() {
	ebiten.SetWindowSize(constants.GetDimensions(), constants.GetDimensions()+500)
	ebiten.SetWindowTitle("Game")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
