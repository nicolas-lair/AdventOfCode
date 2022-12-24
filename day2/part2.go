package day2

import (
	"AoC22/utils"
	"log"
	"strings"
)

func Part2() int {
	scanner := utils.ReadDataLineByLine("data/input.txt")

	var score int
	for scanner.Scan() {
		line := scanner.Text()
		oppHand, desiredOutcome := parseLinePart2(line)
		myHand := findHandToPlay(oppHand, desiredOutcome)
		score += myHand.eval() + getOutcome(myHand, oppHand).eval()
	}
	return score

}

func parseLinePart2(x string) (Hand, Outcome) {
	letters := strings.Fields(x)
	oppHand := parseHandPart1(letters[0])
	desiredOutcome := getDesiredOutcome(letters[1])
	return oppHand, desiredOutcome
}

func getDesiredOutcome(letterHand string) Outcome {
	var desiredOutcome Outcome
	switch letterHand {
	case "X":
		desiredOutcome = Lose
	case "Y":
		desiredOutcome = Draw
	case "Z":
		desiredOutcome = Win
	default:
		log.Fatal("Invalid Hand")
	}
	return desiredOutcome
}

func findHandToPlay(oppHand Hand, outcome Outcome) Hand {
	var myHand Hand
	switch outcome {
	case Lose:
		myHand = LoserHand[oppHand]
	case Draw:
		myHand = oppHand
	case Win:
		myHand = WinnerHand[oppHand]
	}
	return myHand
}
