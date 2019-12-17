package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/text"
	"image/color"
	"strings"

	"github.com/hajimehoshi/ebiten"
	//"github.com/hajimehoshi/ebiten/ebitenutil"
	"golang.org/x/image/font"
)

func NewTextBox(x, y, boxX, boxY, boxWidth, boxHeight int, text string, img *ebiten.Image) (*TextBox, error) {
	maxLineLen := boxWidth / 13

	var lines []string
	words := strings.Fields(text)

	line := ""
	for i, word := range words {
		if len(line)+len(word)+1 <= maxLineLen {
			line = fmt.Sprintf("%s %s", line, word)

			if i == len(words)-1 {
				lines = append(lines, line)
				line = ""
			}
		} else {
			lines = append(lines, line)
			line = ""
		}
	}

	return &TextBox{
		x:            x,
		y:            y,
		boxX:         boxX,
		boxY:         boxY,
		boxHeight:    boxHeight,
		boxWidth:     boxWidth,
		text:         lines,
		font:         defaultFontFace,
		highlightBox: true,
		color:        color.White,
		img:          img,
	}, nil
}

type TextBox struct {
	x            int
	y            int
	boxX         int
	boxY         int
	boxWidth     int
	boxHeight    int
	highlightBox bool
	text         []string
	font         font.Face
	color        color.Color
	img          *ebiten.Image
}

func (tb *TextBox) Draw(screen *ebiten.Image) error {

	if tb.isTouchingMouse() && tb.highlightBox {
		tb.color = color.Color(color.RGBA{0xed, 0xc5, 0x3f, 0xff})
	} else {
		tb.color = color.Black
	}

	if tb.img == nil  && len(tb.text) == 0 {
		ebitenutil.DrawRect(screen, float64(tb.boxX), float64(tb.boxY), float64(tb.boxWidth), float64(tb.boxHeight), color.Black)
	}

	if tb.img != nil {
		imgOptions := &ebiten.DrawImageOptions{}
		imgOptions.GeoM.Translate(float64(tb.boxX), float64(tb.boxY))
		if err := screen.DrawImage(tb.img, imgOptions); err != nil {
			return err
		}
	}

	for i, line := range tb.text {
		text.Draw(screen, line, tb.font, tb.x, tb.y+i*28, tb.color)
	}

	return nil
}

func (tb *TextBox) isTouchingMouse() bool {
	x, y := ebiten.CursorPosition()

	return x > tb.boxX && x < tb.boxX+tb.boxWidth && y > tb.boxY && y < tb.boxY+tb.boxHeight
}
