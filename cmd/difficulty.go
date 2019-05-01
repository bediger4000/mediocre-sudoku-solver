package main

import (
	"fmt"
	"os"
	"sudoku/board"
)

type ClueCountLevel struct {
	minClues   int
	lowerBound int
	level      string
}

var ClueCountLevels []ClueCountLevel = []ClueCountLevel{
	ClueCountLevel{47, 5, "Extremely Easy"},
	ClueCountLevel{36, 4, "Easy"},
	ClueCountLevel{32, 3, "Medium"},
	ClueCountLevel{28, 2, "Difficult"},
	ClueCountLevel{17, 0, "Evil"},
}

func main() {

	bd := board.ReadBoard(os.Stdin)

	clueCount := 0
	for _, row := range board.Rows {
		for _, cell := range row {
			if bd[cell.X][cell.Y].Solved {
				clueCount++
			}
		}
	}

	countLevel := "unset"

	for _, clueLevel := range ClueCountLevels {
		if clueCount >= clueLevel.minClues {
			countLevel = clueLevel.level
			break
		}
	}

	minRowCount := 99
	for _, row := range board.Rows {
		clueCount = 0
		for _, cell := range row {
			if bd[cell.X][cell.Y].Solved {
				clueCount++
			}
		}
		if clueCount < minRowCount {
			minRowCount = clueCount
		}
	}
	minColCount := 99
	for _, col := range board.Columns {
		clueCount = 0
		for _, cell := range col {
			if bd[cell.X][cell.Y].Solved {
				clueCount++
			}
		}
		if clueCount < minColCount {
			minColCount = clueCount
		}
	}
	rowLevel := "unset"
	for _, clueLevel := range ClueCountLevels {
		if minRowCount >= clueLevel.lowerBound {
			rowLevel = clueLevel.level
			break
		}
	}
	colLevel := "unset"
	for _, clueLevel := range ClueCountLevels {
		if minColCount >= clueLevel.lowerBound {
			colLevel = clueLevel.level
			break
		}
	}

	fmt.Printf("%d/%d/%d : %s, %s, %s\n", clueCount, minRowCount, minColCount, countLevel, rowLevel, colLevel)
}
