package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/seanburman/game/assets"
	"github.com/seanburman/game/assets/fonts"
	"github.com/seanburman/game/assets/sprites/box"
	"github.com/seanburman/game/assets/sprites/square"
	"github.com/seanburman/game/client"
	"github.com/seanburman/game/config"
	"github.com/seanburman/game/constants"
)

var MessageClient *client.Client
var Input *fonts.TextInput

func init() {
	MessageClient = client.NewClient("/messages")
	go MessageClient.Dial("/messages")
	go MessageClient.OnPage()

	Input = fonts.NewTextInput(fonts.Roboto)
	assets.PropRegistry.Register([]assets.SpriteInterface{
		box.NewBox(50, 50),
		box.NewBox(150, 150),
		box.NewBox(150, 225),
	})
	assets.CharacterRegistry.Register([]assets.SpriteInterface{
		square.NewSquare(10, 10),
	})
	assets.InputRegistry.Register([]assets.SpriteInterface{
		assets.NewControls(80, 472),
	})
}

type Game struct {
	bounds assets.Collider
}

func (g *Game) Update() error {
	assets.PropRegistry.Update()
	assets.CharacterRegistry.Update()
	assets.InputRegistry.Update()
	Input.Field.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{255, 255, 255, 255})
	assets.PropRegistry.Draw(screen)
	assets.CharacterRegistry.Draw(screen)
	assets.InputRegistry.Draw(screen)
	Input.Field.Draw(screen)
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
	if err := ebiten.RunGame(&Game{bounds: config.GAME_BOUNDARIES}); err != nil {
		log.Fatal(err)
	}
}
