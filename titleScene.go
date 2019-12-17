package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

func NewTitleScene() (Scene, error) {
	textX := screenWidth / 5
	textY := screenHeight*2 / 3
	text := "Rule your kingdom"
	boxX := textX - 10
	boxY := textY - 23
	boxHeight := 30
	boxWidth := 15 * len(text)

	box, err := NewTextBox(textX, textY, boxX, boxY, boxWidth, boxHeight, text, nil)
	if err != nil {
		return nil, err
	}

	return &titleScene{
		titleBox: box,
	}, nil
}

type titleScene struct {
	titleBox      *TextBox
	transitioning bool
	curtainy      int
}

func (ts *titleScene) update(screen *ebiten.Image) error {
	if err := screen.DrawImage(titleImg, &ebiten.DrawImageOptions{}); err != nil {
		return err
	}

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) && !ts.transitioning && ts.titleBox.isTouchingMouse() {
		ts.transitioning = true
		clickSound.Play()
		clickSound.Rewind()
	}

	if err := ts.titleBox.Draw(screen); err != nil {
		return err
	}

	if ts.transitioning {
		curtainOptions := &ebiten.DrawImageOptions{}
		_, curtainHeight := curtainImg.Size()
		curtainOptions.GeoM.Translate(0, float64(ts.curtainy-curtainHeight))
		if err := screen.DrawImage(curtainImg, curtainOptions); err != nil {
			return err
		}

		ts.curtainy += 10
	}

	if ts.curtainy >= screenHeight {
		choiceScene, err := NewChoiceScene()
		if err != nil {
			return err
		}
		game.previousScene = game.currentScene
		game.currentScene = choiceScene

		return nil
	}

	return nil
}
