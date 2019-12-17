package main

import "github.com/hajimehoshi/ebiten"

type Game struct {
	currentScene        Scene
	previousScene       Scene
	mouseAlreadyPressed bool
}

type Scene interface {
	update(*ebiten.Image) error
}
