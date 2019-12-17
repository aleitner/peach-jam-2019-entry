package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/hajimehoshi/ebiten/text"
	"image/color"
	"math/rand"
	"sync"
	"time"
)

func NewChoiceScene() (Scene, error) {
	casualPortrait, err := NewTextBox(60+10, 145+33, 60, 145, 200, 200, "", neutralCasualImg)
	if err != nil {
		return nil, err
	}

	elitePortrait, err := NewTextBox(380+10, 145+33, 380, 145, 200, 200, "", neutralEliteImg)
	if err != nil {
		return nil, err
	}

	creativePortrait, err := NewTextBox(700+10, 145+33, 700, 145, 200, 200, "", neutralCreativeImg)
	if err != nil {
		return nil, err
	}

	choiceScene := &ChoiceScene{
		casualPortrait:   casualPortrait,
		elitePortrait:    elitePortrait,
		creativePortrait: creativePortrait,
		influence: &Influence{
			casual:   500,
			elite:    300,
			creative: 100,
		},
		highscore:   &Influence{},
		done:        make(chan bool),
		currentRate: &Influence{
				casual: 1,
				elite: 5,
				creative: 10,
		},
	}

	if choiceScene.questionBox, choiceScene.answerBoxes, err = choiceScene.popNextQuestion(); err != nil {
		return nil, err
	}

	ticker := time.NewTicker(time.Second)

	go func() {
		tickCount := 0
		for {
			select {
			case <-choiceScene.done:
				return
			case <-ticker.C:
				tickCount++

				choiceScene.mtx.Lock()
				if choiceScene.influence.casual > 0 && choiceScene.currentRate.casual > 0 && tickCount % choiceScene.currentRate.casual == 0 {
					choiceScene.influence.casual -= 10
					if choiceScene.influence.casual < 0 {
						choiceScene.influence.casual = 0
					}
				}

				if choiceScene.influence.elite > 0 && choiceScene.currentRate.elite > 0 && tickCount % choiceScene.currentRate.elite == 0 {
					choiceScene.influence.elite -= 10
					if choiceScene.influence.elite < 0 {
						choiceScene.influence.elite = 0
					}
				}

				if choiceScene.influence.creative > 0 && choiceScene.currentRate.creative > 0 && tickCount % choiceScene.currentRate.creative == 0 {
					choiceScene.influence.creative -= 10
					if choiceScene.influence.creative < 0 {
						choiceScene.influence.creative = 0
					}
				}

				choiceScene.mtx.Unlock()
			}
		}
	}()

	return choiceScene, nil
}

type ChoiceScene struct {
	questionIndex      int
	questionList       []QuestionDef
	questionBox        *TextBox
	currentQuestionDef QuestionDef
	answerBoxes        []*Answer
	characterPortraits *TextBox
	casualPortrait     *TextBox
	elitePortrait      *TextBox
	creativePortrait   *TextBox
	influence          *Influence
	currentRate        *Influence
	highscore          *Influence
	done               chan bool
	mtx                sync.Mutex
}

func (cs *ChoiceScene) update(screen *ebiten.Image) error {

	if err := screen.DrawImage(backgroundImg, &ebiten.DrawImageOptions{}); err != nil {
		return err
	}

	if cs.influence.casual == 0 && cs.influence.creative == 0 && cs.influence.elite == 0 {
		cs.done <- true
		resultsScene, err := NewResultsScene(cs.highscore)
		if err != nil {
			return err
		}

		game.previousScene = game.currentScene
		game.currentScene = resultsScene
	}

	if err := cs.questionBox.Draw(screen); err != nil {
		return err
	}

	for _, answer := range cs.answerBoxes {
		answer.display.Draw(screen)

		if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) && answer.display.isTouchingMouse() {
			clickSound.Play()
			clickSound.Rewind()
			cs.mtx.Lock()
			cs.influence.creative += answer.influence.creative * 10
			if answer.influence.creative > 0 {
				cs.creativePortrait.img = happyCreativeImg
			} else if answer.influence.creative < 0 {
				cs.creativePortrait.img = unhappyCreativeImg
			} else {
				cs.creativePortrait.img = neutralCreativeImg
			}
			if cs.influence.creative > cs.highscore.creative {
				cs.highscore.creative = cs.influence.creative
			}
			if cs.influence.creative < 0 {
				cs.influence.creative = 0
			}

			cs.influence.casual += answer.influence.casual * 10
			if answer.influence.casual > 0 {
				cs.casualPortrait.img = happyCasualImg
			} else if answer.influence.casual < 0 {
				cs.casualPortrait.img = unhappyCasualImg
			} else {
				cs.casualPortrait.img = neutralCasualImg
			}
			if cs.influence.casual > cs.highscore.casual {
				cs.highscore.casual = cs.influence.casual
			}
			if cs.influence.casual < 0 {
				cs.influence.casual = 0
			}

			cs.influence.elite += answer.influence.elite * 10
			if answer.influence.elite > 0 {
				cs.elitePortrait.img = happyEliteImg
			} else if answer.influence.elite < 0 {
				cs.elitePortrait.img = unhappyEliteImg
			} else {
				cs.elitePortrait.img = neutralEliteImg
			}
			if cs.influence.elite > cs.highscore.elite {
				cs.highscore.elite = cs.influence.elite
			}
			if cs.influence.elite < 0 {
				cs.influence.elite = 0
			}

			if cs.currentQuestionDef.qtype == "structural" {
				chance := rand.Intn(9)
				if cs.currentQuestionDef.text == "no" {
					if true {
						cs.currentRate.creative /= 50000
						cs.currentRate.elite /= 50000
						cs.currentRate.casual /= 500000
					}
				} else {
					if chance > 8 {
						cs.currentRate.creative *= 2
						cs.currentRate.elite *= 2
						cs.currentRate.casual *= 2
					}
				}
			}

			cs.mtx.Unlock()

			var err error
			if cs.questionBox, cs.answerBoxes, err = cs.popNextQuestion(); err != nil {
				return err
			}
		}
	}

	text.Draw(screen, fmt.Sprintf("%d", cs.influence.casual), defaultFontFace, 127, 145-33, color.Black)
	text.Draw(screen, fmt.Sprintf("%d", cs.influence.elite), defaultFontFace, 457, 145-33, color.Black)
	text.Draw(screen, fmt.Sprintf("%d", cs.influence.creative), defaultFontFace, 777, 145-33, color.Black)

	cs.casualPortrait.Draw(screen)
	cs.elitePortrait.Draw(screen)
	cs.creativePortrait.Draw(screen)
	return nil
}

func buildQuestionList() []QuestionDef {
	var questions []QuestionDef

	questions = append(questions, simple_question_pool...)
	questions = append(questions, structural_question_pool[rand.Intn(len(structural_question_pool))])

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(questions), func(i, j int) { questions[i], questions[j] = questions[j], questions[i] })

	return questions
}

func randomAnswers(answers []AnswerDef) []AnswerDef {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(answers), func(i, j int) { answers[i], answers[j] = answers[j], answers[i] })

	numAnswers := len(answers)
	if numAnswers > 3 {
		numAnswers = 3
	}

	return answers[:numAnswers]
}

func (cs *ChoiceScene) popNextQuestion() (question *TextBox, answers []*Answer, err error) {
	if len(cs.questionList) == 0 {
		cs.questionList = buildQuestionList()
	}

	next := cs.questionList[0]
	cs.questionList = cs.questionList[1:]

	cs.currentQuestionDef = next

	question, err = NewTextBox(60+10, 380+33, 60, 380, 840, 75, next.text, nil)
	if err != nil {
		return nil, nil, err
	}
	question.highlightBox = false

	var potentialAnswers []AnswerDef
	for _, answerRef := range next.answers {
		potentialAnswers = append(potentialAnswers, answer_pool[answerRef]...)
	}

	// Pick 3 random answers from answer_pool[answerRef] list
	answerChoices := randomAnswers(potentialAnswers)

	for i, answerChoice := range answerChoices {
		answer, err := NewAnswer(60+10, 465+(i*75)+33, 60, 465+(i*75), 840, 75, answerChoice.text, answerChoice.influence)
		if err != nil {
			return nil, nil, err
		}
		answers = append(answers, answer)
	}

	return question, answers, nil
}
