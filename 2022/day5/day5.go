package day5

import (
	"AoC22/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Part1() string {
	fileContent := utils.ReadEntireFile("2022/day5/input")
	input := strings.Split(fileContent, "\n\n")
	gridString, instructionsString := input[0], input[1]

	grid := new(Grid)
	grid.parseGrid(gridString)

	for _, instruction := range parseInstructions(instructionsString) {
		fmt.Println(grid.gridArray)
		fmt.Println(instruction)
		grid.move(instruction)
	}
	result := grid.getTopItems()
	return result
}

type Instruction struct {
	nMove int
	start int
	end   int
}

type Grid struct {
	gridArray [][]string
}

func (g *Grid) parseGrid(gridString string) {
	lineGrid := strings.Split(gridString, "\n")
	nRows := len(lineGrid) - 1
	nCols := len(strings.Fields(lineGrid[nRows]))

	g.gridArray = make([][]string, nCols)
	for _, line := range lineGrid[:nRows] {
		g.parseGridLine(line)
	}
}

func (g *Grid) parseGridLine(gridLine string) {
	re := regexp.MustCompile(`\[`)
	index := re.FindAllIndex([]byte(gridLine), -1)
	for _, i := range index {
		item := string(gridLine[i[0]+1])
		col := i[0] / 4
		g.gridArray[col] = append(g.gridArray[col], item)
	}
}

func (g *Grid) move(i Instruction) {
	var item string
	for m := 0; m < i.nMove; m++ {
		// get item
		item = g.gridArray[i.start-1][0]

		// Move to the right stack
		g.gridArray[i.end-1] = append([]string{item}, g.gridArray[i.end-1]...)

		// Remove from initial stack
		g.gridArray[i.start-1] = g.gridArray[i.start-1][1:]
	}
}

func (g *Grid) getTopItems() string {
	topItems := ""
	for col := 0; col < len(g.gridArray); col++ {
		topItems += g.gridArray[col][0]
	}
	return topItems
}

func parseInstructions(instructionString string) []Instruction {
	var instructionList []Instruction
	for _, line := range strings.Split(instructionString, "\n") {
		if len(line) == 0 {
			break
		}
		instructionList = append(instructionList, parseInstructionLine(line))
	}
	return instructionList
}

func parseInstructionLine(instructionLine string) Instruction {
	re := regexp.MustCompile("\\d+")
	instructionValue := re.FindAll([]byte(instructionLine), -1)
	nMove, _ := strconv.Atoi(string(instructionValue[0]))
	start, _ := strconv.Atoi(string(instructionValue[1]))
	end, _ := strconv.Atoi(string(instructionValue[2]))

	return Instruction{nMove: nMove, start: start, end: end}
}
