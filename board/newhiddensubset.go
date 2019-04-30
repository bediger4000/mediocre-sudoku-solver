package board

import "fmt"

func (bd *Board) HiddenPair(announce bool) int {
	eliminated := 0
	for colNo, column := range Columns {
		digitMap := make(map[int][]CellCoord)
		for _, coord := range column {
			if !bd[coord.X][coord.Y].Solved {
				for _, digit := range bd[coord.X][coord.Y].Possible {
					digitMap[digit] = append(digitMap[digit], coord)
				}
			}
		}
		pairCells := make(map[int][]CellCoord)
		for digit, cells := range digitMap {
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
						if announce {
							fmt.Printf("Column %d: %d & %d a hidden pair in %v and %v\n",
								colNo, digit, otherDigit, cells, otherCells)
						}
						pairedAlready[digit] = true
						pairedAlready[otherDigit] = true
						eliminated += cleanHiddenPair(bd, digit, otherDigit, cells, announce)
					}
				}
			}
		}
	}
	for rowNo, row := range Rows {
		digitMap := make(map[int][]CellCoord)
		for _, coord := range row {
			if !bd[coord.X][coord.Y].Solved {
				for _, digit := range bd[coord.X][coord.Y].Possible {
					digitMap[digit] = append(digitMap[digit], coord)
				}
			}
		}
		pairCells := make(map[int][]CellCoord)
		for digit, cells := range digitMap {
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
						if announce {
							fmt.Printf("Row %d: %d & %d a hidden pair in %v and %v\n",
								rowNo, digit, otherDigit, cells, otherCells)
						}
						pairedAlready[digit] = true
						pairedAlready[otherDigit] = true
						eliminated += cleanHiddenPair(bd, digit, otherDigit, cells, announce)
					}
				}
			}
		}
	}
	for blockNo, block := range Blocks {
		digitMap := make(map[int][]CellCoord)
		for _, coord := range block {
			if !bd[coord.X][coord.Y].Solved {
				for _, digit := range bd[coord.X][coord.Y].Possible {
					digitMap[digit] = append(digitMap[digit], coord)
				}
			}
		}
		pairCells := make(map[int][]CellCoord)
		for digit, cells := range digitMap {
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
						if announce {
							fmt.Printf("Block %d: %d & %d a hidden pair in %v and %v\n",
								blockNo, digit, otherDigit, cells, otherCells)
						}
						pairedAlready[digit] = true
						pairedAlready[otherDigit] = true
						eliminated += cleanHiddenPair(bd, digit, otherDigit, cells, announce)
						// For the block case, it's possible to eliminate the
						// digits of the hidden pair from other cells,
						// if the pair shows up in the same row or column.
						if cells[0].X == cells[1].X {
							// same row
							row := Rows[cells[0].X]
							for _, cell := range row {
								if bd[cell.X][cell.Y].Solved {
									continue
								}
								if (cell.X == cells[0].X && cell.Y == cells[0].Y) ||
									(cell.X == cells[1].X && cell.Y == cells[1].Y) {
									continue
								}
								eliminated += bd.SpliceOut(cell.X, cell.Y, digit)
								eliminated += bd.SpliceOut(cell.X, cell.Y, otherDigit)
							}
						} else if cells[0].Y == cells[1].Y {
							// same column
							col := Columns[cells[0].Y]
							for _, cell := range col {
								if bd[cell.X][cell.Y].Solved {
									continue
								}
								if cell.X == cells[0].X && cell.Y == cells[0].Y ||
									cell.X == cells[1].X && cell.Y == cells[1].Y {
									continue
								}
								eliminated += bd.SpliceOut(cell.X, cell.Y, digit)
								eliminated += bd.SpliceOut(cell.X, cell.Y, otherDigit)
							}
						}
					}
				}
			}
		}
	}
	return eliminated
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

func cleanHiddenPair(bd *Board, digit int, otherDigit int, cells []CellCoord, announce bool) int {
	eliminated := 0
	for _, cell := range cells {
		maybeDigits := make([]int, len(bd[cell.X][cell.Y].Possible))
		copy(maybeDigits, bd[cell.X][cell.Y].Possible)
		for _, possibleDigit := range maybeDigits {
			if (possibleDigit != digit) && possibleDigit != otherDigit {
				n := bd.SpliceOut(cell.X, cell.Y, possibleDigit)
				if n > 0 && announce {
					fmt.Printf("Eliminated %d from <%d,%d>\n", possibleDigit, cell.X, cell.Y)
				}
				eliminated += n
			}
		}
	}
	return eliminated
}
