package board

import "fmt"

func (bd *Board) HiddenPair(announceSolutions bool) int {
	found := 0
	for colNo, column := range Columns {
		digitMap := make(map[int][]CellCoord)
		for _, coord := range column {
			if !bd[coord.X][coord.Y].Solved {
				for _, digit := range bd[coord.X][coord.Y].Possible {
					digitMap[digit] = append(digitMap[digit], coord)
				}
			}
		}
		// fmt.Printf("Column %d:\n", colNo)
		pairCells := make(map[int][]CellCoord)
		for digit, cells := range digitMap {
			// fmt.Printf("\t%d %v\n", digit, cells)
			if len(cells) == 2 {
				pairCells[digit] = cells
			}
		}
		if len(pairCells) >= 2 {
			pairedAlready := make(map[int]bool)
			for digit, cells := range pairCells {
				for otherDigit, otherCells := range pairCells {
					if digit == otherDigit {
						continue
					}
					if pairedAlready[digit] || pairedAlready[otherDigit] {
						continue
					}
					if pairOfCellsEqual(cells, otherCells) {
						if announceSolutions {
							fmt.Printf("Column %d: %d & %d a hidden pair in %v and %v\n",
								colNo, digit, otherDigit, cells, otherCells)
						}
						pairedAlready[digit] = true
						pairedAlready[otherDigit] = true
						found++
						cleanHiddenPair(bd, digit, otherDigit, cells)
					}
				}
			}
		}
	}
	/*
		for _, row := range Rows {
			for _, coord := range row {
			}
		}
	*/
	return found
}

// It says []CellCoord, but each of them is length 2
func pairOfCellsEqual(a, b []CellCoord) bool {
	// a[0] == b[0] && a[1] == b[1]
	// or
	// a[1] == b[0] && a[0] == b[1]
	if (a[0] == b[0] && a[1] == b[1]) || (a[1] == b[0] && a[0] == b[1]) {
		return true
	}
	return false
}

func cellsEqual(a, b CellCoord) bool {
	return a.X == b.X && a.Y == b.Y
}

func cleanHiddenPair(bd *Board, digit int, otherDigit int, cells []CellCoord) {
	for _, cell := range cells {
		for _, possibleDigit := range bd[cell.X][cell.Y].Possible {
			if possibleDigit != digit && possibleDigit != otherDigit {
				bd.SpliceOut(cell.X, cell.Y, possibleDigit)
			}
		}
	}
}
