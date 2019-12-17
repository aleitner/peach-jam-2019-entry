package main

type Influence struct {
	casual   int
	elite    int
	creative int
}

type Answer struct {
	display   *TextBox
	influence Influence
}

func NewAnswer(x, y, boxX, boxY, boxWidth, boxHeight int, text string, values Influence) (*Answer, error) {
	answerBox, err := NewTextBox(x, y, boxX, boxY, boxWidth, boxHeight, text, nil)
	if err != nil {
		return nil, err
	}

	return &Answer{
		display:   answerBox,
		influence: values,
	}, nil
}
