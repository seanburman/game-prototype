package assets

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

// In order of precedence
var Registries = NewRegistries()
var PropRegistry = NewRegistry()
var CharacterRegistry Registry = NewRegistry()
var TileRegistry Registry = NewRegistry()
var InputRegistry Registry = NewRegistry()

type registries []Registry

func NewRegistries() *registries {
	r := make(registries, 4)
	return &r
}

func (rs *registries) Update() {
	for _, r := range *rs {
		r.Update()
	}
}

func (rs *registries) Draw(screen *ebiten.Image) {
	for _, r := range *rs {
		r.Draw(screen)
	}
}

type Registry struct {
	Sprites []SpriteInterface
}

func NewRegistry() Registry {
	return Registry{
		Sprites: []SpriteInterface{},
	}
}

func (r *Registry) Register(si []SpriteInterface) {
	r.Sprites = append(r.Sprites, si...)
	fmt.Println("ok")
}

func (r *Registry) Update() {
	for _, s := range r.Sprites {
		if s != nil {
			s.Update()
		}
	}
}

func (r *Registry) Draw(screen *ebiten.Image) {
	for _, s := range r.Sprites {
		s.Draw(screen)
	}
}

// TODO: Create objects at start up which can be cloned during run time
type ObjectLibrary struct {
}
