package day2

import (
	"AoC22/utils"
	"log"
	"strings"
)

func Part1() int {
	scanner := utils.ReadDataLineByLine("data/input.txt")
	var score int
	for scanner.Scan() {
		line := scanner.Text()
		oppHand, myHand := parseLinePart1(line)
		score += myHand.eval() + getOutcome(myHand, oppHand).eval()
	}
	return score

}

func parseLinePart1(x string) (Hand, Hand) {
	letters := strings.Fields(x)
	oppHand := parseHandPart1(letters[0])
	myHand := parseHandPart1(letters[1])
	return oppHand, myHand
}

func parseHandPart1(stringHand string) Hand {
	var handValue Hand
	switch stringHand {
	case "A", "X":
		handValue = Rock
	case "B", "Y":
		handValue = Paper
	case "C", "Z":
		handValue = Scissors
	default:
		log.Fatal("Invalid Hand")
	}
	return handValue
}

func getOutcome(myHand Hand, oppHand Hand) Outcome {
	var outcome Outcome
	if myHand == WinnerHand[oppHand] {
		outcome = Win
	} else if myHand == LoserHand[oppHand] {
		outcome = Lose
	} else if myHand == oppHand {
		outcome = Draw
	} else {
		log.Fatal("Invalid round")
	}
	return outcome
}
