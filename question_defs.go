package main

import (
	"math/rand"
)

type AnswerDef struct {
	text      string
	influence Influence
}

type QuestionDef struct {
	text    string
	answers []string
	qtype   string
}

var simple_question_pool = []QuestionDef{
	QuestionDef{text: "Let’s add a new character, they are a ...", qtype: "simple", answers: []string{"a"}},
	QuestionDef{text: "Going to write off this character with a ...", qtype: "simple", answers: []string{"b"}},
	QuestionDef{text: "Tone of the next part of the story will be ...", qtype: "simple", answers: []string{"c"}},
	QuestionDef{text: "Next background will be set in ...", qtype: "simple", answers: []string{"d"}},
	QuestionDef{text: "Music will be in <answer> style ", qtype: "simple", answers: []string{"e"}},
	QuestionDef{text: "The crucial macguffin is ...", qtype: "simple", answers: []string{"f"}},
	QuestionDef{text: "Merchandise we’re going to sell will be ...", qtype: "simple", answers: []string{"a", "f"}},
}

var structural_question_pool = []QuestionDef{
	QuestionDef{text: "The next project should be a ...", qtype: "structural", answers: []string{"g"}},
	QuestionDef{text: "The next project is going to be a sequel", qtype: "structural", answers: []string{"h"}},
	QuestionDef{text: "The next project will be a reboot", qtype: "structural", answers: []string{"i"}},
}

var answer_pool = map[string][]AnswerDef{
	"a": []AnswerDef{
		AnswerDef{
			text: "cute baby",
			influence: Influence{
				creative: 1,
				elite:    -1,
				casual:   1,
			},
		},
		AnswerDef{
			text: "brooding dark edgeperson",
			influence: Influence{
				creative: -1,
				elite:    1,
				casual:   1,
			},
		},
		AnswerDef{
			text: "iconoclastic idealist",
			influence: Influence{
				creative: 1,
				elite:    -1,
				casual:   -1,
			},
		},
		AnswerDef{
			text: "know-it-all",
			influence: Influence{
				creative: -1,
				elite:    1,
				casual:   -1,
			},
		},
		AnswerDef{
			text: "sexy fanservice character ",
			influence: Influence{
				creative: -1,
				elite:    -1,
				casual:   1,
			},
		},
		AnswerDef{
			text: "honorable warrior",
			influence: Influence{
				creative: -1,
				elite:    1,
				casual:   1,
			},
		},
		AnswerDef{
			text: "animal mascot",
			influence: Influence{
				creative: 1,
				elite:    -1,
				casual:   1,
			},
		},
	},
	"b": []AnswerDef{
		AnswerDef{
			text: "kill them all",
			influence: Influence{
				creative: 1,
				elite:    -1,
				casual:   -1,
			},
		},
		AnswerDef{
			text: "tragic meaningless death",
			influence: Influence{
				creative: -1,
				elite:    -1,
				casual:   1,
			},
		},
		AnswerDef{
			text: "gone to their home planet",
			influence: Influence{
				creative: -1,
				elite:    -1,
				casual:   -1,
			},
		},
		AnswerDef{
			text: "sent away on a bus",
			influence: Influence{
				creative: -1,
				elite:    -1,
				casual:   -1,
			},
		},
		AnswerDef{
			text: "heroic sacrifice",
			influence: Influence{
				creative: 1,
				elite:    1,
				casual:   1,
			},
		},
		AnswerDef{
			text: "heavily foreshadowed",
			influence: Influence{
				creative: 1,
				elite:    1,
				casual:   -1,
			},
		},
	},
	"c": []AnswerDef{
		AnswerDef{
			text: "dark",
			influence: Influence{
				creative: -1,
				elite:    1,
				casual:   1,
			},
		},
		AnswerDef{
			text: "grimdark",
			influence: Influence{
				creative: 1,
				elite:    -1,
				casual:   -1,
			},
		},
		AnswerDef{
			text: "comedic",
			influence: Influence{
				creative: 1,
				elite:    -1,
				casual:   1,
			},
		},
		AnswerDef{
			text: "melodramatic",
			influence: Influence{
				creative: -1,
				elite:    -1,
				casual:   1,
			},
		},
		AnswerDef{
			text: "romantic",
			influence: Influence{
				creative: 1,
				elite:    -1,
				casual:   1,
			},
		},
	},
	"d": []AnswerDef{
		AnswerDef{
			text: "snow covered mountains",
			influence: Influence{
				creative: func() int { return -1 + rand.Intn(2) }(),
				elite:    func() int { return -1 + rand.Intn(2) }(),
				casual:   func() int { return -1 + rand.Intn(2) }(),
			},
		},
		AnswerDef{
			text: "exquisite antique mansion",
			influence: Influence{
				creative: func() int { return -1 + rand.Intn(2) }(),
				elite:    func() int { return -1 + rand.Intn(2) }(),
				casual:   func() int { return -1 + rand.Intn(2) }(),
			},
		},
		AnswerDef{
			text: "dystopian cityscape",
			influence: Influence{
				creative: func() int { return -1 + rand.Intn(2) }(),
				elite:    func() int { return -1 + rand.Intn(2) }(),
				casual:   func() int { return -1 + rand.Intn(2) }(),
			},
		},
		AnswerDef{
			text: "quiet suburbs",
			influence: Influence{
				creative: func() int { return -1 + rand.Intn(2) }(),
				elite:    func() int { return -1 + rand.Intn(2) }(),
				casual:   func() int { return -1 + rand.Intn(2) }(),
			},
		},
		AnswerDef{
			text: "sunny beaches",
			influence: Influence{
				creative: func() int { return -1 + rand.Intn(2) }(),
				elite:    func() int { return -1 + rand.Intn(2) }(),
				casual:   func() int { return -1 + rand.Intn(2) }(),
			},
		},
	},
	"e": []AnswerDef{
		AnswerDef{
			text: "classical",
			influence: Influence{
				creative: func() int { return -1 + rand.Intn(2) }(),
				elite:    func() int { return -1 + rand.Intn(2) }(),
				casual:   func() int { return -1 + rand.Intn(2) }(),
			},
		},
		AnswerDef{
			text: "8-bit",
			influence: Influence{
				creative: func() int { return -1 + rand.Intn(2) }(),
				elite:    func() int { return -1 + rand.Intn(2) }(),
				casual:   func() int { return -1 + rand.Intn(2) }(),
			},
		},
		AnswerDef{
			text: "rock",
			influence: Influence{
				creative: func() int { return -1 + rand.Intn(2) }(),
				elite:    func() int { return -1 + rand.Intn(2) }(),
				casual:   func() int { return -1 + rand.Intn(2) }(),
			},
		},
		AnswerDef{
			text: "pop",
			influence: Influence{
				creative: func() int { return -1 + rand.Intn(2) }(),
				elite:    func() int { return -1 + rand.Intn(2) }(),
				casual:   func() int { return -1 + rand.Intn(2) }(),
			},
		},
		AnswerDef{
			text: "emo",
			influence: Influence{
				creative: -1,
				elite:    1,
				casual:   -1,
			},
		},
		AnswerDef{
			text: "kpop",
			influence: Influence{
				creative: 1,
				elite:    1,
				casual:   1,
			},
		},
	},
	"f": []AnswerDef{
		AnswerDef{
			text: "golden idol",
			influence: Influence{
				creative: -1,
				elite:    1,
				casual:   -1,
			},
		},
		AnswerDef{
			text: "friendship",
			influence: Influence{
				creative: 1,
				elite:    -1,
				casual:   -1,
			},
		},
		AnswerDef{
			text: "rogue ai",
			influence: Influence{
				creative: -1,
				elite:    1,
				casual:   -1,
			},
		},
		AnswerDef{
			text: "giant robots",
			influence: Influence{
				creative: 1,
				elite:    -1,
				casual:   1,
			},
		},
		AnswerDef{
			text: "magic",
			influence: Influence{
				creative: 1,
				elite:    -1,
				casual:   1,
			},
		},
		AnswerDef{
			text: "gun hung on a wall",
			influence: Influence{
				creative: -1,
				elite:    1,
				casual:   1,
			},
		},
	},
	"g": []AnswerDef{
		AnswerDef{
			text: "TV",
		},
		AnswerDef{
			text: "Movie",
		},
		AnswerDef{
			text: "Book",
		},
		AnswerDef{
			text: "whatever, we will make a game",
		},
	},
	"h": []AnswerDef{
		AnswerDef{
			text: "yes",
		},
		AnswerDef{
			text: "no",
		},
	},
	"i": []AnswerDef{
		AnswerDef{
			text: "yes",
		},
		AnswerDef{
			text: "no",
		},
	},
}
