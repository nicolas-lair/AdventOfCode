package day2

type Hand int

const (
	Rock Hand = iota + 1
	Paper
	Scissors
)

func (h Hand) eval() int {
	return int(h)
}

var WinnerHand = map[Hand]Hand{
	Rock:     Paper,
	Paper:    Scissors,
	Scissors: Rock,
}

var LoserHand = map[Hand]Hand{
	Rock:     Scissors,
	Paper:    Rock,
	Scissors: Paper,
}

type Outcome int

const (
	Lose Outcome = 0
	Draw         = 3
	Win          = 6
)

func (h Outcome) eval() int {
	return int(h)
}
