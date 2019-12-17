package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/hajimehoshi/ebiten/text"
	"image/color"
)

func NewResultsScene(influence *Influence) (Scene, error) {
	textX := screenWidth / 5
	textY := screenHeight * 4 / 5
	text := "Restart"
	boxX := textX - 5
	boxY := textY - 23
	boxHeight := 30
	boxWidth := 15 * len(text)
	restartButton, _ := NewTextBox(textX, textY, boxX, boxY, boxWidth, boxHeight, text, nil)

	creativityMeter, _ := NewTextBox(60, 120+33, 60, 120, influence.creative, 75, "", nil)
	eliteMeter, _ := NewTextBox(60, 240+33, 60, 240, influence.elite, 75, "", nil)
	casualMeter, _ := NewTextBox(60, 360+33, 60, 360, influence.casual, 75, "", nil)


	return &InfluenceScene{
		restartButton: restartButton,
		influence: influence,
		casualMeter: casualMeter,
		eliteMeter: eliteMeter,
		creativityMeter: creativityMeter,
	}, nil
}

type InfluenceScene struct {
	influence       *Influence
	creativityMeter *TextBox
	eliteMeter      *TextBox
	casualMeter     *TextBox
	restartButton *TextBox
}

func (is *InfluenceScene) update(screen *ebiten.Image) error {
	if err := screen.DrawImage(backgroundImg, &ebiten.DrawImageOptions{}); err != nil {
		return err
	}

	text.Draw(screen, "Creative", defaultFontFace, 60, 120-15, color.Black)
	text.Draw(screen, fmt.Sprintf("%d", is.influence.creative), defaultFontFace, 60+is.influence.creative+10, 120+40, color.Black)

	text.Draw(screen, "Elite", defaultFontFace, 60, 240-15, color.Black)
	text.Draw(screen, fmt.Sprintf("%d", is.influence.elite), defaultFontFace, 60+is.influence.elite+10, 240+40, color.Black)

	text.Draw(screen, "Casual", defaultFontFace, 60, 360-15, color.Black)
	text.Draw(screen, fmt.Sprintf("%d", is.influence.casual), defaultFontFace, 60+is.influence.casual+10, 360+40, color.Black)


	is.creativityMeter.Draw(screen)
	is.eliteMeter.Draw(screen)
	is.casualMeter.Draw(screen)
	is.restartButton.Draw(screen)

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) && is.restartButton.isTouchingMouse() {
		clickSound.Play()
		clickSound.Rewind()
		choiceScene, err := NewChoiceScene()
		if err != nil {
			return err
		}
		game.previousScene = game.currentScene
		game.currentScene = choiceScene
	}

	return nil
}
