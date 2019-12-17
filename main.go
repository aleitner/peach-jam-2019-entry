package main

import (
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/audio/mp3"
	"github.com/hajimehoshi/ebiten/audio/wav"
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"golang.org/x/image/font"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	screenWidth  = 960
	screenHeight = 720
)

var (
	curtainImg *ebiten.Image
	titleImg   *ebiten.Image
	backgroundImg *ebiten.Image

	neutralCreativeImg   *ebiten.Image
	unhappyCreativeImg *ebiten.Image
	happyCreativeImg   *ebiten.Image

	neutralCasualImg   *ebiten.Image
	unhappyCasualImg *ebiten.Image
	happyCasualImg   *ebiten.Image

	neutralEliteImg   *ebiten.Image
	unhappyEliteImg *ebiten.Image
	happyEliteImg   *ebiten.Image

	defaultFontFace font.Face

	clickSound *audio.Player
	music *audio.Player
)

func init() {
	var err error

	audioContext, err := audio.NewContext(48000)
	if err != nil {
		log.Fatal(err)
	}

	f, err := ebitenutil.OpenFile("resources/click.wav")
	if err != nil {
		log.Fatal(err)
	}

	d, err := wav.Decode(audioContext, f)
	if err != nil {
		log.Fatal(err)
	}

	clickSound, err = audio.NewPlayer(audioContext, d)
	if err != nil {
		log.Fatal(err)
	}

	f, err = ebitenutil.OpenFile("resources/gradual_sunrise.mp3")
	if err != nil {
		log.Fatal(err)
	}

	d2, err := mp3.Decode(audioContext, f)
	if err != nil {
		log.Fatal(err)
	}

	music, err = audio.NewPlayer(audioContext, d2)
	if err != nil {
		log.Fatal(err)
	}

	// images
	curtainImg, _, err = ebitenutil.NewImageFromFile("resources/curtain.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	titleImg, _, err = ebitenutil.NewImageFromFile("resources/titlescreen.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	happyCreativeImg, _, err = ebitenutil.NewImageFromFile("resources/happy_creative.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	unhappyCreativeImg, _, err = ebitenutil.NewImageFromFile("resources/unhappy_creative.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	neutralCreativeImg, _, err = ebitenutil.NewImageFromFile("resources/neutral_creative.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	happyCasualImg, _, err = ebitenutil.NewImageFromFile("resources/happy_casual.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	unhappyCasualImg, _, err = ebitenutil.NewImageFromFile("resources/unhappy_casual.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	neutralCasualImg, _, err = ebitenutil.NewImageFromFile("resources/neutral_casual.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	happyEliteImg, _, err = ebitenutil.NewImageFromFile("resources/happy_elite.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	unhappyEliteImg, _, err = ebitenutil.NewImageFromFile("resources/unhappy_elite.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	neutralEliteImg, _, err = ebitenutil.NewImageFromFile("resources/neutral_elite.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	backgroundImg, _, err = ebitenutil.NewImageFromFile("resources/background.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}


	tt, err := truetype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	defaultFontFace = truetype.NewFace(tt, &truetype.Options{
		Size:    22,
		DPI:     80,
		Hinting: font.HintingFull,
	})
}

var game *Game

func update(screen *ebiten.Image) error {
	screen.Fill(color.White)
	if !music.IsPlaying() {
		music.Rewind()
		music.Play()
	}
	return game.currentScene.update(screen)
}

func main() {

	currentScene, err := NewTitleScene()
	if err != nil {
		log.Fatal(err)
	}

	game = &Game{
		currentScene: currentScene,
	}

	if err := ebiten.Run(update, screenWidth, screenHeight, 1, "Kingdom of Fandom"); err != nil {
		log.Fatal(err)
	}
}
