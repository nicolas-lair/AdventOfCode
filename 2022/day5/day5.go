package day5

import (
	"AoC22/utils"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func Day5(crateMoverModel int) string {
	fileContent := utils.ReadEntireFile("2022/day5/input")
	input := strings.Split(fileContent, "\n\n")
	gridString, instructionsString := input[0], input[1]

	grid := new(Grid)
	grid.parseGrid(gridString)

	for _, instruction := range parseInstructions(instructionsString) {
		//fmt.Println(grid)
		//fmt.Println(grid.getItemCount())
		//fmt.Println("\n")
		//fmt.Println(instruction)
		grid.move(instruction, crateMoverModel)
	}
	result := grid.getTopItems()
	return result
}

type Instruction struct {
	nMove int
	start int
	end   int
}

func (i Instruction) String() string {
	return fmt.Sprintf("Move %v from %v to %v", i.nMove, i.start, i.end)
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

func (g *Grid) move(i Instruction, crateMoverModel int) {
	switch crateMoverModel {
	case 9000:
		g.move9000(i)
	case 9001:
		g.move9001(i)
	default:
		log.Fatal("No such model")
	}
}

func (g *Grid) move9000(i Instruction) {
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

func (g *Grid) move9001(i Instruction) {
	itemStack := make([]string, i.nMove)
	copy(itemStack, g.gridArray[i.start-1][:i.nMove])

	newStartStack := make([]string, len(g.gridArray[i.start-1])-i.nMove)
	copy(newStartStack, g.gridArray[i.start-1][i.nMove:])
	g.gridArray[i.start-1] = newStartStack

	newEndStack := make([]string, len(g.gridArray[i.end-1])+i.nMove)
	copy(newEndStack, append(itemStack, g.gridArray[i.end-1]...))
	g.gridArray[i.end-1] = newEndStack
}

func (g *Grid) getTopItems() string {
	topItems := ""
	for col := 0; col < len(g.gridArray); col++ {
		topItems += g.gridArray[col][0]
	}
	return topItems
}

func (g *Grid) getItemCount() map[string]int {
	nItems := make(map[string]int)
	for _, stack := range g.gridArray {
		for _, item := range stack {
			_, ok := nItems[item]
			if ok {
				nItems[item] += 1
			} else {
				nItems[item] = 1
			}
		}
	}
	return nItems
}

func (g *Grid) getMaxHeight() int {
	heightList := utils.SliceMap[[]string, int](g.gridArray, func(i []string) int {
		return len(i)
	})
	maxHeight := utils.MaxInts(heightList)
	return maxHeight
}

func (g *Grid) String() string {
	nCol := len(g.gridArray)
	repr := ""

	var newLine string
	var value string
	maxHeight := g.getMaxHeight()
	for h := maxHeight; h >= 0; h-- {
		newLine = ""
		for i := 0; i < nCol; i++ {
			if len(g.gridArray[i]) > h {
				value = g.gridArray[i][len(g.gridArray[i])-h-1]
				newLine += " " + value + " "
			} else {
				newLine += "   "
			}
		}
		repr += newLine + "\n"
	}

	for i := 0; i < nCol; i++ {
		repr += fmt.Sprintf(" %v ", i+1)
	}

	return repr
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
